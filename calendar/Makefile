GOOS?=linux
PUBLISHER_APP?=publisher

build: clean calendar publisher scheduler
calendar: ./apps/calendar/calendar.go
	go build -o ./apps/calendar/bin/calendar ./apps/calendar/calendar.go
publisher: ./apps/publisher/publisher.go
	go build -o ./apps/publisher/bin/publisher ./apps/publisher/publisher.go
scheduler: ./apps/scheduler/scheduler.go
	go build -o ./apps/scheduler/bin/scheduler ./apps/scheduler/scheduler.go
clean:
	rm -rf ./bin
run: docker-compose.yml
	sudo docker-compose up -d

