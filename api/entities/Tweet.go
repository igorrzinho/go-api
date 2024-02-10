package entities

import (
	"github.com/google/uuid"
)
type Tweet struct{
	ID string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Description  string `json:"description" bson:"description"`
}

func NewTweet() *Tweet {
	tweet := Tweet {
		ID: uuid.New().String(),
	}

	return &tweet
}