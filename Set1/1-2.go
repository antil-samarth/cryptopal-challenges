/* Fixed XOR
Write a function that takes two equal-length buffers and produces their XOR combination.

If your function works properly, then when you feed it the string:

1c0111001f010100061a024b53535009181c
... after hex decoding, and when XOR'd against:

686974207468652062756c6c277320657965
... should produce:

746865206b696420646f6e277420706c6179 */

package main

import (
	"encoding/hex"
	"fmt"
)

func xorHexStrings(a, b string) (string, error) {
	aRaw, err := hex.DecodeString(a)
	if err != nil {
		return "", fmt.Errorf("error decoding first hex string: %v", err)
	}

	bRaw, err := hex.DecodeString(b)
	if err != nil {
		return "", fmt.Errorf("error decoding second hex string: %v", err)
	}

	if len(aRaw) != len(bRaw) {
		return "", fmt.Errorf("hex strings have different lengths")
	}

	xorResult := make([]byte, len(aRaw))
	for i := range aRaw {
		xorResult[i] = aRaw[i] ^ bRaw[i]
	}

	return hex.EncodeToString(xorResult), nil
}

func challenge1_2() {
	a := "1c0111001f010100061a024b53535009181c"
	b := "686974207468652062756c6c277320657965"
	expectedResult := "746865206b696420646f6e277420706c6179"

	result, err := xorHexStrings(a, b)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if result != expectedResult {
		fmt.Println("Output does not match expected result")
	} else {
		fmt.Println("Output matches expected result")
	}
}
