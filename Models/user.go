package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)


/**
 * User
 * Model user
 */
type User struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name" json:"name,omitempty"`
	Lastname string `bson:"lastname" json:"lastname,omitempty"`
	BirthDate time.Time `bson:"birthDate" json:"birthDate,omitempty"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password,omitempty"`
	Avatar string `bson:"Avatars" json:"Avatars,omitempty"`
	Banner string `bson:"banner" json:"banner,omitempty"`
	Biography string `bson:"biography" json:"biography,omitempty"`
	Location string `bson:"location" json:"location,omitempty"`
	Website string `bson:"website" json:"website,omitempty"`
}