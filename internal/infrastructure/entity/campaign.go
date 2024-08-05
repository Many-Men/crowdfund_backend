package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Campaign struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Title         string             `bson:"title"`
	Description   string             `bson:"description"`
	Goal          float64            `bson:"goal"`
	CurrentAmount float64            `bson:"current_amount"`
	Creator       primitive.ObjectID `bson:"creator"`
	CreatedAt     time.Time          `bson:"created_at"`
}
