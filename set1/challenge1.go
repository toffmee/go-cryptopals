package set1

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func Challenge1(input string) (string, error) {
	data, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}

	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded, nil
}
