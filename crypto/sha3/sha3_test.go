// Copyright 2020 The TurtleCoin Developers
//
// Please see the included LICENSE file for more information.

package sha3

import (
	"bytes"
	"testing"

	"github.com/turtlecoin/go-turtlecoin/utils"
)

func TestKeccak(t *testing.T) {
	tests := []struct {
		input []byte
		hash  string
	}{
		{
			[]byte(""),
			"c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470",
		},
		{
			[]byte("The quick brown fox jumps over the lazy dog"),
			"4d741b6f1eb29cb2a9b9911c82f56fa8d73b04959d3d9d222895df6c0b28aa15",
		},
		{
			[]byte("The quick brown fox jumps over the lazy dog."),
			"578951e24efd62a3d63a86f7cd19aaa53c898fe287d2552133220370240b572d",
		},
		{
			[]byte("I'd just like to interject for a moment. What you're referring to as Linux, is in fact, GNU/Linux, or as I've recently taken to calling it, GNU plus Linux. Linux is not an operating system unto itself"),
			"d6a63dc2e3ab16360c1dd26fa4b343af9dde6b4ae275793b1d64eaffdc02f1d9",
		},
	}

	for _, test := range tests {
		got := Keccak(test.input)
		want, _ := utils.DecodeHex(test.hash)
		if !bytes.Equal(got, want) {
			t.Errorf("unexpected hash : got '%x' want '%s'", got, test.hash)
		}
	}
}

func TestKeccak1600(t *testing.T) {
	tests := []struct {
		input []byte
		hash  string
	}{
		{
			[]byte(""),
			"c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4703dbb9a2cd87ca974b9a2b0ec61119bcb5cedf9c0c411221f6141a25f17c60d82d24680abbcbfba815b762b24b751d5b1e85325ba5e6df23c10725bfe986ace3ba2d24535a79f7dbabb153bb0d33c0dfa09cec712ebd7fe3b49a9194e859c82ebff11a645651a5d1b726be100f44641069fab7164e13487fe3609bbeebd88309cbaacb2a7ecb8e8de2145cf1db7623b16916d7210991b576bbe182362cf22fab7d7af9f77f71afea3",
		},
		{
			[]byte("The quick brown fox jumps over the lazy dog"),
			"4d741b6f1eb29cb2a9b9911c82f56fa8d73b04959d3d9d222895df6c0b28aa15d92ae6ccbaccd8a14b02c9877ec141985a0fbe2214e17a69d328ff18dc4a952e2ca82016467aedbf7ed95909eb3d7a4b084657031e7e229afa2ce03fef3801756f8ccf7a3c71236b04a36e6bf0da3316424e538782f4a6d5ef04ba77c55e0e107c9ebeb0978ca595e38d02397d017a1fbcdbd78747195b709417c39f7107b8c5bea4c408bc2b6a7fc1f38709da3ed9ea29e43c7852181be5c98802a871d64574941c948edd7da976",
		},
		{
			[]byte("The quick brown fox jumps over the lazy dog."),
			"578951e24efd62a3d63a86f7cd19aaa53c898fe287d2552133220370240b572d9f1e22e5490be85e3bdea844e6f687d630305e2d64b04b518363b2b350831104d067bb2545279d51d1827548d19d85f7d47e72a4c9b3a52fbb044451d44d23a282a894fa6596d31d7f4dcaa346684424b9967e350fe149e4801176c934954b21b3b12138076367b3189c346b9c12f4defe531438f420e5f0f790785995e648a3395a37a05b222e6d14fb7b299247ccfab3919d8095b1567e4a556b2cb57a45dfeeec273d9a6a99a0",
		},
		{
			[]byte("I'd just like to interject for a moment. What you're referring to as Linux, is in fact, GNU/Linux, or as I've recently taken to calling it, GNU plus Linux. Linux is not an operating system unto itself"),
			"d6a63dc2e3ab16360c1dd26fa4b343af9dde6b4ae275793b1d64eaffdc02f1d9778936236831e1588491709a18b2bda7267f270d6313b80b4a4baf40f0305e03605db78f0e10ce8d484334e1f4e7a923463306980b69bd5bc78f90e999830bd169df08ff9a6f961bc9191eb6ecd1d8c2a512a4cae09444b05600a19d72741d8d8e8f713361ea0d306e990341d252b815003ed8bbd1ae8335f74cc31a31470350bea29eabff794b704796ed0d82e10ff0bb382621ae7ef477cc9b785692582b736d5c286b4598bee6",
		},
	}

	for _, test := range tests {
		want, _ := utils.DecodeHex(test.hash)
		got := Keccak1600(test.input)
		if !bytes.Equal(got, want) {
			t.Errorf("unexpected hash : got '%x' want '%s'", got, test.hash)
		}
	}
}
