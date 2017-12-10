package blocks

import (
	"encoding/binary"
	"fmt"
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
	println("Version: ", fmt.Sprintf("%x", buf[:4]))
	copy(hashInput[0:4], buf[:4])
	buf = make([]byte, 32)
	reverseEndian(buf, h.HashPrevBlock[:])
	println("HashPrevBlock: ", fmt.Sprintf("%x", buf))
	copy(hashInput[4:36], buf)
	buf = make([]byte, 32)
	reverseEndian(buf, h.HashMerkleRoot[:])
	println("HashMerkleRoot: ", fmt.Sprintf("%x", buf))
	copy(hashInput[36:68], buf)
	buf = make([]byte, binary.MaxVarintLen32)
	binary.LittleEndian.PutUint32(buf, h.Time)
	println("Time: ", fmt.Sprintf("%x", buf[:4]))
	copy(hashInput[68:72], buf[:4])
	buf = make([]byte, binary.MaxVarintLen32)
	binary.LittleEndian.PutUint32(buf, h.Bits)
	println("Bits: ", fmt.Sprintf("%x", buf[:4]))
	copy(hashInput[72:76], buf[:4])
	buf = make([]byte, binary.MaxVarintLen32)
	binary.LittleEndian.PutUint32(buf, h.Nonce)
	println("Nonce: ", fmt.Sprintf("%x", buf[:4]))
	copy(hashInput[76:80], buf[:4])
	println("Final Hash Input: ", fmt.Sprintf("%x", hashInput[:]))

	hash1 := sha256.Sum256(hashInput[:])

	hash2 := sha256.Sum256(hash1[:])
	out := [32]byte{}
	reverseEndian(out[:], hash2[:])

	return out
}
