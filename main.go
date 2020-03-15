package main

import (
	//"flag"
	"flag"
	"fmt"
	"net/http"

	"advance-go/internal/api"
	"advance-go/internal/config"
)

func main() {
	port := flag.String("port", "8080", "default port: 8080")
	stage := flag.String("stage", "dev", "default as dev")
	cfgPath := flag.String("config", "configs", "config path")
	flag.Parse()

	conf := &config.Config{}
	if err := conf.Init(*stage, *cfgPath); err != nil {
		return
	}

	route := api.Init(conf)
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", *port),
		Handler: route,
	}
	server.ListenAndServe()
	fmt.Println("what!!")
}
