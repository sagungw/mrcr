package datastore

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	db *mongo.Database
	c  *mongo.Client
)

func Open(databaseUrl string) error {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseUrl))
	if err != nil {
		return errors.Wrap(err, "datastore: error establishing datastore connection")
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return errors.Wrap(err, "datastore: error establishing datastore connection")
	}

	c = client
	db = client.Database("mercari")
	return nil
}

func Close() error {
	if err := c.Disconnect(context.Background()); err != nil {
		return errors.Wrap(err, "datastore: error closing datastore connection")
	}

	return nil
}
