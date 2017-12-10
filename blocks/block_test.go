package blocks

import (
	"testing"
)


func TestVersionBits(t *testing.T) {
	flags := int32(0x42)

	out := VersionBits(flags)

	expected := int32(0x20000042)

	if out != expected {
		t.Errorf("VersionBits did not match expected:\n Expected: %08x\n Actual: %08x", expected, out)
	}
}
