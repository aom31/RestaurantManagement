package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Note struct {
	ID        primitive.ObjectID `bson: "id"`
	Text      string             `json: "text"`
	Title     string             `json: "title"`
	Create_at time.Time          `json: "create_at"`
	Update_at time.Time          `json: "update_at"`
	Note_id   string             `json: "note_id"`
}
