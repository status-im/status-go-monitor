default: status-monitor

run:
	go run ./*.go

build:
	go build -o bin/status-monitor

clean:
	rm -fr bin/*
