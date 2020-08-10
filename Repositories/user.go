package Repositories

import (
	"context"
	"fmt"
	"github.com/d97arkslayer/twitter-go/Database"
	"github.com/d97arkslayer/twitter-go/Models"
	"github.com/d97arkslayer/twitter-go/Utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/**
 * InsertUser
 * This function inserts a new user in mongo
 */
func InsertUser(user Models.User)(string, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	database := Database.MongoConnection.Database("twitter-go")
	collection := database.Collection("users")
	user.Password, _ = Utils.EncryptPassword(user.Password)
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}
	ObjId, _ := result.InsertedID.(primitive.ObjectID)
	return ObjId.String(), true, nil
}

/**
 * ExistUser
 * This function search an user with the provided email
 */
func ExistUser(email string)(Models.User, bool, string){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	database := Database.MongoConnection.Database("twitter-go")
	collection := database.Collection("users")
	condition := bson.M{"email":email}
	var result Models.User
	err := collection.FindOne(ctx, condition).Decode(&result);
	Id := result.Id.Hex()
	if err != nil {
		return result, false, Id
	}
	return result, true, Id
}

/**
 * Login
 * Use to log user
 */
func Login(email string, password string) (Models.User, bool)  {
	user, exist, _ := ExistUser(email)
	if exist == false {
		return user, false
	}
	compare := Utils.ComparePassword(password, user.Password)
	if compare == false {
		return user, false
	}
	return user, true
}

/**
 * ShowProfile
 * Use to find a user profile
 */
func ShowProfile(Id string)(Models.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	database := Database.MongoConnection.Database("twitter-go")
	collection := database.Collection("users")
	var profile Models.User
	objId, _ := primitive.ObjectIDFromHex(Id)
	condition := bson.M{ "_id":objId }
	err := collection.FindOne(ctx, condition).Decode(&profile)
	profile.Password=""
	if err != nil {
		fmt.Println("Profile not found " + err.Error())
		return profile, err
	}
	return profile, nil
}
