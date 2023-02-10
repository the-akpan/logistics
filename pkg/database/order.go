package database

import (
	"context"
	"tracka/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const orderCollection = "orders"

var orderColl *OrderCollection

func initOrders(db *mongo.Database) {
	collection := db.Collection(orderCollection)
	orderColl = &OrderCollection{collection}
}

type OrderCollection struct {
	*mongo.Collection
}

func (coll *OrderCollection) GetOrders(page int64, limit int64) ([]*models.Order, error) {
	var orders []*models.Order = make([]*models.Order, 0)
	skip := page - 1*limit
	opt := options.Find().
		SetSkip(skip).SetLimit(limit).
		SetSort(bson.D{{Key: "createdAt", Value: -1}})

	cursor, err := coll.Find(context.TODO(), bson.D{}, opt)
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (coll *OrderCollection) GetOrder(tracker string) (*models.Order, error) {
	var order models.Order

	err := coll.FindOne(context.TODO(), bson.D{{Key: "tracker", Value: tracker}}).Decode(&order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func OrderColl() *OrderCollection {
	return orderColl
}
