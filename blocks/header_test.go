package blocks

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHeader_Hash(t *testing.T) {
	strHashPrevBlock := "000000000000000000434a86dc49d3cf5349e584b6cbbdf292c68164db3af950"
	strHashMerkleRoot := "a0775e03f0ffe2a6f609ce5e6ac731035f1c716f020e56ae33e3e72215f1c430"
	version := int32(0x20000000)
	time := uint32(1512886999)
	bits := uint32(402698477)
	nonce := uint32(3288773616)

	hashPrevBlockParsed, err := hex.DecodeString(strHashPrevBlock)
	if err != nil {
		panic(err)
	}
	hashMerkleRootParsed, err := hex.DecodeString(strHashMerkleRoot)
	if err != nil {
		panic(err)
	}

	var hashPrevBlock [32]byte
	copy(hashPrevBlock[:], hashPrevBlockParsed)
	var hashMerkleRoot [32]byte
	copy(hashMerkleRoot[:], hashMerkleRootParsed)

	head := NewHeader(version, hashPrevBlock, hashMerkleRoot, time, bits, nonce)

	hash := head.Hash()

	expectedHash := "000000000000000000538e17c5a650c9a9405385b6208a7ea4763737e70e0ff3"
	if fmt.Sprintf("%032x", hash) != expectedHash {
		t.Errorf("Result hash did not match expected.\nExpected: %s\nActual: %032x", expectedHash, hash)
	}

	println("Result:", fmt.Sprintf("%032x", hash))
}
