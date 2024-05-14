package repository

import (
	"context"
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	client *mongo.Client
	log    *log.Logger
}

func New(ctx context.Context, log *log.Logger) (*repository, error) {
	uri := os.Getenv("DATABASE_URI")
	if len(uri) == 0 {
		return nil, errors.New("DATABASE_URI environment variable not set")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &repository{
		client: client,
		log:    log,
	}, nil
}
