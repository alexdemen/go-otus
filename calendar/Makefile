GOOS?=linux

build: clean calendar publisher

calendar: calendar.go
	go build -o ./bin/calendar ./calendar.go
publisher: ./apps/publisher/publisher.go
	go build -o ./bin/publisher ./apps/publisher/publisher.go
clean:
	rm -rf ./bin