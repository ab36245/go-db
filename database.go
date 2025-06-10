package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Database struct {
	mongo *mongo.Database
}

func Open(config Config, name string) (*Database, error) {
	host := config.Host
	if host == "" {
		host = "localhost"
	}
	port := config.Port
	if port == 0 {
		port = 27017
	}
	username := config.Username
	if username == "" {
		return nil, fmt.Errorf("no username")
	}
	password := config.Password
	if password == "" {
		return nil, fmt.Errorf("no password")
	}

	url := fmt.Sprintf("mongodb://%s:%s@%s:%d", username, password, host, port)
	opts := options.Client().ApplyURI(url)
	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, err
	}
	mongo := client.Database(name)
	database := &Database{
		mongo: mongo,
	}
	return database, nil
}

func (d *Database) Drop() error {
	ctx := context.TODO()
	return d.mongo.Drop(ctx)
}
