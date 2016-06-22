# DOCKER-VERSION 1.7

# gomon (github.com/rmorriso/gomon)

# Scratch container requires a static link of gomon:
# CGO_ENABLED=0 go build -a -installsuffix cgo

FROM alpine:latest
 
MAINTAINER Rick Morrison <rmorriso@fedroot.com>

COPY files/root /

EXPOSE 80 443 8080 

CMD /api -port :8080 -c /etc/monitor.conf && /dashboard -api http://localhost:8080/

