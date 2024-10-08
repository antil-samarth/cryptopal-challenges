/* Convert hex to base64
The string:

49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
Should produce:

SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t

Cryptopals Rule
Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing. */

package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func challenge1_1() {
	hexInput := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	output := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	rawBytes, err := hex.DecodeString(hexInput)
	if err != nil {
		fmt.Printf("Error decoding hex string: %v\n", err)
		return
	}

	base64Output := base64.StdEncoding.EncodeToString(rawBytes)

	/* fmt.Printf("Input (hex):  %s\n", hexInput)
	fmt.Printf("Output (base64): %s\n", base64Output) */

	if base64Output != output {
		fmt.Printf("Output does not match expected result\n")
		return
	}
	fmt.Printf("Output matches expected result\n")

}
