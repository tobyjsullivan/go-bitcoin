package blocks

import (
	"encoding/binary"
	"crypto/sha256"
)

type Header struct {
	Version        int32
	HashPrevBlock  [32]byte
	HashMerkleRoot [32]byte
	Time           uint32
	Bits           uint32
	Nonce          uint32
	vtx            []Transaction
}

func (h *Header) Hash() [32]byte {
	var hashInput [80]byte

	buf := make([]byte, binary.MaxVarintLen32)
	binary.LittleEndian.PutUint32(buf, uint32(h.Version))
	copy(hashInput[0:4], buf[:4])
	buf = make([]byte, 32)
	reverseEndian(buf, h.HashPrevBlock[:])
	copy(hashInput[4:36], buf)
	buf = make([]byte, 32)
	reverseEndian(buf, h.HashMerkleRoot[:])
	copy(hashInput[36:68], buf)
	buf = make([]byte, binary.MaxVarintLen32)
	binary.LittleEndian.PutUint32(buf, h.Time)
	copy(hashInput[68:72], buf[:4])
	buf = make([]byte, binary.MaxVarintLen32)
	binary.LittleEndian.PutUint32(buf, h.Bits)
	copy(hashInput[72:76], buf[:4])
	buf = make([]byte, binary.MaxVarintLen32)
	binary.LittleEndian.PutUint32(buf, h.Nonce)
	copy(hashInput[76:80], buf[:4])

	hash1 := sha256.Sum256(hashInput[:])

	hash2 := sha256.Sum256(hash1[:])
	out := [32]byte{}
	reverseEndian(out[:], hash2[:])

	return out
}
