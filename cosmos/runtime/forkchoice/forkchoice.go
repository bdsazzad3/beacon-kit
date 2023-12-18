// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

// Package forker implements the Ethereum forker.
package forkchoice

import (
	"time"

	"github.com/prysmaticlabs/prysm/v4/consensus-types/interfaces"
	pb "github.com/prysmaticlabs/prysm/v4/proto/engine/v1"

	"cosmossdk.io/log"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/itsdevbear/bolaris/beacon/blockchain"
	"github.com/itsdevbear/bolaris/beacon/execution"
	beaconkeeper "github.com/itsdevbear/bolaris/cosmos/x/beacon/keeper"
)

// Service implements the baseapp.TxSelector interface.
type Service struct {
	execution.EngineCaller
	logger log.Logger
	ek     *beaconkeeper.Keeper
	bk     *blockchain.Service
	// TODO: handle this better with a proper LRU cache.
	cachedPayload *pb.PayloadIDBytes
}

// New produces a cosmos forker from a geth forker.
func New(gm execution.EngineCaller, bk *blockchain.Service, ek *beaconkeeper.Keeper, logger log.Logger) *Service {
	return &Service{
		EngineCaller: gm,
		logger:       logger,
		ek:           ek,
		bk:           bk,
	}
}

func (m *Service) BuildBlockV2(ctx sdk.Context) (interfaces.ExecutionData, error) {
	// Reference: https://hackmd.io/@danielrachi/engine_api#Block-Building
	// builder := (&execution.Builder{EngineCaller: m.EngineCaller.(*execution.Service)})
	m.logger.Info("entering build-block-v2")
	defer m.logger.Info("exiting build-block-v2")
	var payloadID *pb.PayloadIDBytes

	// EDGE CASE
	if m.cachedPayload == nil {
		m.logger.Info("No cached payload found, building new payload")
		// Trigger the execution client to begin building the block, and update
		// the proposers forkchoice state accordingly. By setting the HeadBlockHash
		// to the last finalized block that we have seen, we are telling the execution client
		// to begin building a block on top of that block.

		// TODO: SHOULD THIS BE LATEST OR FINALIZED????
		// IN THEORY LATEST MEANS THAT THE PROPOSER CAN APPEND ARBITRARY BLOCKS
		// AND SET THE CANONICAL CHAIN TO WHATEVER IT WANTS (i.e) could
		// apply a deep deep reorg?
		// On the flip side, if we force LatestFinalizedBlock(), we can
		// at the Consensus Layer ensure the reorg will always be 1 deep.
		b, err := m.EngineCaller.LatestFinalizedBlock(ctx)
		if err != nil {
			m.logger.Error("failed to get block number", "err", err)
		}

		attrs, err := m.bk.GetPayloadAttributes(
			ctx, uint64(ctx.BlockHeight()), uint64(ctx.BlockTime().Unix()),
		)
		if err != nil {
			m.logger.Error("failed to get payload attributes", "err", err)
		}

		// We start by setting the head of our execution client to the
		// latest block that we have seen.
		zeroHash := [32]byte{}
		fc := &pb.ForkchoiceState{
			HeadBlockHash:      b.Hash.Bytes(),
			SafeBlockHash:      zeroHash[:],
			FinalizedBlockHash: zeroHash[:],
		}
		payloadID, _, err = m.EngineCaller.ForkchoiceUpdated(ctx, fc, attrs)
		if err != nil {
			m.logger.Error("failed to get forkchoice updated", "err", err)
			return nil, err
		}

		// TODO: this should be something that is 80% of proposal timeout, or so?
		// TODO: maybe this should be some sort of event that we wait for?
		// But the TLDR is that we need to wait for the execution client to
		// build the payload before we can include it in the beacon block.
		time.Sleep(8000 * time.Millisecond) //nolint:gomnd // temp.
	} else {
		payloadID = m.cachedPayload
	}

	builtPayload, err := m.bk.ProposeNewFinalBlock(ctx, ctx.HeaderInfo(), payloadID)
	if err != nil {
		return nil, err
	}

	return builtPayload, nil
}

func (m *Service) ValidateBlock(ctx sdk.Context, builtPayload interfaces.ExecutionData) error {
	payload, err := m.bk.ProcessBlock(ctx, ctx.HeaderInfo(), builtPayload)
	if err != nil {
		return err
	}
	m.cachedPayload = payload
	return nil
}
