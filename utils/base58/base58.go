// Copyright 2020 The TurtleCoin Developers
//
// Please see the included LICENSE file for more information.

package base58

import "fmt"

// Encode returns the base58 encoding of src
func Encode(src []byte) string {
	binsz := len(src)
	var i, j, zcount, high int
	var carry uint32

	for zcount < binsz && src[zcount] == 0 {
		zcount++
	}

	size := ((binsz-zcount)*138/100 + 1)

	// allocate one big buffer up front
	buf := make([]byte, size*2+zcount)

	// use the second half for the temporary buffer
	tmp := buf[size+zcount:]

	high = size - 1
	for i = zcount; i < binsz; i++ {
		j = size - 1
		for carry = uint32(src[i]); j > high || carry != 0; j-- {
			carry = carry + 256*uint32(tmp[j])
			tmp[j] = byte(carry % 58)
			carry /= 58
		}
		high = j
	}

	for j = 0; j < size && tmp[j] == 0; j++ {
	}

	// Use the first half for the result
	b58 := buf[:size-j+zcount]

	if zcount != 0 {
		for i = 0; i < zcount; i++ {
			b58[i] = alphabetIdx0
		}
	}

	for i = zcount; j < size; i++ {
		b58[i] = alphabet[tmp[j]]
		j++
	}

	return string(b58)
}

// Decode returns the bytes represented by the base58 string s
func Decode(s string) ([]byte, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("zero length string")
	}

	var (
		t, c   uint64
		zmask  uint32
		zcount int

		b58u  = []rune(s)
		b58sz = len(b58u)

		outisz    = (b58sz + 3) >> 2
		binu      = make([]byte, (b58sz+3)*3)
		bytesleft = b58sz & 3
	)

	if bytesleft > 0 {
		zmask = 0xffffffff << uint32(bytesleft*8)
	} else {
		bytesleft = 4
	}

	var outi = make([]uint32, outisz)

	for i := 0; i < b58sz && b58u[i] == alphabetIdx0; i++ {
		zcount++
	}

	for _, r := range b58u {
		if r > 127 {
			return nil, fmt.Errorf("high-bit set on invalid digit")
		}
		if decodeMap[r] == 255 {
			return nil, fmt.Errorf("invalid base58 digit (%q)", r)
		}

		c = uint64(decodeMap[r])

		for j := outisz - 1; j >= 0; j-- {
			t = uint64(outi[j])*58 + c
			c = (t >> 32) & 0x3f
			outi[j] = uint32(t & 0xffffffff)
		}

		if c > 0 {
			return nil, fmt.Errorf("output number too big (carry to the next int32)")
		}

		if outi[0]&zmask != 0 {
			return nil, fmt.Errorf("output number too big (last int32 filled too far)")
		}
	}

	var j, cnt int
	for j, cnt = 0, 0; j < outisz; j++ {
		for mask := byte(bytesleft-1) * 8; mask <= 0x18; mask, cnt = mask-8, cnt+1 {
			binu[cnt] = byte(outi[j] >> mask)
		}
		if j == 0 {
			bytesleft = 4 // because it could be less than 4 the first time through
		}
	}

	for n, v := range binu {
		if v > 0 {
			start := n - zcount
			if start < 0 {
				start = 0
			}
			return binu[start:cnt], nil
		}
	}
	return binu[:cnt], nil
}
