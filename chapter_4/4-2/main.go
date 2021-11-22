package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]

	sha_384_flag := flag.Bool("sha384", false, "print as sha384")
	sha_512_flag := flag.Bool("sha512", false, "print as sha512")
	flag.Parse()

	if *sha_384_flag {
		fmt.Printf("SHA 384 input: %v\n", sha512.Sum384([]byte(input)))
	} else if *sha_512_flag {
		fmt.Printf("SHA 512 input: %v\n", sha512.Sum512([]byte(input)))
	} else {
		fmt.Printf("SHA 256 input: %v\n", sha256.Sum256([]byte(input)))
	}
}
