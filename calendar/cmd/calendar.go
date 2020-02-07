package main

import (
	"fmt"
	"github.com/alexdemen/go-otus/calendar/internal/config"
	"github.com/alexdemen/go-otus/calendar/internal/middleware"
	flag "github.com/spf13/pflag"
	"log"
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

	logger := middleware.Logger{}

	handler := logger.SetLogger(http.HandlerFunc(resolvePath))
	if err := http.ListenAndServe(runConfig.ListenAddress, handler); err != nil {
		fmt.Println(err)
	}
}

func resolvePath(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/hello":
		w.Write([]byte("Hello world!"))
		return
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
