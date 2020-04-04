// Copyright 2020 The TurtleCoin Developers
//
// Please see the included LICENSE file for more information.

package crypto

import (
	crand "crypto/rand"
	"io"

	"github.com/turtlecoin/go-turtlecoin/crypto/edwards25519"
	"github.com/turtlecoin/go-turtlecoin/types"
)

func randomScalar() (privKey types.PrivateKey, err error) {
	rand := crand.Reader

	var seed [64]byte
	if _, err := io.ReadFull(rand, seed[:]); err != nil {
		return privKey, err
	}

	var privateKey [32]byte

	edwards25519.ScReduce(&privateKey, &seed)

	return privateKey, nil
}
