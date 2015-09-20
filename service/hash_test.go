package service

import (
	"testing"
)

const expected = "Ho1dq1of4MPYvhkk3JV0DPOUmy0="

func TestGetHash(t *testing.T) {
	hash, err := getHash("./testme.txt")
	if err != nil {
		t.Fatalf("unexpected error (%s) getting hash\n", err)
	}
	if hash != expected {
		t.Fatalf("expected %s, got %s\n", expected, hash)
	}
	_, err = getHash("/dev/nosuchfile")
	if err == nil {
		t.Fatalf("expected error opening non-existent file\n")
	}
}
