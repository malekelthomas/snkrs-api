package mongo

import (
	"context"
	"fmt"
	"os"

	"errors"

	"snkrs/pkg/create"
	"snkrs/pkg/get"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	db *mongo.Database
}

func NewMongoStore(conn string) (*Store, error) {
	ctx := context.TODO()
	var clientOps options.ClientOptions
	clientOps.ApplyURI(conn)
	client, err := mongo.Connect(ctx, &clientOps)
	if err != nil {
		return nil, fmt.Errorf("unable to establish connection: %v", err)
	}
	return &Store{db: client.Database(os.Getenv("MONGODB_NAME"))}, nil
}

func (s *Store) getCollection(collection string) *mongo.Collection {
	return s.db.Collection(collection)
}

func (s *Store) CreateSneaker(ctx context.Context, sneaker create.Sneaker) (*create.Sneaker, error) {
	if inserted, err := upsert(ctx, s.getCollection("sneakers"), bson.D{{"sku", sneaker.Sku}}, sneaker); err != nil {
		return nil, err
	} else if !inserted {
		return nil, errors.New("sneaker was not inserted")
	}
	return &sneaker, nil
}

func (s *Store) GetSneakerByModel(ctx context.Context, model string) (get.Sneaker, error) {

	var result get.Sneaker
	if err := s.getCollection("sneakers").FindOne(ctx, bson.D{{"model", model}}).Decode(&result); err != nil {
		return get.Sneaker{}, err
	}
	return result, nil

}
func (s *Store) GetSneakerBySKU(ctx context.Context, sku string) (get.Sneaker, error) {

	var result get.Sneaker
	if err := s.getCollection("sneakers").FindOne(ctx, bson.D{{"sku", sku}}).Decode(&result); err != nil {
		return get.Sneaker{}, err
	}

	return result, nil

}

func (s *Store) GetAllSneakers(ctx context.Context) ([]get.Sneaker, error) {

	var sneakers []get.Sneaker
	cur, err := s.getCollection("sneakers").Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var sneaker get.Sneaker
		if err := cur.Decode(&sneaker); err != nil {
			return nil, err
		}
		sneakers = append(sneakers, sneaker)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return sneakers, nil

}

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
