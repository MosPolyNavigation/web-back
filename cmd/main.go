package main

import (
	"back/internal/app"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := logrus.New()

	a := app.New()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	complete := make(chan struct{})

	go func() {
		<-sig
		if err := a.Stop(complete); err != nil {
			log.Fatal(err)
		}
	}()

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
