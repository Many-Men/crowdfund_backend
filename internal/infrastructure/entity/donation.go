package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Donation struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Amount   float64            `bson:"amount"`
	Campaign primitive.ObjectID `bson:"campaign"`
	Donor    primitive.ObjectID `bson:"donor"`
	Date     time.Time          `bson:"date"`
}

func NewDonation(amount float64, campaign, donor primitive.ObjectID) *Donation {
	return &Donation{
		Amount:   amount,
		Campaign: campaign,
		Donor:    donor,
		Date:     time.Now(),
	}
}
