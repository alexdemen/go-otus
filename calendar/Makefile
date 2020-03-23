GOOS?=linux

build: clean calendar publisher scheduler
calendar: ./apps/calendar/calendar.go
	go build -o ./bin/calendar ./apps/calendar/calendar.go
publisher: ./apps/publisher/publisher.go
	go build -o ./bin/publisher ./apps/publisher/publisher.go
scheduler: ./apps/scheduler/scheduler.go
	go build -o ./bin/scheduler ./apps/scheduler/scheduler.go
clean:
	rm -rf ./bin
run: docker-compose.yml
	sudo docker-compose up -d
