// Copyright 2015 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/rmorriso/gomon/service"
)

type Services struct {
	Files        []*service.File
	HttpServices []*service.HttpService
	TcpServices  []*service.TcpService
}

// Init unmarshalls Services from JSON configuration in filename
func (services *Services) Init(filename string) {
	if conf, err := ioutil.ReadFile(filename); err != nil {
		log.Fatalf("failed to read %s: %s\n", filename, err)
	} else if err = json.Unmarshal(conf, &services); err != nil {
		log.Fatalf("Config error at %s (while reading %s)\n", err, filename)
	}	
}