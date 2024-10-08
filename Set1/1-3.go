/* Single-byte XOR cipher
The hex encoded string:

1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
... has been XOR'd against a single character. Find the key, decrypt the message.

You can do this by hand. But don't: write code to do it for you. */

package main

import (
	"encoding/hex"
	"fmt"
	"strings"
	"unicode"
)

func xorStringAndByte(input []byte, key byte) string {
	output := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ key
	}
	return string(output)
}

func isPrintable(s string) bool {
	for _, char := range s {
		if !unicode.IsPrint(char) {
			return false
		}
	}
	return true
}

func scoreEnglish(s string) int {
	commonLetters := "ETAOIN SHRDLU"
	score := 0
	for _, char := range strings.ToUpper(s) {
		if strings.ContainsRune(commonLetters, char) {
			score++
		}
	}
	return score
}

func challenge1_3() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	// convert hex string to byte array
	bytes, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	bestScore := 0
	var bestResult string
	var bestKey byte

	// loop through all possible single-byte XOR keys
	for i := 0; i < 256; i++ {
		result := xorStringAndByte(bytes, byte(i))
		if isPrintable(result) {
			score := scoreEnglish(result)
			if score > bestScore {
				bestScore = score
				bestResult = result
				bestKey = byte(i)
			}
		}
	}

	fmt.Printf("Best Key: %d, Best Result: %s\n", bestKey, bestResult)
}
