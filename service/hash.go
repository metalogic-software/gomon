// Copyright 2015 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package service

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"io"
	"os"
)

const EMPTY_HASH = "ZERO"

// Given a path string and a Hash interface, calculates the hash of the file's
// content.
func getHash(path string) (string, error) {
	// make a new hash calculator
	hash := sha1.New()

	// if we can open the file...
	file, err := os.Open(path)
	if err == nil {
		// if we can copy its contents to the hash
		_, err = io.Copy(hash, file)
		if err == nil { // return the hash
			return encodeBase64(hash.Sum(nil)), nil
		}
	}
	// error return
	return "", err
}

// this is a fairly nasty way of getting a string out of a byte array in Base64
func encodeBase64(source []byte) string {
	// make a byte slice just big enough for the result of the encode operation
	dest := make([]byte, base64.StdEncoding.EncodedLen(len(source)))
	// encode it
	base64.StdEncoding.Encode(dest, source)
	// convert this byte buffer to a string
	return bytes.NewBuffer(dest).String()
}
