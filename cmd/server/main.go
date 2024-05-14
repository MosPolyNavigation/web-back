package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/MosPolyNavigation/web-back/internal/app"

	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()

	a, err := app.New(logger)
	if err != nil {
		logger.Fatal(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	complete := make(chan struct{})

	go func() {
		<-sig
		if err := a.Stop(complete); err != nil {
			logger.Fatal(err)
		}
	}()

	if err := a.Run(); err != nil {
		logger.Fatal(err)
	}
}
