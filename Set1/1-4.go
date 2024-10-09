/* Detect single-character XOR
One of the 60-character strings in the file 1-4.txt has been encrypted by single-character XOR.

Find it.*/

package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"unicode"
)

func challenge1_4() {
	file, err := os.Open("1-4.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bestScore := 0
	var bestLine, bestDecrypted string

	for scanner.Scan() {
		line := scanner.Text()
		decoded, _ := hex.DecodeString(line)

		for key := 0; key < 256; key++ {
			decrypted := xorDecrypt(decoded, byte(key))
			score := scoreText(decrypted)

			if score > bestScore {
				bestScore = score
				bestLine = line
				bestDecrypted = string(decrypted)
			}
		}
	}

	fmt.Println("Best matching line:", bestLine)
	fmt.Print("Decrypted message:", bestDecrypted)
}

func xorDecrypt(input []byte, key byte) []byte {
	output := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ key
	}
	return output
}

func scoreText(text []byte) int {
	score := 0
	for _, c := range text {
		if unicode.IsLetter(rune(c)) || unicode.IsSpace(rune(c)) {
			score++
		}
	}
	return score
}
