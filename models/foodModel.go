package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Food struct {
	ID         primitive.ObjectID `bson: "id"`
	Name       *string            `json: "name" validate: "required, min = 2, max= 100"`
	Price      *float64           `json: "price" validate: "required"`
	Food_image *string            `json: "food_image" validate: "required"`
	Create_at  time.Time          `json: "create_at"`
	Update_at  time.Time          `json: "updated_id"`
	Food_id    string             `json: "food_id"`
	Menu_id    *string            `json: "menu_id" validate: "required"`
}
