default: status-monitor

run:
	go run ./*.go

status-monitor:
	go build -o bin/status-monitor

clean:
	rm -fr bin/*
