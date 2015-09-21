# Copyright 2015 The Metalogic Software Authors. All rights reserved.
# Use of this source code is governed by an MIT
# license that can be found in the LICENSE file.

all: clean
	@rm -rf docker/build/files/gomon docker/build/files/root/html docker/build/files/root/inc
	go fmt
	go test -i
	go test
	CGO_ENABLED=0 go build -a -installsuffix cgo -o gomon
	make todo

lint: all
	go vet
	golint .

install: all
	go install

docker: all
	@cp gomon docker/build/files
	@cp -r html inc docker/build/files/root
	sudo docker build -t metalogic/gomon docker/build

todo:
	@grep -n ^[[:space:]]*_[[:space:]]*=[[:space:]][[:alnum:]] *.go || true
	@grep -n TODO *.go || true
	@grep -n FIXME *.go || true
	@grep -n BUG *.go || true

clean:
	rm -f gomon *~

