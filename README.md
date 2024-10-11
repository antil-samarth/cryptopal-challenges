# Cryptopals Challenges

This repository contains solutions to the [Cryptopals Crypto Challenges](https://www.cryptopals.com/) implemented in Go.

## Set 1: Basics

### Challenges Completed:

1. Convert hex to base64 (1-1.go)
2. Fixed XOR (1-2.go)
3. Single-byte XOR cipher (1-3.go)
4. Detect single-character XOR (1-4.go)
5. Implement repeating-key XOR (1-5.go)
6. Break repeating-key XOR (1-6.go)
7. AES in ECB mode (1-7.go)
8. Detect AES in ECB mode (1-8.go)

### Running the Challenges

To run all the challenges in Set 1, execute the following command:

```bash
cd Set1
go run Set1/set1.go
```


This will run the main function in set1.go, which calls each challenge function sequentially.

### Individual Challenges

Each challenge is implemented in its own file:

- 1-1.go: Converts a hexadecimal string to base64
- 1-2.go: Performs a fixed XOR operation on two equal-length buffers
- 1-3.go: Decrypts a single-byte XOR cipher
- 1-4.go: Detects single-character XOR in a file
- 1-5.go: Implements repeating-key XOR encryption
- 1-6.go: Breaks repeating-key XOR encryption
- 1-7.go: Implements AES in ECB mode
- 1-8.go: Detects AES in ECB mode
