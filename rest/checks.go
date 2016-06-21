// Copyright 2012-2016 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/rmorriso/gomon/check"
)

// Checks aggregates the lists of Gomon checks
type Checks struct {
	FileChecks  []*check.FileCheck
	HTTPChecks  []*check.HTTPCheck
	ShellChecks []*check.ShellCheck
	TcpChecks   []*check.TcpCheck
}

// Init unmarshalls Checks from JSON configuration in filename
func (checks *Checks) Init(filename string) {
	if conf, err := ioutil.ReadFile(filename); err != nil {
		log.Fatalf("failed to read %s: %s\n", filename, err)
	} else if err = json.Unmarshal(conf, &checks); err != nil {
		log.Fatalf("Config error at %s (while reading %s)\n", err, filename)
	}
}
