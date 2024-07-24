package services

import (
	"context"
	"copilot/pkg/copilot/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AbastecimentoService struct {
	collection *mongo.Collection
}

func NewAbastecimentoService(client *mongo.Client, dbName string, collectionName string) *AbastecimentoService {
	collection := client.Database(dbName).Collection(collectionName)
	return &AbastecimentoService{collection}
}

func (s *AbastecimentoService) GetAbastecimentoByCustomer(customer string) ([]domain.Fills, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(customer)
	filter := bson.M{"customer": objectID}
	cursor, err := s.collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var abastecimentos []domain.Fills
	if err = cursor.All(ctx, &abastecimentos); err != nil {
		return nil, err
	}

	return abastecimentos, nil
}
