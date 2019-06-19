
default: status-monitor

run:
	go run ./*.go

status-monitor:
	go build -o build/bin/status-monitor

clean:
	rm -fr build/bin/*
