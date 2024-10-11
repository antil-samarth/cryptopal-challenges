package main

import "fmt"

func main() {
	fmt.Println("Set 1")
	fmt.Println()

	// 1-1
	fmt.Println("Challenge 1-1 : Convert hex to base64")
	challenge1_1()
	fmt.Println()

	// 1-2
	fmt.Println("Challenge 1-2 : Fixed XOR")
	challenge1_2()
	fmt.Println()

	// 1-3
	fmt.Println("Challenge 1-3 : Single-byte XOR cipher")
	challenge1_3()
	fmt.Println()

	// 1-4
	fmt.Println("Challenge 1-4 : Detect single-character XOR")
	challenge1_4()
	fmt.Println()

	// 1-5
	fmt.Println("Challenge 1-5 : Implement repeating-key XOR")
	challenge1_5()
	fmt.Println()

	// 1-6
	fmt.Println("Challenge 1-6 : Break repeating-key XOR")
	challenge1_6()
	fmt.Println()

	// 1-7
	fmt.Println("Challenge 1-7 : AES in ECB mode")
	challenge1_7()
	fmt.Println()

	// 1-8
	fmt.Println("Challenge 1-8 : Detect AES in ECB mode")
	challenge1_8()
}
