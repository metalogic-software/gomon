# DOCKER-VERSION 1.7

# gomon (github.com/rmorriso/gomon)

# Scratch container requires a static link (Go 1.4.1) of gomon:
# CGO_ENABLED=0 go build -a -installsuffix cgo

FROM busybox
 
MAINTAINER Rick Morrison <rmorriso@fedroot.com>

COPY files/root /
ADD files/gomon /gomon
ADD files/monitor.conf /monitor.conf

EXPOSE 80 443 8080 

