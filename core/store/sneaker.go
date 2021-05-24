package store

import (
	"context"
	"log"
	"main/models"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SneakerStore struct {
	DB    *sqlx.DB
	mongo *mongo.Collection
}

type sneaker struct {
	price int64          `db:"price"`
	brand string         `db:"brand"`
	model string         `db:"model"`
	sku   string         `db:"sku"`
	sites pq.StringArray `db:"sites"`
}

func NewSneakerStore(conn *DatabaseConnection) *SneakerStore {
	if conn != nil && conn.Connection != "" {
		db, err := sqlx.Connect("postgres", conn.Connection)
		if err != nil {
			log.Println("unable to establish connection", err)
			return nil
		}
		return &SneakerStore{
			DB: db,
		}
	} else {
		if conn == nil || conn.MongoDB == nil {
			return nil
		}

		return &SneakerStore{
			mongo: conn.MongoDB.Collection("sneakers"),
		}
	}

}

func (s SneakerStore) CreateSneaker(ctx context.Context, model, brand, sku string, photos []string, siteSizePrice models.SiteSizePrice) (*models.Sneaker, error) {

	sneaker := models.Sneaker{
		Brand:            brand,
		Model:            model,
		Sku:              sku,
		SitesSizesPrices: &siteSizePrice,
	}

	if s.DB != nil {
		//
	}
	if s.mongo != nil {
		if _, err := upsert(ctx, s.mongo, bson.D{{"sku", sneaker.Sku}}, sneaker); err != nil {
			return nil, err
		}
	}
	return &sneaker, nil
}

func (s SneakerStore) GetAllSneakers(ctx context.Context) ([]models.Sneaker, error) {

	if s.DB != nil {
		//get values from db scan into store sneaker type
		var sneakers []sneaker
		if err := s.DB.Select(&sneakers, `SELECT * FROM sneakers`); err != nil {
			return nil, err
		}

		//convert and return array of model type
		var converted []models.Sneaker
		for _, s := range sneakers {
			converted = append(converted, s.ToSneaker())
		}

		return converted, nil
	} else {

		var sneakers []models.Sneaker
		cur, err := s.mongo.Find(ctx, bson.D{{}})
		if err != nil {
			return nil, err
		}
		defer cur.Close(ctx)

		for cur.Next(ctx) {
			var sneaker models.Sneaker
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

}

func (s SneakerStore) GetSneakerByModel(ctx context.Context, model string) (models.Sneaker, error) {
	var sneaker sneaker
	if s.DB != nil {
		if err := s.DB.Get(&sneaker, `SELECT * FROM sneakers WHERE model=$1`, model); err != nil {
			return models.Sneaker{}, err
		}
		return sneaker.ToSneaker(), nil
	} else {

		var result models.Sneaker
		if err := s.mongo.FindOne(ctx, bson.D{{"model", model}}).Decode(&result); err != nil {
			return models.Sneaker{}, err
		}

		return result, nil

	}

}

func (s sneaker) ToSneaker() models.Sneaker {

	//convert values returned from db to site_sold_on type so it's methods can be used
	return models.Sneaker{
		Brand: s.brand,
		Model: s.model,
		Sku:   s.sku,
	}
}
