# gomon
*a simple example using goroutines to poll things of interest*

This is one of the first programs I wrote in Go. At the time - 2011 - that
was pre-Go 1.0. It has compiled under each subsequent release, now at
1.5.1 without modification. At least in my case Go has more than lived up to
its "compatibility promise."

I chose this problem to work on since it gave me an opportunity to play
with goroutines, channels, interfaces and even HTML templates.
Perhaps not all was strictly necessary, but it was a good learning exercise.

## Building

```
% go get github.com/rmorriso/gomon
% make gomon
```

Note that the build is configured to produce a fully static binary for inclusion
in a bare Docker container (scratch or busybox).

## Running

```% ./gomon```

will use the included default monitor configuration defined in monitor.conf.

You can see the effect of changing the state of a monitored item
in the default config by editing
the contents of the monitored file "changeme" or by stopping/starting
sshd on localhost.

## Testing

There are just a few tests:

```% make test```

## Build the Docker Container (Linux Only)

```% make docker```

## Run the Docker Container

The container uses the default config file monitor.conf.

To run the docker container, and supply the monitored file changeme copy changeme to /tmp and execute the following:

```% sudo docker run --name gomon -d -v /tmp/changeme:/changeme -p 8080:8080 metalogic/gomon /gomon```

Make a change to the file in the container:

```% sudo docker exec -it gomon /bin/sh -c "echo you are changed  >> /changeme"```

## TODO

* expose REST API for accessing monitor state, updating pollable properties, etc
* replace front-end with Angular
* delete old monitor data
* more tests

... but hey, it's just an exercise. Maybe one day :-)

