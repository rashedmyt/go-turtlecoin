// Copyright 2020 The TurtleCoin Developers
//
// Please see the included LICENSE file for more information.

package mnemonics

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"strings"

	"github.com/turtlecoin/go-turtlecoin/types"
)

// MnemonicToPrivateKey converts a string of mnemonic words
// to a private key
func MnemonicToPrivateKey(seed string) (privateKey types.PrivateKey, err error) {
	words := strings.Split(seed, " ")
	wLen := len(words)

	if wLen < 25 {
		var wordPlural string

		if wLen == 1 {
			wordPlural = "word"
		} else {
			wordPlural = "words"
		}

		err = fmt.Errorf("The mnemonic seed given is of wrong length. It"+
			" should be 25 words long, but it is %d %s long.", wLen, wordPlural)
		return privateKey, err
	}

	wIndices, err := getWordIndices(words)
	if err != nil {
		return privateKey, err
	}

	if !hasValidChecksum(words) {
		return privateKey, errors.New("Given mnemonic seed has invalid checksum")
	}

	wlLen := len(english)

	for i := 0; i < (wLen-1)/3; i++ {
		w1 := wIndices[i*3]
		w2 := wIndices[i*3+1]
		w3 := wIndices[i*3+2]

		val := w1 + wlLen*(((wlLen-w1)+w2)%wlLen) + wlLen*wlLen*(((wlLen-w2)+w3)%wlLen)

		if val%wlLen != w1 {
			return privateKey, errors.New("Invalid mnemonic")
		}

		binary.LittleEndian.PutUint32(privateKey[i*4:], uint32(val))
	}

	return
}

// PrivateKeyToMnemonic converts a private key to
// a string of mnemonic words
func PrivateKeyToMnemonic(privateKey types.PrivateKey) string {
	var words []string
	wlLen := uint32(len(english))

	for i := 0; i < 31; i += 4 {
		val := binary.LittleEndian.Uint32(privateKey[i:])

		w1 := val % wlLen
		w2 := ((val / wlLen) + w1) % wlLen
		w3 := (((val / wlLen) / wlLen) + w2) % wlLen

		words = append(words, english[w1])
		words = append(words, english[w2])
		words = append(words, english[w3])
	}

	words = append(words, getChecksumWord(words))

	return strings.Join(words, " ")
}

func hasValidChecksum(words []string) bool {
	return words[len(words)-1] == getChecksumWord(words[:len(words)-1])
}

func getChecksumWord(words []string) string {
	var trimmed string

	for _, v := range words {
		trimmed += v[0:3]
	}

	hash := crc32.ChecksumIEEE([]byte(trimmed))

	return words[hash%uint32(len(words))]
}

func getWordIndices(words []string) ([]int, error) {

	// Create a map to retrieve indices
	indices := map[string]int{}

	maxLen := len(english)

	for i := 0; i < maxLen; i++ {
		indices[english[i]] = i
	}

	res := make([]int, len(words))

	for i, v := range words {
		if index, ok := indices[v]; ok {
			res[i] = index
		} else {
			err := fmt.Errorf("The mnemonic seed given has a word that is not"+
				" present in the english word list (%s).", v)
			return nil, err
		}
	}

	return res, nil
}
