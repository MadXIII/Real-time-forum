package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func (s *Store) InitMongoStore(uri string) error {
	var err error
	s.Client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	s.Collection = s.Client.Database("forum").Collection("user")

	return nil
}

func (s *Store) CheckCategoryID(categoryID int) error {
	return nil
}
