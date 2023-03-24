package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// read in transactions from file
	txnsHex, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	txnsStr := strings.TrimSuffix(string(txnsHex), "\n") // remove newline characters
	txns := make([][]byte, len(txnsStr)/64)
	for i := 0; i < len(txns); i++ {
		txns[i], err = hex.DecodeString(txnsStr[i*64 : (i+1)*64])
		if err != nil {
			panic(err)
		}
	}

	// calculate merkle tree
	root := merkleRoot(txns)

	fmt.Printf("Merkle root: %x\n", root)
}

func merkleRoot(txns [][]byte) []byte {
	if len(txns) == 0 {
		return nil
	}
	if len(txns) == 1 {
		return txns[0]
	}

	// recursively calculate left and right merkle roots
	mid := len(txns) / 2
	left := merkleRoot(txns[:mid])
	right := merkleRoot(txns[mid:])

	// concatenate and hash left and right merkle roots
	combined := append(left, right...)
	hash := sha256.Sum256(combined)

	return hash[:]
}
