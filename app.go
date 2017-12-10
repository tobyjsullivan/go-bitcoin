package main

import (
	"github.com/tobyjsullivan/go-bitcoin/blocks"
	"encoding/hex"
	"fmt"
)

const (
	strHashPrevBlock = "000000000000000000434a86dc49d3cf5349e584b6cbbdf292c68164db3af950"
	strHashMerkleRoot =  "a0775e03f0ffe2a6f609ce5e6ac731035f1c716f020e56ae33e3e72215f1c430"
	time = 1512886999
	bits = 402698477
	nonce = 3288773616
)

var (
	version = blocks.VersionBits(0)
)

func main() {
	println("Building header...")
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

	head := blocks.Header{
		Version: version,
		HashPrevBlock: hashPrevBlock,
		HashMerkleRoot: hashMerkleRoot,
		Time: time,
		Bits: bits,
		Nonce: nonce,
	}

	println("Computing hash of header")
	hash := head.Hash()

	println("Result:", fmt.Sprintf("%032x", hash))
}

