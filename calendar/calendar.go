package main

import (
	"fmt"
	"github.com/alexdemen/go-otus/calendar/internal/calendarpb"
	"github.com/alexdemen/go-otus/calendar/internal/config"
	"github.com/alexdemen/go-otus/calendar/internal/middleware/logger"
	"github.com/alexdemen/go-otus/calendar/internal/service"
	"github.com/alexdemen/go-otus/calendar/internal/store"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

var configPath = flag.String("config", "", "path to configuration file")

func init() {
	flag.Parse()
}

func main() {
	runConfig, err := config.GetConfiguration(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", runConfig.ListenAddress)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	storage := store.NewMemoryStore()
	var eventServer calendarpb.EventServiceServer = service.NewEventServer(storage)
	eventServer = logger.NewMiddlewareLogger(eventServer)
	calendarpb.RegisterEventServiceServer(grpcServer, eventServer)

	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt)
	go func() {
		<-done
		fmt.Println("interrupt")
		grpcServer.GracefulStop()
	}()

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

	fmt.Println("end")
}
