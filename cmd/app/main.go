package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mezmerizxd/go-app/pkg/cache/redis"
	"github.com/mezmerizxd/go-app/pkg/config"
	"github.com/mezmerizxd/go-app/pkg/data"
	"github.com/mezmerizxd/go-app/pkg/firebase"
	"github.com/mezmerizxd/go-app/pkg/server"
	"github.com/mezmerizxd/go-app/pkg/version"
)

func main() {
	// Initialize the server
	cfg := config.New(":3001")
	cacheImpl := redis.New(cfg.Redis)
	firebaseImpl := firebase.New(&firebase.Config{
		Cache: cacheImpl,
		DatabaseUrl: "https://locus-public-default-rtdb.firebaseio.com/",
		CredentialPath: "/home/mezmerizxd/Documents/locus_firebase.json",
	})
	dataImpl := data.New(&data.Config{
		Cache: cacheImpl,
		Firebase: firebaseImpl,
	})

	srv := server.New(cfg.ServerAddr, &version.Config{
		Cache: cacheImpl,
		Data:  dataImpl,
		Firebase: firebaseImpl,
	})

	// Create channel for quitting server
	quitChannel := make(chan bool, 1)

	// Start the server
	go func() {
		log.Println("Starting at http://localhost" + cfg.ServerAddr)
		if err := srv.Start(); err != nil {
			select {
			case <-quitChannel:
				return
			default:
				log.Println("Failed to start server: " + err.Error())
			}
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	quitChannel <- true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := srv.Stop(ctx); err != nil {
		log.Println("Failed to stop server: " + err.Error())
	}

	log.Println("exiting")
}
