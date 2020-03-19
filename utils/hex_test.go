// Copyright 2020 The TurtleCoin Developers
//
// Please see the included LICENSE file for more information.

package utils

import "testing"

func TestDecodeHex(t *testing.T) {
	// should give error
	_, err := DecodeHex("xcpg")
	if err == nil {
		t.Error("Non hex string is decoded without any errors")
	}

	// should not give error
	_, err = DecodeHex("fb9c")
	if err != nil {
		t.Error("Error while decoding a hex string - " + err.Error())
	}
}
