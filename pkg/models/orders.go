package models

import "time"

type Order struct {
	From      string    `json:"from" bson:"from"`
	To        string    `json:"to" bson:"to"`
	Status    string    `json:"status" bson:"status"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

func (order *Order) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	// result["from"] = order.

	return result
}
