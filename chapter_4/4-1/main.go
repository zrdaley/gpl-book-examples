package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("y"))

	fmt.Printf("SHA 256 x: %v\n", c1)
	fmt.Printf("SHA 256 x: %v\n", c2)
	fmt.Printf("Number of different bits: %v\n", BitCompare(c1, c2))
}

func BitCompare(x [32]byte, y [32]byte) int {
	diff_bits := 0
	for i := range x {
		for j := 1; i <= 8; i++ {
			if x[i]&byte(j*8) != y[i]&byte(j*8) {
				diff_bits++
			}
		}
	}
	return diff_bits
}
