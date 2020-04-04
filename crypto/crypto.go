// Copyright 2020 The TurtleCoin Developers
//
// Please see the included LICENSE file for more information.

package crypto

import (
	"github.com/turtlecoin/go-turtlecoin/crypto/edwards25519"
	"github.com/turtlecoin/go-turtlecoin/types"
)

// GenerateKeys generates private and
// public key pair
func GenerateKeys() (privKey types.PrivateKey, pubKey types.PublicKey, err error) {

	privateKey, err := randomScalar()
	if err != nil {
		return privKey, pubKey, err
	}

	publicKey := NewKeyFromSeed(privateKey[:])

	return privateKey, publicKey, nil
}

// GenerateDeterministicKeys generates a
// deterministic private and public key pair
// from the given seed
func GenerateDeterministicKeys(seed types.PrivateKey) (types.PrivateKey, types.PublicKey) {
	var privateKey [32]byte
	var hBytes [32]byte

	copy(hBytes[:], seed[:])
	edwards25519.ScReduce32(&privateKey, &hBytes)

	publicKey := NewKeyFromSeed(privateKey[:])

	return privateKey, publicKey
}

// NewKeyFromSeed generates the public key
// corresponding to the given seed.
func NewKeyFromSeed(seed []byte) types.PublicKey {
	var A edwards25519.ExtendedGroupElement
	var hBytes [32]byte
	copy(hBytes[:], seed[:])
	edwards25519.GeScalarMultBase(&A, &hBytes)
	var publicKeyBytes [32]byte
	A.ToBytes(&publicKeyBytes)

	return publicKeyBytes
}
