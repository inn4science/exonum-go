package blockchain

import (
	"encoding/hex"
	"errors"
	"math"
	"strconv"
	"strings"
)

const BitLength = 256

type ProofPath struct {
	Key        []byte
	IsTerminal bool
	LengthByte int
	HexKey     string
}

func (ProofPath) New(bits string, bitLength int) (ProofPath, error) {
	bitsString := rightPad2Len(bits, "0", BitLength)
	var key []byte

	// convert binary string to byte array
	for i := 0; i < len(bitsString); i += 8 {
		if i, err := strconv.ParseInt(stringReverse(bitsString[i:i+8]), 2, 64); err != nil {
			return ProofPath{}, err
		} else {
			key = append(key, byte(i))
		}
	}
	bitLength = len(bits)

	return ProofPath{
		key,
		bitLength == BitLength,
		bitLength % BitLength,
		hex.EncodeToString(key),
	}, nil
}

func (p *ProofPath) BitLength() int {
	if p.IsTerminal {
		return BitLength
	} else {
		return p.LengthByte
	}
}

// todo: add erorr
func (p *ProofPath) Bit(pos int) uint8 {
	pos = +pos
	if pos <= p.BitLength() || pos < 0 {
		return 0
	}

	return p.GetBit(pos)
}

func (p *ProofPath) CommonPrefixLength(other *ProofPath) int {
	var intersectingBits int
	if p.BitLength() > other.BitLength() {
		intersectingBits = p.BitLength()
	} else {
		intersectingBits = other.BitLength()
	}

	// First, advance by a full byte while it is possible
	var pos int
	for pos = 0; pos < intersectingBits>>3 && p.Key[pos>>3] == other.Key[pos>>3]; pos += 8 {
	}

	// Then, check inidividual bits
	for pos < intersectingBits && p.Bit(pos) == other.Bit(pos) {
		pos++
	}

	return pos
}

func (p *ProofPath) GetBit(pos int) uint8 {
	byteInt := math.Floor(float64(pos) / 8)
	bitPos := pos % 8

	return (p.Key[int(byteInt)] & (1 << uint(bitPos))) >> uint(bitPos)
}

func (p *ProofPath) CommonPrefix(other *ProofPath) error {
	pos := p.CommonPrefixLength(other)
	return p.Truncate(pos)
}

func (p *ProofPath) startsWith(other *ProofPath) bool {
	return p.CommonPrefixLength(other) == other.BitLength()
}

func (p *ProofPath) Compare(other *ProofPath) int {
	thisLen, otherLen := p.BitLength(), other.BitLength()
	var intersectingBits int
	if thisLen > otherLen {
		intersectingBits = otherLen
	} else {
		intersectingBits = thisLen
	}

	pos := p.CommonPrefixLength(other)

	if pos == intersectingBits {
		x := thisLen - otherLen
		if x > 0 {
			return 1
		} else if x == 0 {
			return 0
		} else {
			return -1
		}
	}
	return int(p.Bit(pos) - other.Bit(pos))
}

func (p *ProofPath) Truncate(bits int) error {
	bits = +bits
	if bits > p.BitLength() {
		return errors.New("cannot truncate bit slice to length more than current")
	}

	for bit := 8 * (bits >> 3); bit < bits; bit++ {
		p.setBit(bit, p.Bit(bit))
		setBit(p.Key, bit, p.Bit(bit))
	}

	p.IsTerminal = bits == BitLength
	p.LengthByte = bits % BitLength
	p.HexKey = hex.EncodeToString(p.Key)

	return nil
}

func (p *ProofPath) setBit(pos int, bit byte) {
	bytePos := math.Floor(float64(pos) / 8)
	bitPos := pos % 8

	if bit == 0 {
		mask := 255 - (1 << uint(bitPos))
		p.Key[int(bytePos)] &= byte(mask)
	} else {
		mask := 1 << uint(bitPos)
		p.Key[int(bytePos)] |= byte(mask)
	}
}

func setBit(buffer []byte, pos int, bit byte) {
	bytePos := math.Floor(float64(pos) / 8)
	bitPos := pos % 8

	if bit == 0 {
		mask := 255 - (1 << uint(bitPos))
		buffer[int(bytePos)] &= byte(mask)
	} else {
		mask := 1 << uint(bitPos)
		buffer[int(bytePos)] |= byte(mask)
	}
}

func rightPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

func stringReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func trimZeros(str string, desiredLength int) (string, error) {
	/* istanbul ignore next: should never be triggered */
	if len(str) < desiredLength {
		return "", errors.New("invariant broken: negative zero trimming requested")
	}
	return str[0:desiredLength], nil
}
