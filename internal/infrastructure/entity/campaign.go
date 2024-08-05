package entity

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Campaign struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	CampaignID      string             `bson:"campaign_id"`
	Title           string             `bson:"title"`
	Description     string             `bson:"description"`
	Goal            float64            `bson:"goal"`
	CurrentAmount   float64            `bson:"current_amount"`
	CreatorID       primitive.ObjectID `bson:"creator"`
	CreatorUsername string             `bson:"creator_username"`
	LikeCount       int                `bson:"like_count"`
	CreatedAt       time.Time          `bson:"created_at"`
}

func NewCampaign(title, description, creatorUsername string, goal float64, creatorID primitive.ObjectID) *Campaign {
	return &Campaign{
		CampaignID:      uuid.New().String(),
		Title:           title,
		Description:     description,
		Goal:            goal,
		CurrentAmount:   0,
		CreatorID:       creatorID,
		CreatorUsername: creatorUsername,
		LikeCount:       0,
		CreatedAt:       time.Now(),
	}
}
