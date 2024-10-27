package set1

import (
	"encoding/hex"
	"log"
)

func Challenge2(input string, xor_string string) (string, error) {
	if len(input) != len(xor_string) {
		log.Fatal("The input string and the XOR string lengths need to match.")
	}

	decodedInput, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}

	decodedXORString, err := hex.DecodeString(xor_string)
	if err != nil {
		log.Fatal(err)
	}

	result := make([]byte, len(decodedInput))
	for i := range decodedInput {
		result[i] = decodedInput[i] ^ decodedXORString[i]
	}

	resultHex := hex.EncodeToString(result)
	return resultHex, nil
}
