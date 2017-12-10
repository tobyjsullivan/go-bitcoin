package blocks

import "errors"

func reverseEndian(dst, src []byte) error {
	dl := len(dst)
	sl := len(src)
	if dl != sl {
		return errors.New("reverseEndian: dst and src len must match")
	}

	for i := 0; i < sl; i++ {
		dst[dl-(i+1)] = src[i]
	}

	return nil
}
