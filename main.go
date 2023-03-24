package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func main() {
	// Read transactions from file
	txns, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Convert transactions from hex to byte slices
	var txnData [][]byte
	for _, txn := range bytes.Split(txns, []byte("\n")) {
		if len(txn) > 0 {
			txnData = append(txnData, txn)
		}
	}

	// Construct merkle tree
	tree := buildMerkleTree(txnData)

	// Output merkle root
	fmt.Println(hex.EncodeToString(tree[0]))
}

// buildMerkleTree constructs a merkle tree from the given data and returns the root hash.
func buildMerkleTree(data [][]byte) [][]byte {
	// Create initial leaf nodes
	var leaves [][]byte
	for _, d := range data {
		hash := sha256.Sum256(d)
		leaves = append(leaves, hash[:])
	}

	// Iterate until we reach the root node
	for len(leaves) > 1 {
		// If there's an odd number of leaves, duplicate the last one
		if len(leaves)%2 == 1 {
			leaves = append(leaves, leaves[len(leaves)-1])
		}

		// Create parent nodes
		var parents [][]byte
		for i := 0; i < len(leaves); i += 2 {
			hash := sha256.Sum256(append(leaves[i], leaves[i+1]...))
			parents = append(parents, hash[:])
		}

		// Set parents as new leaves
		leaves = parents
	}

	return leaves
}
