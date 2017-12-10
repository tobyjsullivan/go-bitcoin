package blocks

import (
	"crypto/sha256"
	"encoding/binary"
)

//type Header struct {
//	Version        int32
//	HashPrevBlock  [32]byte
//	HashMerkleRoot [32]byte
//	Time           uint32
//	Bits           uint32
//	Nonce          uint32
//	vtx            []Transaction
//}

type Header [80]byte

func NewHeader(
	nVersion int32,
	hashPrevBlock [32]byte,
	hashMerkleRoot [32]byte,
	nTime uint32,
	nBits uint32,
	nNonce uint32,
) Header {
	var hashInput [80]byte

	binary.LittleEndian.PutUint32(hashInput[0:4], uint32(nVersion))
	reverseEndian(hashInput[4:36], hashPrevBlock[:])
	reverseEndian(hashInput[36:68], hashMerkleRoot[:])
	binary.LittleEndian.PutUint32(hashInput[68:72], nTime)
	binary.LittleEndian.PutUint32(hashInput[72:76], nBits)
	binary.LittleEndian.PutUint32(hashInput[76:80], nNonce)

	return Header(hashInput)
}

func (h *Header) SetNonce(nonce uint32) {
	binary.LittleEndian.PutUint32(h[76:80], nonce)
}

func (h *Header) Hash() [32]byte {
	hash1 := sha256.Sum256(h[:])
	hash2 := sha256.Sum256(hash1[:])

	out := [32]byte{}
	reverseEndian(out[:], hash2[:])
	return out
}
