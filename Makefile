# Copyright 2012-2016 The Metalogic Software Authors. All rights reserved.
# Use of this source code is governed by an MIT
# license that can be found in the LICENSE file.

gomon: rest/api web/dashboard
	gofmt -w .

all:
	make gomon
	make test
	make lint
	make todo

rest/api: rest/*.go check/*.go monitor/*.go
	make test
	cd check && CGO_ENABLED=0 go build -a -installsuffix cgo
	cd monitor && CGO_ENABLED=0 go build -a -installsuffix cgo
	cd rest && CGO_ENABLED=0 go build -a -installsuffix cgo -o api

web/dashboard: web/*.go
	go test -v github.com/rmorriso/gomon/web
	cd web && CGO_ENABLED=0 go build -a -installsuffix cgo -o dashboard

test: 
	go test -i github.com/rmorriso/gomon/check
	go test ${short} -v github.com/rmorriso/gomon/check
	go test -i github.com/rmorriso/gomon/monitor
	go test ${short} -v github.com/rmorriso/gomon/monitor
	go test -i github.com/rmorriso/gomon/rest
	go test ${short} -v github.com/rmorriso/gomon/rest
	go test -i github.com/rmorriso/gomon/web
	go test ${short} -v github.com/rmorriso/gomon/web

shorttest:
	make test short="-short"
lint: 
	-go vet github.com/rmorriso/gomon/check github.com/rmorriso/gomon/monitor github.com/rmorriso/gomon/rest github.com/rmorriso/gomon/web
	golint check monitor rest web

install: all
	go install

vendor:
	govendor update +vendor

docker: all
	@cp rest/api web/dashboard docker/assets/root/
	@cp rest/monitor.conf docker/assets/root/etc/monitor.conf
	@cp -r web/html web/inc docker/assets/root
	sudo docker build -t metalogic/gomon docker

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
	rm -f rest/api
	rm -f web/dashboard
	rm -rf docker/build/files/root/{api,dashboard} docker/build/files/root/etc/monitor.conf docker/build/files/root/html docker/build/files/root/inc

