package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Table struct {
	ID               primitive.ObjectID `bson:"id"`
	Number_of_guests *int               `json: "number_of_guests" validate: "required"`
	Table_number     *int               `json: "table_number" validate: "required"`
	Create_at        time.Time          `json : "create_at"`
	Update_at        time.Time          `json: "update_at"`
	Table_id         string             `json: "table_id"`
}
