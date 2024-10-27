package set1

import (
	"testing"
)

func TestChallenge2(t *testing.T) {
	challenge_string := "1c0111001f010100061a024b53535009181c"
	xor_against := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"
	msg, err := Challenge2(challenge_string, xor_against)
	if err != nil || expected != msg {
		t.Fatalf(`Challenge1(%s) = %s, %v; want %s, nil`, challenge_string, msg, err, expected)
	}
}
