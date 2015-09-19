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

```% go get github.com/rmorriso/gomon```

## Running

```% ./gomon```

will use the included default monitor configuration defined in monitor.conf.

You can see the effect of changing the state of a monitored item
in the default config by editing
the contents of the monitored file "changeme" or by stopping/starting
sshd on localhost.

## TODO

* expose REST API for accessing monitor state, updating pollable properties, etc
* replace front-end with Angular
* delete old monitor data

... but hey, it's just an exercise. Maybe one day :-)

