package app

import (
	"github.com/MosPolyNavigation/web-back/internal/adapters/api"
	"github.com/MosPolyNavigation/web-back/internal/delivery"
	"github.com/gofiber/fiber/v3"
	log "github.com/sirupsen/logrus"
	"os"
)

type app struct {
	log    *log.Logger
	server *fiber.App
}

func New(log *log.Logger) (*app, error) {
	githubApi := api.New(log)
	client := fiber.New()
	handlers := delivery.NewHandlers(githubApi, client, log)
	handlers.RegisterRoute()
	return &app{
		log:    log,
		server: client,
	}, nil
}

func (app *app) Run() error {
	return app.server.Listen(":" + os.Getenv("SERVER_PORT"))
}

func (app *app) Stop(done chan struct{}) error {
	err := app.server.Shutdown()
	if err != nil {
		return err
	}
	close(done)
	return nil
}
