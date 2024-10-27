package set1

import (
	"testing"
)

func TestChallenge3(t *testing.T) {
	challenge_string := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expected := "test"
	msg, err := Challenge3(challenge_string)
	if err != nil || expected != msg {
		t.Fatalf(`Challenge1(%s) = %s, %v; want %s, nil`, challenge_string, msg, err, expected)
	}
}
