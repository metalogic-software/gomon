# gomon container

## Prerequistes

This container relies on a static build of github.com/rmorriso/gomon:

```CGO_ENABLED=0 go build -a -installsuffix cgo```

Build the executable as above and copy to build/files/gomon.

Edit and copy etc/monitor.conf configuration to /data/gomon/etc/monitor.conf.
This will be mounted at /etc in the container.

## Building

From build directory:

`# docker build -t metalogic/gomon .`

## Running from the Command Line

```# docker run --name gomon -p 8080:8080 -v /data/gomon/etc:/etc/gomon  metalogic/gomon /gomon```

## Deploying

Copy the systemd unit file systemd/notbot.service to /etc/systemd/system.
 
```# systemctl start gomon```

To enable at boot time:

```# systemctl enable gomon```

