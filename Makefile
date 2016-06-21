# Copyright 2012-2016 The Metalogic Software Authors. All rights reserved.
# Use of this source code is governed by an MIT
# license that can be found in the LICENSE file.

all: clean test gomon
	go fmt
	make todo

gomon: clean
	CGO_ENABLED=0 go build -a -installsuffix cgo -o gomon

test: 
	go test -i github.com/rmorriso/gomon
	go test -v github.com/rmorriso/gomon
	go test -i github.com/rmorriso/gomon/check
	go test -v github.com/rmorriso/gomon/check

lint: 
	@go vet
	@golint .

install: all
	go install

docker: all
	@cp gomon docker/build/files/root/gomon
	@cp monitor.conf docker/build/files/root/etc/monitor.conf
	@cp -r html inc docker/build/files/root
	sudo docker build -t metalogic/gomon docker/build

run: docker
	sudo docker stop gomon
	sudo docker rm gomon
	sudo docker run -d --name gomon -p 8080:8080 metalogic/gomon /gomon

todo:
	@find . -name \*.go -exec grep -Hn ^[[:space:]]*_[[:space:]]*=[[:space:]][[:alnum:]] \{\} \;
	@find . -name \*.go -exec grep -Hn TODO \{\} \;
	@find . -name \*.go -exec grep -Hn FIXME \{\} \;
	@find . -name \*.go -exec grep -Hn BUG \{\} \;

clean:
	go clean
	rm -f gomon *~
	rm -rf docker/build/files/root/gomon docker/build/files/root/etc/monitor.conf docker/build/files/root/html docker/build/files/root/inc
