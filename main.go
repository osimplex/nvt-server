package main

import (
	"log"
	"nvt-server/src/common"
	"nvt-server/src/server"
	"os"
	"os/signal"
	"syscall"
)

const (
	// configFile = "etc/oracle.toml"
	configFile = "etc/example.toml"
)

func main() {
	var server server.Server

	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, syscall.SIGINT, syscall.SIGTERM)
	sigHup := make(chan os.Signal, 1)
	signal.Notify(sigHup, syscall.SIGHUP)

serverLoop:
	for {
		server = serverSetup()

		select {
		case <-sigHup:
			log.Printf("HUP signal received")
			server.End()
		case <-sigTerm:
			log.Printf("TERM or INT signal received")
			server.End()
			break serverLoop
		}
	}
}

func serverSetup() server.Server {
	config, err := common.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("config loading error: %s", err)
	}

	server, err := server.Factory(config)
	if err != nil {
		log.Fatalf("server provisioning error: %s", err)
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server setup error: %s", err)
	}

	return server
}
