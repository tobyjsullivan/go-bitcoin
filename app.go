package main

import (
	"encoding/hex"
	"fmt"
	"github.com/tobyjsullivan/go-bitcoin/blocks"
	time2 "time"
)

const (
	strHashPrevBlock  = "000000000000000000434a86dc49d3cf5349e584b6cbbdf292c68164db3af950"
	strHashMerkleRoot = "a0775e03f0ffe2a6f609ce5e6ac731035f1c716f020e56ae33e3e72215f1c430"
	time              = 1512886999
	bits              = 402698477
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

	nextNonce := make(chan uint32, 200)
	done := make(chan bool)
	go func(nextNonce chan uint32) {
		for nonce := uint32(0); ; nonce++ {
			nextNonce <- nonce
		}
		done <- true
	}(nextNonce)

	numParallel := 8

	start := time2.Now()
	var best [32]byte
	var set bool
	for i := 0; i < numParallel; i++ {
		go func(nextNonce chan uint32, start time2.Time) {
			head := blocks.NewHeader(version, hashPrevBlock, hashMerkleRoot, time, bits, 0)

			for nonce := range nextNonce {
				head.SetNonce(nonce)
				hash := head.Hash()

				if !set {
					set = true
					best = hash
				}

				for i := 0; i < len(hash); i++ {
					if best[i] < hash[i] {
						break
					}
					if hash[i] < best[i] {
						best = hash
						break
					}
				}

				if nonce%5000000 == 0 {
					now := time2.Now()
					seconds := now.Sub(start).Seconds()

					rate := float64(nonce) / seconds
					println(fmt.Sprintf("Nonce: %d\nHash: %032x\nBest: %032x\nRate: %.00f/s", nonce, hash, best, rate))
				}
			}
		}(nextNonce, start)
	}

	<-done
}
