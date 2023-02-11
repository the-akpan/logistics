package models

import "time"

type OrderIn struct {
	From      string    `json:"from" bson:"from"`
	To        string    `json:"to" bson:"to"`
	Mode      string    `json:"mode" bson:"mode"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

type Order struct {
	From      string         `json:"from" bson:"from"`
	To        string         `json:"to" bson:"to"`
	Mode      string         `json:"mode" bson:"mode"`
	Status    string         `json:"status" bson:"status"`
	Tracker   string         `json:"tracker" bson:"tracker"`
	Updates   []*OrderUpdate `json:"updates,omitempty" bson:"updates"`
	CreatedAt time.Time      `json:"createdAt" bson:"createdAt"`
}

func (order *Order) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	// result["from"] = order.

	return result
}

type OrderUpdate struct {
	Content   string    `json:"content" bson:"content"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	Location  string    `json:"country" bson:"country"`
}

func (update *OrderUpdate) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["content"] = update.Content
	result["createdAt"] = update.CreatedAt

	return result
}
