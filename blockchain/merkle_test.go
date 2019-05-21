package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"

	exonumCrypto "github.com/inn4science/exonum-go/crypto"
)

func TestMerkleProof_CheckProof(t *testing.T) {
	merkleTree := MerkleTree{
		Left: &MerkleTree{
			Left: &MerkleTree{
				HashString: "0438082601f8b38ae010a621a48f4b4cd021c4e6e69219e3c2d8abab482039e9",
			},
			Right: &MerkleTree{
				Left: &MerkleTree{
					HashString: "53d8e8a89be4b96326ff43c3cf6ad4a25cf2e1ec44dbd75937c3cca9d632abaf",
				},
				Right: &MerkleTree{
					Left: &MerkleTree{
						HashString: "91f3479b74c954440ec98bc1dac97fa8a3aac38ceeeebb3996b3ee7aa2e498c8",
					},
					Right: &MerkleTree{
						Left: &MerkleTree{
							Left: &MerkleTree{
								Left: &MerkleTree{
									HashString: "77aed0fb2f30b04e29c006a137d9d879c7fee98d9757fde7580bccd72dce8946",
								},
								Right: &MerkleTree{
									Value: []byte{153, 209, 189, 13, 222, 26, 107, 28, 238, 121},
								},
							},
							Right: &MerkleTree{
								Left: &MerkleTree{
									Value: []byte{98, 142, 223, 244, 216, 184, 203, 213, 158, 53},
								},
								Right: &MerkleTree{
									Value: []byte{152, 213, 96, 73, 235, 62, 222, 64, 239, 47},
								},
							},
						},
						Right: &MerkleTree{
							Left: &MerkleTree{
								Left: &MerkleTree{
									Value: []byte{43, 93, 185, 75, 181, 229, 155, 34, 13, 95},
								},
								Right: &MerkleTree{
									Value: []byte{30, 13, 8, 60, 72, 173, 0, 135, 185, 216},
								},
							},
							Right: &MerkleTree{
								Left: &MerkleTree{
									Value: []byte{210, 127, 166, 231, 147, 251, 53, 14, 79, 139},
								},
								Right: &MerkleTree{
									HashString: "cf0710f238e771f3e73d08ad1eb5bc8e6cbf88afa343ea94954f95005c4553f0",
								},
							},
						},
					},
				},
			},
		},
		Right: &MerkleTree{
			HashString: "b267fa0930dede7557b805fe643a3ce8ebe4434e366924df1d622785365cf0fc",
		},
	}
	rootHash, err := exonumCrypto.Hash{}.FromString("280a704efafae9410d7b07140bb130e4995eeb381ba90939b4eaefcaf740ca25")
	assert.NoError(t, err)

	merkleProof := MerkleProof{}.New(rootHash, merkleTree)
	elements, err := merkleProof.CheckProof()
	assert.NoError(t, err)

	expectedElements := [][]byte{
		{153, 209, 189, 13, 222, 26, 107, 28, 238, 121},
		{98, 142, 223, 244, 216, 184, 203, 213, 158, 53},
		{152, 213, 96, 73, 235, 62, 222, 64, 239, 47},
		{43, 93, 185, 75, 181, 229, 155, 34, 13, 95},
		{30, 13, 8, 60, 72, 173, 0, 135, 185, 216},
		{210, 127, 166, 231, 147, 251, 53, 14, 79, 139},
	}

	assert.Equal(t, expectedElements, elements)
}
