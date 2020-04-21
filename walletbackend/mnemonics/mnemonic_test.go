// Copyright 2020 The TurtleCoin Developers
//
// Please see the included LICENSE file for more information.

package mnemonics

import (
	"encoding/hex"
	"testing"

	"github.com/turtlecoin/go-turtlecoin/types"
)

var tests = []struct {
	seed string
	key  string
}{
	{
		"cement wallets dwarf karate directed surfer insult logic desk spout pipeline inroads elapse mobile syringe launching pool ravine bomb eagle happens poetry toilet touchy syringe",
		"68367745ff0af79e7b1b24b145e398c215de1c4a96f76f079fc2af22b2349b01",
	},
	{
		"semifinal ailments punch federal army faked bobsled anvil incur ticket having bias bamboo bagpipe lair arena utility affair wrong total lair directed taunts terminal bagpipe",
		"8bb480a81cdf1e37f253055cc4a02fbd90de7d624ec18119d18b349b1e1c2602",
	},
	{
		"sulking testing unbending dexterity mobile neon lazy argue wobbly mumble pyramid legion eternal pool eggs woven upcoming southern jukebox hectare females emulate weekday wolf emulate",
		"b6c9d20bef4d250786de94ea5b3f4fcb6fbb3391be6d6fe587641beb953f8604",
	},
}

func TestMnemonicToPrivateKey(t *testing.T) {
	for _, test := range tests {
		key, err := MnemonicToPrivateKey(test.seed)
		if err != nil {
			t.Error(err)
		}

		if hex.EncodeToString(key[:]) != test.key {
			t.Errorf("Invalid conversion from mnemonic to private key: got %s want %s", key, test.key)
		}
	}
}

func TestPrivateKeyToMnemonic(t *testing.T) {
	for _, test := range tests {
		var testKey types.PrivateKey
		keyBytes, _ := hex.DecodeString(test.key)
		copy(testKey[:], keyBytes)
		seed := PrivateKeyToMnemonic(testKey)
		if seed != test.seed {
			t.Errorf("Invalid conversion from private key to mnemonic: got %s want %s", seed, test.seed)
		}
	}
}
