package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Requests

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserUpdateBalanceRequest struct {
	Balance float64 `json:"balance"`
}

type DonationRequest struct {
	CampaignID primitive.ObjectID `json:"campaign_id"`
	DonorID    primitive.ObjectID `json:"donor_id"`
	Amount     float64            `json:"amount"`
}

type CampaignRequest struct {
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	GoalAmount    float64 `json:"goal_amount"`
	CurrentAmount float64 `json:"current_amount"`
}

// Responses

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type UserResponse struct {
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Logo      []byte    `json:"logo"`
	CreatedAt time.Time `json:"created_at"`
}

type DonationResponse struct {
	ID         primitive.ObjectID `json:"id"`
	CampaignID primitive.ObjectID `json:"campaign_id"`
	DonorID    primitive.ObjectID `json:"donor_id"`
	Amount     float64            `json:"amount"`
	CreatedAt  time.Time          `json:"created_at"`
}

type CampaignResponse struct {
	CampaignID      string    `json:"campaign_id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Goal            float64   `json:"goal"`
	CurrentAmount   float64   `json:"current_amount"`
	CreatorUsername string    `json:"creator"`
	LikeCount       int       `json:"like_count"`
	CreatedAt       time.Time `json:"created_at"`
}
