package app

import (
	"context"
	"fmt"
	"github.com/MosPolyNavigation/web-back/internal/adapters/repository"
	log "github.com/sirupsen/logrus"
)

type app struct {
	log log.Logger
}

func New(log *log.Logger) (*app, error) {
	ctx := context.Background()
	repo, err := repository.New(ctx, log)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize repository: %w", err)
	}
	_ = repo
	return &app{}, nil
}

func (app *app) Run() error {
	return nil
}

func (app *app) Stop(done chan struct{}) error {
	return nil
}
