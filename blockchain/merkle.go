package blockchain

import (
	"bytes"
	"errors"
	"fmt"

	exonumCrypto "github.com/inn4science/exonum-go/crypto"
)

type MerkleTree struct {
	Left       *MerkleTree
	Right      *MerkleTree
	Hash       []byte
	HashString string
	Value      []byte
}

type MerkleProof struct {
	rootHash  exonumCrypto.Hash
	proofNode MerkleTree
	elements  [][]byte
}

func (MerkleProof) New(rootHash exonumCrypto.Hash, proofNode MerkleTree) MerkleProof {
	return MerkleProof{
		rootHash,
		proofNode,
		make([][]byte, 0),
	}
}

func (mp *MerkleProof) CheckProof() ([][]byte, error) {
	actualHash, err := mp.getHash(&mp.proofNode)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(mp.rootHash.GetData(), actualHash) {
		return nil, errors.New("rootHash parameter is not equal to actual hash")
	}

	return mp.elements, nil
}

func (mp *MerkleProof) getHash(mTree *MerkleTree) ([]byte, error) {
	var err error

	if mTree.Value != nil {
		mp.elements = append(mp.elements, mTree.Value)
		hash := exonumCrypto.Hash{}.FromData(mTree.Value)
		return hash.GetData(), nil
	}

	if mTree.HashString != "" {
		hash, err := exonumCrypto.Hash{}.FromString(mTree.HashString)
		if err != nil {
			return nil, errors.New("tree element of wrong type is passed. Hexadecimal expected")
		}
		return hash.GetData(), nil
	}

	if mTree.Hash != nil {
		return mTree.Hash, nil
	}

	if mTree.Left == nil {
		return nil, errors.New("left node is missed")
	}

	hashLeft, err := mp.getHash(mTree.Left)
	if err != nil {
		fmt.Println("error", err)
	}

	hashRight, err := mp.getHash(mTree.Right)
	if err != nil {
		fmt.Println("error", err)
	}

	hash := exonumCrypto.Hash{}.FromData(append(hashLeft, hashRight...))
	return hash.GetData(), nil
}
