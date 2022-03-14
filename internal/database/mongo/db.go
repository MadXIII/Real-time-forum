package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	Collection *mongo.Collection
}

func (s *Store) InitMongoStore(ctx context.Context, uri string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return client, err
	}

	if err = client.Connect(ctx); err != nil {
		return client, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return client, err
	}

	s.Collection = client.Database("forum").Collection("user")

	return client, nil
}
