package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OrderItem struct {
	ID            primitive.ObjectID `bson:"id"`
	Quantity      *string            `json: "quantity" validate: "required, eq=S|eq=M|eq=L"`
	Unit_price    *float64           `json: "unit_price" validate: "required"`
	Created_at    time.Time          `json: "create_at"`
	Updated_at    time.Time          `json: "update_at"`
	food_id       *string            `json: "food_id" validate: "required"`
	Order_item_id string             `json: "order_item_id"`
	Order_id      string             `json: "order_id" validate: "required"`
}
