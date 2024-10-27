package set1

import (
	"testing"
)

func TestChallenge1(t *testing.T) {
	challenge_string := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	msg, err := Challenge1(challenge_string)
	if err != nil || expected != msg {
		t.Fatalf(`Challenge1(%s) = %s, %v; want %s, nil`, challenge_string, msg, err, expected)
	}
}
