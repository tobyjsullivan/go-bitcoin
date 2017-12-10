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

	buf := make([]byte, binary.MaxVarintLen32)
	binary.LittleEndian.PutUint32(buf, uint32(nVersion))
	copy(hashInput[0:4], buf[:4])
	buf = make([]byte, 32)
	reverseEndian(buf, hashPrevBlock[:])
	copy(hashInput[4:36], buf)
	buf = make([]byte, 32)
	reverseEndian(buf, hashMerkleRoot[:])
	copy(hashInput[36:68], buf)
	buf = make([]byte, binary.MaxVarintLen32)
	binary.LittleEndian.PutUint32(buf, nTime)
	copy(hashInput[68:72], buf[:4])
	buf = make([]byte, binary.MaxVarintLen32)
	binary.LittleEndian.PutUint32(buf, nBits)
	copy(hashInput[72:76], buf[:4])
	buf = make([]byte, binary.MaxVarintLen32)
	binary.LittleEndian.PutUint32(buf, nNonce)
	copy(hashInput[76:80], buf[:4])

	return Header(hashInput)
}

func (h *Header) SetNonce(nonce uint32) {

	buf := make([]byte, binary.MaxVarintLen32)
	binary.LittleEndian.PutUint32(buf, nonce)
	copy(h[76:80], buf[:4])
}

func (h *Header) Hash() [32]byte {
	hash1 := sha256.Sum256(h[:])

	hash2 := sha256.Sum256(hash1[:])
	out := [32]byte{}
	reverseEndian(out[:], hash2[:])

	return out
}
