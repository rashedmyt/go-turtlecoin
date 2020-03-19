// Copyright 2020 The TurtleCoin Developers
//
// Please see the included LICENSE file for more information.

package sha3

import "encoding/binary"

// Keccak1600 computes the 200 byte keccak
// hash of given input.
func Keccak1600(input []byte) []byte {
	return keccak(input, 200)
}

// Keccak computes the 32 byte keccak
// hash of given input
func Keccak(input []byte) []byte {
	return keccak(input, 32)
}

// KeccakF permutes the state for fixed
// number of rounds with corresponding
// round constant.
func KeccakF(st []uint64) {
	var t uint64
	var bc [5]uint64

	for round := 0; round < keccakRounds; round++ {

		// Theta
		for i := 0; i < 5; i++ {
			bc[i] = st[i] ^ st[i+5] ^ st[i+10] ^ st[i+15] ^ st[i+20]
		}

		for i := 0; i < 5; i++ {
			t = bc[(i+4)%5] ^ rotl64(bc[(i+1)%5], 1)
			for j := 0; j < 25; j += 5 {
				st[i+j] ^= t
			}
		}

		// Rho Pi
		t = st[1]

		for i := 0; i < 24; i++ {
			j := piln[i]
			bc[0] = st[j]
			st[j] = rotl64(t, uint64(rotc[i]))
			t = bc[0]
		}

		// Chi
		for j := 0; j < 25; j += 5 {
			for i := 0; i < 5; i++ {
				bc[i] = st[i+j]
			}

			for i := 0; i < 5; i++ {
				st[i+j] ^= (^bc[(i+1)%5]) & bc[(i+2)%5]
			}
		}

		// Iota
		st[0] ^= rc[round]
	}
}

func keccak(input []byte, outputSize int) []byte {

	st := make([]uint64, 25)

	rsiz := hashDataArea

	if outputSize != 200 {
		rsiz = 200 - 2*outputSize
	}

	rsizw := rsiz / 8

	inputLength := len(input)
	offset := 0

	for ; inputLength >= rsiz; inputLength, offset = inputLength-rsiz, offset+rsiz {
		for i := 0; i < rsizw; i++ {
			st[i] ^= binary.LittleEndian.Uint64(input[offset+i*8:])
		}

		KeccakF(st)
	}

	temp := make([]byte, 144)
	for i := 0; i < inputLength; i++ {
		temp[i] = input[offset+i]
	}

	temp[inputLength] = 1
	inputLength++

	for i := 0; i < rsiz-inputLength; i++ {
		temp[inputLength+i] = 0
	}

	temp[rsiz-1] |= 0x80

	for i := 0; i < rsizw; i++ {
		st[i] ^= binary.LittleEndian.Uint64(temp[8*i:])
	}

	KeccakF(st)

	output := make([]byte, outputSize)

	for i := 0; i < outputSize; i += 8 {
		binary.LittleEndian.PutUint64(output[i:], st[i/8])
	}

	return output
}

func rotl64(x, y uint64) uint64 {
	return x<<y | x>>(64-y)
}
