package blocks

import (
	"fmt"
)

const (
	BlockVersion = 0x2
)

type Transaction struct {
	Version int32

}

type Input struct {
	Hash [32]byte
	N uint32
	ScriptSig []byte
	Sequence uint32
}

type Output struct {
	Value int64
	ScriptPubkey []byte
}

func VersionBits(softForkFlags int32) int32 {
	if uint32(softForkFlags) & uint32(0xf000000) != 0 {
		panic(fmt.Sprintf("VersionBits: invalid softForkFlags: %08x", softForkFlags))
	}

	return (BlockVersion << 28) | softForkFlags
}

