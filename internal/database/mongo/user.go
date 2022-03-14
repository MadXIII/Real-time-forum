package mongo

import (
	"context"
	"fmt"
	"forum/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Store) InsertUser(ctx context.Context, user models.User) error {
	res, err := s.Collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("InsertUser, InsertOne: %w", err)
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if ok {
		user.ID = id.Hex()
		return nil
	}
	return nil
}

func (s *Store) GetUserByLogin(ctx context.Context, username string) (models.User, error) {
	var user models.User
	filter := bson.M{"username": username}

	res := s.Collection.FindOne(ctx, filter)

	if res.Err() != nil {
		return user, fmt.Errorf("GetUserByLogin, FindOne: %w", res.Err())
	}

	if err := res.Decode(&user); err != nil {
		return user, fmt.Errorf("GetUserByLogin, Decode: %w", err)
	}

	return user, nil
}

func (s *Store) GetUsernameByID(ctx context.Context, id string) (string, error) {
	var username string
	filter := bson.M{"_id": id, "username": 1}
	projection := bson.D{
		{"username", 1},
	}

	res := s.Collection.FindOne(ctx, filter, options.FindOne().SetProjection(projection))
	// s.Collection.FindOne(ctx, bson.M{"_id": id}, &options.FindOneOptions{})
	if res.Err() != nil {
		return username, fmt.Errorf("GetUsernameByID, FindOne: %w", res.Err())
	}

	if err := res.Decode(&username); err != nil {
		return username, fmt.Errorf("GetUsernameByID, Decode: %w", err)
	}

	return username, nil
}
