package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Tweet struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId string `bson:"userId" json:"userId,omitempty"`
	Message string `bson:"message" json:"message,omitempty"`
	Date time.Time `bson:"date" json:"date,omitempty"`
}