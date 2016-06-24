# DOCKER-VERSION 1.7

# gomon (github.com/rmorriso/gomon)

# Scratch container requires a static link of gomon:
# CGO_ENABLED=0 go build -a -installsuffix cgo

FROM alpine:latest
 
MAINTAINER Rick Morrison <rmorriso@fedroot.com>

RUN apk update
RUN apk --no-cache add ca-certificates curl

EXPOSE 80 443 8080 

CMD /api -port :8081 -c /etc/monitor.conf && /dashboard -port 8080 -api http://localhost:8081/

