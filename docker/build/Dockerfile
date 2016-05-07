# gomon (github.com/rmorriso/gomon)

# Scratch container requires a static link of gomon:
# CGO_ENABLED=0 go build -a -installsuffix cgo

FROM alpine
 
MAINTAINER Rick Morrison <rmorriso@fedroot.com>

COPY files/root /

EXPOSE 80 443 8080 

CMD /gomon -port :80 -c /etc/monitor.conf

