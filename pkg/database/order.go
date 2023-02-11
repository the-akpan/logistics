package database

import (
	"context"
	"fmt"
	"log"
	"tracka/pkg/models"
	"tracka/pkg/utils"

	"github.com/biter777/countries"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const orderCollection = "orders"

var orderColl *OrderCollection
var MODES []string
var STATUS []string

func initOrders(db *mongo.Database) {
	collection := db.Collection(orderCollection)
	orderColl = &OrderCollection{collection}

	MODES = make([]string, 0)
	MODES = append(MODES, "SHIP", "AIR", "TRAIN")

	STATUS = make([]string, 0)
	STATUS = append(STATUS, "CREATED", "IN_TRANSIT", "ON_HOLD", "DELIVERED")
}

type OrderCollection struct {
	*mongo.Collection
}

func (coll *OrderCollection) GetOrders(page int64, limit int64) ([]*models.Order, error) {
	var orders []*models.Order = make([]*models.Order, 0)
	skip := (page - 1) * limit
	opt := options.Find().
		SetSkip(skip).
		SetLimit(limit).
		SetSort(bson.D{{Key: "createdAt", Value: -1}}).
		SetProjection(bson.D{{Key: "updates", Value: 0}})

	cursor, err := coll.Find(context.TODO(), bson.D{{}}, opt)
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
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Fatal(err)
	}

	return &order, nil
}

func (coll *OrderCollection) CreateOrder(data *models.OrderIn) (*models.Order, error) {
	var order models.Order
	from := countries.ByName(data.From).Alpha3()
	to := countries.ByName(data.To).Alpha3()
	track_num := utils.GenerateNumericString(10)
	tracker := fmt.Sprintf("%s%s%s", from, track_num, to)

	err := coll.FindOne(context.TODO(), bson.D{{Key: "tracker", Value: tracker}}).Decode(&order)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return coll.CreateOrder(data)
		}
	}

	order = models.Order{
		From:      data.From,
		To:        data.To,
		Mode:      data.Mode,
		Status:    STATUS[0],
		CreatedAt: data.CreatedAt,
		Tracker:   tracker,
	}

	_, err = coll.InsertOne(context.TODO(), order)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (coll *OrderCollection) UpdateOrder(order *models.Order) error {
	_, err := coll.UpdateOne(context.TODO(), bson.D{{Key: "tracker", Value: order.Tracker}}, bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: order.Status}}}})
	return err
}

func OrderColl() *OrderCollection {
	return orderColl
}
