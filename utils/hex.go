// Copyright 2020 The TurtleCoin Developers
//
// Please see the included LICENSE file for more information.

package utils

import "encoding/hex"

// DecodeHex converts a hex-encoded string into a raw byte string.
func DecodeHex(s string) ([]byte, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return b, nil
}
