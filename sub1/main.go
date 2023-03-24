package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func main() {
	// Read the transactions from the file
	transactions, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Convert the hex-encoded transactions to byte arrays
	var txs [][]byte
	for _, tx := range bytes.Split(transactions, []byte("\n")) {
		if len(tx) > 0 {
			txs = append(txs, hexDecode(string(tx)))
		}
	}

	// Calculate the merkle root
	root := calculateMerkleRoot(txs)
	fmt.Println(hex.EncodeToString(root))
}

// hexDecode decodes a hex-encoded string to a byte array.
func hexDecode(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

// calculateMerkleRoot calculates the merkle root of a set of transactions.
func calculateMerkleRoot(txs [][]byte) []byte {
	if len(txs) == 0 {
		return nil
	}
	if len(txs) == 1 {
		return txs[0]
	}
	if len(txs)%2 == 1 {
		txs = append(txs, txs[len(txs)-1])
	}
	var hashes [][]byte
	for i := 0; i < len(txs); i += 2 {
		hash := sha256.Sum256(append(txs[i], txs[i+1]...))
		hashes = append(hashes, hash[:])
	}
	return calculateMerkleRoot(hashes)
}
