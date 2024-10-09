/*
Break repeating-key XOR

It is officially on, now.
This challenge isn't conceptually hard, but it involves actual error-prone coding. The other challenges in this set are
there to bring you up to speed. This one is there to qualify you. If you can do this one, you're probably just fine up to
Set 6.

There's a file here name '1-6.txt'. It's been base64'd after being encrypted with repeating-key XOR.

Decrypt it.

Here's how:

1. Let KEYSIZE be the guessed length of the key; try values from 2 to (say) 40.

2. Write a function to compute the edit distance/Hamming distance between two strings. The Hamming distance is just the
number of differing bits.
	The distance between: 'this is a test' and 'wokka wokka!!!' is 37.
	Make sure your code agrees before you proceed.

3. For each KEYSIZE, take the first KEYSIZE worth of bytes, and the second KEYSIZE worth of bytes, and find the edit
distance between them. Normalize this result by dividing by KEYSIZE.

4. The KEYSIZE with the smallest normalized edit distance is probably the key. You could proceed perhaps with the smallest
2-3 KEYSIZE values. Or take 4 KEYSIZE blocks instead of 2 and average the distances.

5. Now that you probably know the KEYSIZE: break the ciphertext into blocks of KEYSIZE length.

6. Now transpose the blocks: make a block that is the first byte of every block, and a block that is the second byte of
every block, and so on.

7. Solve each block as if it was single-character XOR. You already have code to do this.

8. For each block, the single-byte XOR key that produces the best looking histogram is the repeating-key XOR key byte for
that block. Put them together and you have the key.

This code is going to turn out to be surprisingly useful later on. Breaking repeating-key XOR ("Vigenere") statistically
is obviously an academic exercise, a "Crypto 101" thing. But more people "know how" to break it than can actually break it,
and a similar technique breaks something much more important.
*/

/*
func main () {
	// Define input strings
	input1 := "this is a test"
	input2 := "wokka wokka!!!"

	// Calculate and print the Hamming distance
	fmt.Println(hammingDistance(input1, input2))

	// Verify if the calculated distance is correct (should be 37)
	fmt.Println(hammingDistance(input1, input2) == 37)
}
*/

package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"sort"
)

func challenge1_6() {
	// Read the base64-encoded file
	raw, err := os.ReadFile("1-6.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Decode the base64-encoded data into raw bytes
	data, err := base64.StdEncoding.DecodeString(string(raw))
	if err != nil {
		fmt.Println("Error decoding base64:", err)
		return
	}

	type KeySizeResult struct {
		keySize            int
		normalizedDistance float64
	}

	var results []KeySizeResult

	// Iterate over different KEYSIZE values from 2 to 40
	for KEYSIZE := 2; KEYSIZE <= 40; KEYSIZE++ {
		// Make sure we have enough data for 4 chunks of size KEYSIZE
		if len(data) < 4*KEYSIZE {
			continue
		}

		// Split the data into chunks
		chunk1 := data[:KEYSIZE]
		chunk2 := data[KEYSIZE : 2*KEYSIZE]
		chunk3 := data[2*KEYSIZE : 3*KEYSIZE]
		chunk4 := data[3*KEYSIZE : 4*KEYSIZE]

		// Calculate the Hamming distances between consecutive chunks
		distance1 := hammingDistance(chunk1, chunk2)
		distance2 := hammingDistance(chunk3, chunk4)
		distance3 := hammingDistance(chunk1, chunk3)
		distance4 := hammingDistance(chunk2, chunk4)

		// Average the distances and normalize by dividing by KEYSIZE
		normalizedDistance := (float64(distance1+distance2+distance3+distance4) / float64(4*KEYSIZE))

		// Add result to the slice
		results = append(results, KeySizeResult{KEYSIZE, normalizedDistance})

		/* fmt.Printf("KEYSIZE: %d, Normalized Distance: %.2f\n", KEYSIZE, normalizedDistance)*/
	}

	// Sort results by normalized distance
	sort.Slice(results, func(i, j int) bool {
		return results[i].normalizedDistance < results[j].normalizedDistance
	})

	// Print top 3 best KEYSIZEs
	/* fmt.Println("Top 3 best KEYSIZEs:")
	for i := 0; i < 3 && i < len(results); i++ {
		fmt.Printf("%d. KEYSIZE: %d, Normalized Distance: %.2f\n", i+1, results[i].keySize, results[i].normalizedDistance)
	} */

	// Transpose the blocks based on the best KEYSIZE
	// Solve each block as if it was single-character XOR
	// Find the best scoring key for this block
	keys := make([][]byte, 3)
	for i := 0; i < 3; i++ {
		bestKEYSIZE := results[i].keySize
		keys[i] = computeKey(bestKEYSIZE, data)
	}

	// score the key
	bestScore := 0
	var bestKey []byte
	for i := 0; i < 3; i++ {
		score := scoreText(keys[i])
		if score > bestScore {
			bestScore = score
			bestKey = keys[i]
		}
	}

	fmt.Println("Best key:", string(bestKey))

	// Decrypt the full data using the key
	decryptedData := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		decryptedData[i] = data[i] ^ bestKey[i%len(bestKey)]
	}

	// Print first 10 characters of decrypted data
	fmt.Println(string(decryptedData[:34]))
}

func computeKey(bestKEYSIZE int, data []byte) []byte {
	transposedBlocks := make([][]byte, bestKEYSIZE)
	for i := 0; i < bestKEYSIZE; i++ {
		for j := i; j < len(data); j += bestKEYSIZE {
			transposedBlocks[i] = append(transposedBlocks[i], data[j])
		}
	}

	key := make([]byte, bestKEYSIZE)

	for i := 0; i < bestKEYSIZE; i++ {
		bestScore := 0
		for j := 0; j < 256; j++ {
			decrypted := xorDecrypt(transposedBlocks[i], byte(j))
			score := scoreText(decrypted)

			if score > bestScore {
				bestScore = score
				key[i] = byte(j)
			}
		}
	}
	return key
}

// hammingDistance calculates the Hamming distance between two byte slices
func hammingDistance(input1, input2 []byte) int {
	distance := 0

	for i := 0; i < len(input1); i++ {
		// XOR the bytes and count the number of differing bits
		distance += countBits(input1[i] ^ input2[i])
	}

	return distance
}

// countBits returns the number of 1s in the binary representation of a byte
func countBits(b byte) int {
	count := 0
	for b != 0 {
		count += int(b & 1)
		b >>= 1
	}
	return count
}
