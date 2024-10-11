/* Detect AES in ECB mode
In this file are a bunch of hex-encoded ciphertexts.

One of them has been encrypted with ECB.

Detect it.

Remember that the problem with ECB is that it is stateless and deterministic;
the same 16 byte plaintext block will always produce the same 16 byte ciphertext.
*/

package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
)

func challenge1_8() {
	// Read the file
	file, err := os.Open("1-8.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Convert the line to a byte slice
		ciphertext, err := hex.DecodeString(line)
		if err != nil {
			fmt.Println("Error decoding hex:", err)
			continue
		}

		// Detect ECB
		if isECB(ciphertext) {
			/* fmt.Println("ECB detected in:", line) */
			fmt.Println("ECB found!")
			return
		}
	}
}

func isECB(ciphertext []byte) bool {
	// Check if the ciphertext has repeated
	// 16-byte blocks
	blockSize := 16
	for i := 0; i < len(ciphertext); i += blockSize {
		for j := i + blockSize; j < len(ciphertext); j += blockSize {
			if bytes.Equal(ciphertext[i:i+blockSize], ciphertext[j:j+blockSize]) {
				return true
			}
		}
	}
	return false
}
