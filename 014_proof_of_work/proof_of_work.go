package main

import (
	"crypto/sha256"
	"fmt"
)

var format = "%s%d"

func proofOfWork(zeros int, data []byte) string {
	i := uint64(0)
	digest := sha256.Sum256([]byte(fmt.Sprintf(format, data, i)))
	s := fmt.Sprintf("%x", digest)

	for i = 0; countZeros(s) != zeros; i++ {
		digest = sha256.Sum256([]byte(fmt.Sprintf(format, data, i)))
		s = fmt.Sprintf("%x: %d", digest, i)
	}

	return fmt.Sprintf("%x", digest)
}

func countZeros(s string) int {
	i := 0

	for i = 0; i < len(s) && s[i] == '0'; i++ {
	}

	return i
}
