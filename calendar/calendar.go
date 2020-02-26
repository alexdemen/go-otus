package main

import (
	"errors"
	"github.com/alexdemen/go-otus/calendar/internal/calendarpb"
	"github.com/alexdemen/go-otus/calendar/internal/config"
	"github.com/alexdemen/go-otus/calendar/internal/middleware"
	"github.com/alexdemen/go-otus/calendar/internal/service"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
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
	eventServer := service.NewEventServer()
	calendarpb.RegisterEventServiceServer(grpcServer, eventServer)
	grpcServer.Serve(lis)

	//level, err := chooseLoggerLevel(runConfig.LogLevel)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//err = middleware.ConfigureLogger(runConfig.LogFile, level)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//handler := middleware.SetLogger(http.HandlerFunc(resolvePath))
	//if err := http.ListenAndServe(runConfig.ListenAddress, handler); err != nil {
	//	fmt.Println(err)
	//}
	//
	//middleware.CloseLogger()
}

func resolvePath(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/hello":
		w.Write([]byte("Hello world!"))
		return
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func chooseLoggerLevel(levelName string) (int, error) {
	switch levelName {
	case "debug":
		return middleware.Debug, nil
	case "warn":
		return middleware.Warning, nil
	case "error":
		return middleware.Error, nil
	case "info":
		return middleware.Info, nil
	default:
		return 0, errors.New("failed to select logger level")
	}
}
