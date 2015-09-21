# Copyright 2015 The Metalogic Software Authors. All rights reserved.
# Use of this source code is governed by an MIT
# license that can be found in the LICENSE file.

all: clean
	go fmt
	go test -i
	go test
	CGO_ENABLED=0 go build -a -installsuffix cgo -o gomon
	go vet
	golint .
	make todo

install: all
	go install

docker: all
	@cp gomon docker/build/files
	docker build -t metalogic/gomon docker/build

todo:
	@grep -n ^[[:space:]]*_[[:space:]]*=[[:space:]][[:alnum:]] *.go || true
	@grep -n TODO *.go || true
	@grep -n FIXME *.go || true
	@grep -n BUG *.go || true

clean:
	rm -f gomon *~

