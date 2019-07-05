# Description

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

* Command line options for connecting to `status-go`
* Confirmation dialogs for actions on peers
* Syncing of mailservers
* Info about current node
* Help screen / bindings list
