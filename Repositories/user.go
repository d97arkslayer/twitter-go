package Repositories

import (
	"fmt"
	"github.com/d97arkslayer/twitter-go/Models"
	"github.com/d97arkslayer/twitter-go/Utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/**
 * InsertUser
 * This function inserts a new user in mongo
 */
func InsertUser(user Models.User)(string, bool, error){
	collection, ctx, cancel := setupConnection("twitter-go", "users")
	defer cancel()
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
	collection, ctx, cancel := setupConnection("twitter-go", "users")
	defer cancel()
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
	collection, ctx, cancel := setupConnection("twitter-go", "users")
	defer cancel()
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

/**
 * UpdateUser
 * Use this to update a user record
 */
func UpdateUser(u Models.User, Id string) (bool, error){
	collection, ctx, cancel := setupConnection("twitter-go", "users")
	defer cancel()
	register := make(map[string]interface{})
	if len(u.Name) > 0 {
		register["name"] = u.Name
	}
	if len(u.Password) > 0 {
		password,_ := Utils.EncryptPassword(u.Password)
		register["password"] = password
	}
	if len(u.Lastname) > 0 {
		register["lastname"] = u.Lastname
	}
	register["birthDate"] = u.BirthDate
	if len(u.Avatar) > 0 {
		register["Avatars"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		register["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		register["location"] = u.Location
	}
	if len(u.Website) > 0 {
		register["website"] = u.Website
	}
	updateString := bson.M{
		"$set":register,
	}
	objId, _ := primitive.ObjectIDFromHex(Id)
	filter := bson.M{"_id": bson.M{"$eq":objId}}

	_, err := collection.UpdateOne(ctx,filter,updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
