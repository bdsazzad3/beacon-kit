// Code generated by fastssz. DO NOT EDIT.
// Hash: bc5bb5973dfa5cde76b6d1feada124a134b31ec04c6d8d2ca1cbecb88cc2647f
// Version: 0.1.3
package types

import (
	"github.com/berachain/beacon-kit/mod/primitives/pkg/eip4844"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BeaconBlockBodyDenebPlus object
func (b *BeaconBlockBodyDenebPlus) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconBlockBodyDenebPlus object to a target array
func (b *BeaconBlockBodyDenebPlus) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(220)

	// Field (0) 'RandaoReveal'
	dst = append(dst, b.RandaoReveal[:]...)

	// Field (1) 'Eth1Data'
	if b.Eth1Data == nil {
		b.Eth1Data = new(Eth1Data)
	}
	if dst, err = b.Eth1Data.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'Graffiti'
	dst = append(dst, b.Graffiti[:]...)

	// Offset (3) 'Deposits'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Deposits) * 192

	// Offset (4) 'ExecutionPayload'
	dst = ssz.WriteOffset(dst, offset)
	if b.ExecutionPayload == nil {
		b.ExecutionPayload = new(ExecutionPayload)
	}
	offset += b.ExecutionPayload.SizeSSZ()

	// Offset (5) 'Attestations'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Attestations) * 48

	// Offset (6) 'SlashingInfo'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.SlashingInfo) * 16

	// Offset (7) 'BlobKzgCommitments'
	dst = ssz.WriteOffset(dst, offset)

	// Field (3) 'Deposits'
	if size := len(b.Deposits); size > 16 {
		err = ssz.ErrListTooBigFn("BeaconBlockBodyDenebPlus.Deposits", size, 16)
		return
	}
	for ii := 0; ii < len(b.Deposits); ii++ {
		if dst, err = b.Deposits[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (4) 'ExecutionPayload'
	if dst, err = b.ExecutionPayload.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (5) 'Attestations'
	if size := len(b.Attestations); size > 256 {
		err = ssz.ErrListTooBigFn("BeaconBlockBodyDenebPlus.Attestations", size, 256)
		return
	}
	for ii := 0; ii < len(b.Attestations); ii++ {
		if dst, err = b.Attestations[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (6) 'SlashingInfo'
	if size := len(b.SlashingInfo); size > 256 {
		err = ssz.ErrListTooBigFn("BeaconBlockBodyDenebPlus.SlashingInfo", size, 256)
		return
	}
	for ii := 0; ii < len(b.SlashingInfo); ii++ {
		if dst, err = b.SlashingInfo[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (7) 'BlobKzgCommitments'
	if size := len(b.BlobKzgCommitments); size > 16 {
		err = ssz.ErrListTooBigFn("BeaconBlockBodyDenebPlus.BlobKzgCommitments", size, 16)
		return
	}
	for ii := 0; ii < len(b.BlobKzgCommitments); ii++ {
		dst = append(dst, b.BlobKzgCommitments[ii][:]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconBlockBodyDenebPlus object
func (b *BeaconBlockBodyDenebPlus) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 220 {
		return ssz.ErrSize
	}

	tail := buf
	var o3, o4, o5, o6, o7 uint64

	// Field (0) 'RandaoReveal'
	copy(b.RandaoReveal[:], buf[0:96])

	// Field (1) 'Eth1Data'
	if b.Eth1Data == nil {
		b.Eth1Data = new(Eth1Data)
	}
	if err = b.Eth1Data.UnmarshalSSZ(buf[96:168]); err != nil {
		return err
	}

	// Field (2) 'Graffiti'
	copy(b.Graffiti[:], buf[168:200])

	// Offset (3) 'Deposits'
	if o3 = ssz.ReadOffset(buf[200:204]); o3 > size {
		return ssz.ErrOffset
	}

	if o3 < 220 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (4) 'ExecutionPayload'
	if o4 = ssz.ReadOffset(buf[204:208]); o4 > size || o3 > o4 {
		return ssz.ErrOffset
	}

	// Offset (5) 'Attestations'
	if o5 = ssz.ReadOffset(buf[208:212]); o5 > size || o4 > o5 {
		return ssz.ErrOffset
	}

	// Offset (6) 'SlashingInfo'
	if o6 = ssz.ReadOffset(buf[212:216]); o6 > size || o5 > o6 {
		return ssz.ErrOffset
	}

	// Offset (7) 'BlobKzgCommitments'
	if o7 = ssz.ReadOffset(buf[216:220]); o7 > size || o6 > o7 {
		return ssz.ErrOffset
	}

	// Field (3) 'Deposits'
	{
		buf = tail[o3:o4]
		num, err := ssz.DivideInt2(len(buf), 192, 16)
		if err != nil {
			return err
		}
		b.Deposits = make([]*Deposit, num)
		for ii := 0; ii < num; ii++ {
			if b.Deposits[ii] == nil {
				b.Deposits[ii] = new(Deposit)
			}
			if err = b.Deposits[ii].UnmarshalSSZ(buf[ii*192 : (ii+1)*192]); err != nil {
				return err
			}
		}
	}

	// Field (4) 'ExecutionPayload'
	{
		buf = tail[o4:o5]
		if b.ExecutionPayload == nil {
			b.ExecutionPayload = new(ExecutionPayload)
		}
		if err = b.ExecutionPayload.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (5) 'Attestations'
	{
		buf = tail[o5:o6]
		num, err := ssz.DivideInt2(len(buf), 48, 256)
		if err != nil {
			return err
		}
		b.Attestations = make([]*AttestationData, num)
		for ii := 0; ii < num; ii++ {
			if b.Attestations[ii] == nil {
				b.Attestations[ii] = new(AttestationData)
			}
			if err = b.Attestations[ii].UnmarshalSSZ(buf[ii*48 : (ii+1)*48]); err != nil {
				return err
			}
		}
	}

	// Field (6) 'SlashingInfo'
	{
		buf = tail[o6:o7]
		num, err := ssz.DivideInt2(len(buf), 16, 256)
		if err != nil {
			return err
		}
		b.SlashingInfo = make([]*SlashingInfo, num)
		for ii := 0; ii < num; ii++ {
			if b.SlashingInfo[ii] == nil {
				b.SlashingInfo[ii] = new(SlashingInfo)
			}
			if err = b.SlashingInfo[ii].UnmarshalSSZ(buf[ii*16 : (ii+1)*16]); err != nil {
				return err
			}
		}
	}

	// Field (7) 'BlobKzgCommitments'
	{
		buf = tail[o7:]
		num, err := ssz.DivideInt2(len(buf), 48, 16)
		if err != nil {
			return err
		}
		b.BlobKzgCommitments = make([]eip4844.KZGCommitment, num)
		for ii := 0; ii < num; ii++ {
			copy(b.BlobKzgCommitments[ii][:], buf[ii*48:(ii+1)*48])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconBlockBodyDenebPlus object
func (b *BeaconBlockBodyDenebPlus) SizeSSZ() (size int) {
	size = 220

	// Field (3) 'Deposits'
	size += len(b.Deposits) * 192

	// Field (4) 'ExecutionPayload'
	if b.ExecutionPayload == nil {
		b.ExecutionPayload = new(ExecutionPayload)
	}
	size += b.ExecutionPayload.SizeSSZ()

	// Field (5) 'Attestations'
	size += len(b.Attestations) * 48

	// Field (6) 'SlashingInfo'
	size += len(b.SlashingInfo) * 16

	// Field (7) 'BlobKzgCommitments'
	size += len(b.BlobKzgCommitments) * 48

	return
}

// HashTreeRoot ssz hashes the BeaconBlockBodyDenebPlus object
func (b *BeaconBlockBodyDenebPlus) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconBlockBodyDenebPlus object with a hasher
func (b *BeaconBlockBodyDenebPlus) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'RandaoReveal'
	hh.PutBytes(b.RandaoReveal[:])

	// Field (1) 'Eth1Data'
	if b.Eth1Data == nil {
		b.Eth1Data = new(Eth1Data)
	}
	if err = b.Eth1Data.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'Graffiti'
	hh.PutBytes(b.Graffiti[:])

	// Field (3) 'Deposits'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Deposits))
		if num > 16 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.Deposits {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 16)
	}

	// Field (4) 'ExecutionPayload'
	if err = b.ExecutionPayload.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (5) 'Attestations'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Attestations))
		if num > 256 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.Attestations {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 256)
	}

	// Field (6) 'SlashingInfo'
	{
		subIndx := hh.Index()
		num := uint64(len(b.SlashingInfo))
		if num > 256 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.SlashingInfo {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 256)
	}

	// Field (7) 'BlobKzgCommitments'
	{
		if size := len(b.BlobKzgCommitments); size > 16 {
			err = ssz.ErrListTooBigFn("BeaconBlockBodyDenebPlus.BlobKzgCommitments", size, 16)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.BlobKzgCommitments {
			hh.PutBytes(i[:])
		}
		numItems := uint64(len(b.BlobKzgCommitments))
		hh.MerkleizeWithMixin(subIndx, numItems, 16)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BeaconBlockBodyDenebPlus object
func (b *BeaconBlockBodyDenebPlus) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
