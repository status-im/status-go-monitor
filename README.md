# Status Go Monitor [![Go Report Card](https://goreportcard.com/badge/github.com/status-im/status-go-monitor)](https://goreportcard.com/report/github.com/status-im/status-go-monitor)

This is a console client for manging peers connected to an instance of [status-go](https://github.com/status-im/status-go).

![](/img/status-go-monitor.png)

# Details

It is written in Go using [gocui](https://github.com/jroimartin/gocui) library.

# Building

The simplest way is to just run:
```
make build
```
And use it from the `bin` folder:
```
./bin/status-monitor
```

# TODO

* Syncing of mailservers
* Info about current node
