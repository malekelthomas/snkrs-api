package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func upsert(ctx context.Context, collection *mongo.Collection, filter bson.D, item interface{}) (bool, error) {
	upsert := true
	result, err := collection.ReplaceOne(ctx, filter, item, &options.ReplaceOptions{
		Upsert: &upsert,
	})

	if err != nil {
		return false, err
	}

	if result.UpsertedCount > 0 {
		return true, nil
	}
	return false, nil
}
