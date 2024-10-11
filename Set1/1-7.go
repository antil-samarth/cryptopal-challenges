/* The Base64-encoded content in this file has been encrypted via AES-128 in ECB mode under the key

"YELLOW SUBMARINE".
(case-sensitive, without the quotes; exactly 16 characters; I like "YELLOW SUBMARINE" because it's exactly 16 bytes long,
and now you do too).

Decrypt it. You know the key, after all.

Easiest way: use OpenSSL::Cipher and give it AES-128-ECB as the cipher.

Do this with code.
You can obviously decrypt this using the OpenSSL command-line tool, but we're having you get ECB working in code for a
reason. You'll need it a lot later on, and not just for attacking ECB.
*/

package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"os"
)

func challenge1_7() {

	raw, err := os.ReadFile("1-7.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	data, err := base64.StdEncoding.DecodeString(string(raw))
	if err != nil {
		fmt.Println("Error decoding base64:", err)
		return
	}

	KEY := "YELLOW SUBMARINE"

	cipher, err := aes.NewCipher([]byte(KEY))
	if err != nil {
		fmt.Println("Error creating cipher:", err)
		return
	}

	decrypted := make([]byte, len(data))
	for i := 0; i < len(data); i += 16 {
		cipher.Decrypt(decrypted[i:], data[i:])
	}

	fmt.Println("String decrypted successfully!")
	/* fmt.Println(string(decrypted)) */

}
