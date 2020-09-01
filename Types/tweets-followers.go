package Types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TweetsFollowers struct {
	Id primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId string `bson:"userId" json:"userId,omitempty"`
	UserRelationId string `bson:"userRelationId" json:"userRelationId,omitempty"`
	Tweet struct{
		Message string `bson:"message" json:"message,omitempty"`
		Date time.Time `bson:"date" json:"date,omitempty"`
		id string `bson:"_id" json:"_id,omitempty"`
	}
}