package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Balance  float64            `bson:"balance"`
}

func NewUser(username, email, password string, balance float64) *User {
	return &User{
		Username: username,
		Email:    email,
		Password: password,
		Balance:  balance,
	}
}
