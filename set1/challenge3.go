package set1

import (
	"encoding/hex"
	"fmt"
	"log"
)

func Challenge3(input string) (string, error) {
	decodedXORString, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(decodedXORString)
	return "test2", nil
}
