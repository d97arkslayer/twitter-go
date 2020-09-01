package Repositories

import (
	"context"
	"github.com/d97arkslayer/twitter-go/Models"
	"github.com/d97arkslayer/twitter-go/Types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

/**
 * StoreTweet
 * Use to store a new Tweet
 */
func StoreTweet(tweet Models.Tweet)(string, bool, error){
	collection, ctx, cancel := setupConnection("twitter-go", "tweet")
	defer cancel()
	register := bson.M{
		"userId": tweet.UserId,
		"message": tweet.Message,
		"date": tweet.Date,
	}
	result, err := collection.InsertOne(ctx, register)
	if err != nil {
		return "", false,err
	}
	objId, _ := result.InsertedID.(primitive.ObjectID)
	return objId.String(), true, nil
}

/**
 * GetTweets
 * Use to get all tweets from user
 */
func GetTweets(id string, page int64)([]*Models.Tweet, bool)  {
	collection, ctx, cancel := setupConnection("twitter-go", "tweet")
	defer cancel()
	var results []*Models.Tweet
	condition := bson.M{
		"userId":id,
	}
	findOptions := options.Find()
	findOptions.SetLimit(20)
	findOptions.SetSort(bson.D{{Key: "date", Value: -1}})
	findOptions.SetSkip((page-1)*20)
	pointer, err := collection.Find(ctx, condition, findOptions)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for pointer.Next(context.TODO()){
		var tweet Models.Tweet
		err := pointer.Decode(&tweet)
		if err != nil {
			return results, false
		}
		results = append(results, &tweet)
	}
	return results, true
}

/**
 * DeleteTweet
 * Use this function to delete a tweet
 */
func DeleteTweet(id string, userId string)(error){
	collection, ctx, cancel := setupConnection("twitter-go", "tweet")
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{
		"_id":objId,
		"userId": userId,
	}
	_, err := collection.DeleteOne(ctx, condition)
	return err

}

/**
 * TweetsFollowers
 * Use to get the tweets from the followers
 */
func TweetsFollowers(id string, page int) ([] Types.TweetsFollowers, bool) {
	collection, ctx, cancel := setupConnection("twitter-go", "relation")
	defer cancel()
	skip := (page - 1) * 20
	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match":bson.M{"userId":id}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from": "tweet",
			"localField": "userRelationId",
			"foreignField": "userId",
			"as": "tweet",
		},
	})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})
	cursor, err := collection.Aggregate(ctx, conditions)
	var tweets []Types.TweetsFollowers
	err = cursor.All(ctx, &tweets)
	if err != nil {
		return tweets, false
	}
	return tweets, true

}