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

// getHash computes the hash of the file contents at path
func getHash(path string) (hash string, err error) {
	// make a new hash calculator
	hasher := sha1.New()

	// if we can open the file...
	var file *os.File
	if file, err = os.Open(path); err == nil {
		// if we can copy its contents to the hash
		if _, err = io.Copy(hasher, file); err == nil {
			// return the hash
			hash, err = encodeBase64(hasher.Sum(nil)), nil
		}
	}
	return hash, err
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
