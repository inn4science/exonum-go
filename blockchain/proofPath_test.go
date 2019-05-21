package blockchain

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var binaryStrings = []string{
	"0",
	"1",
	"10",
	"011",
	"10101",
	"1010101010010001",
	"10101010010100010010100101",
	"100000000000000000000001000000000000000000000000000000000000000000000000",
}

func TestProofPath_New(t *testing.T) {
	for _, value := range binaryStrings {
		path, err := ProofPath{}.New(value, BitLength)
		assert.NoError(t, err)

		for i := 0; i < len(value); i++ {
			intValue, err := strconv.ParseInt(value[i:i+1], 2, 8)
			if err != nil {
				assert.NoError(t, err)
			}
			assert.Equal(t, uint8(intValue), path.GetBit(i))
		}
	}
}
