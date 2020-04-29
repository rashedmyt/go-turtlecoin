// Copyright 2020 The TurtleCoin Developers
//
// Please see the included LICENSE file for more information.

package utils

import (
	"github.com/turtlecoin/go-turtlecoin/crypto/sha3"
	"github.com/turtlecoin/go-turtlecoin/types"
	"github.com/turtlecoin/go-turtlecoin/utils/base58"
)

const addressPrefix int64 = 3914525

func AddressFromKeys(pubSpendKey types.PublicKey, pubViewKey types.PublicKey) (result string) {
	combined := PackPrefixAsByteList(addressPrefix)
	combined = append(combined, pubSpendKey[:]...)
	combined = append(combined, pubViewKey[:]...)

	checksum := GetAddressChecksum(combined)

	combined = append(combined, checksum...)

	for i := 0; i < len(combined)/8; i++ {
		result += base58.Encode(combined[i*8 : (i+1)*8])
	}

	if len(combined)%8 > 0 {
		result += base58.Encode(combined[(len(combined)/8)*8:])
	}

	return
}

func GetAddressChecksum(addressInBytes []byte) []byte {
	res := sha3.Keccak(addressInBytes)
	return res[:4]
}

func PackPrefixAsByteList(prefix int64) []byte {
	var res []byte

	for prefix >= 0x80 {
		res = append(res, byte(prefix&0x7f|0x80))
		prefix >>= 7
	}

	res = append(res, byte(prefix))

	return res
}
