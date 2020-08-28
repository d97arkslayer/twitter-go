package Repositories

import (
	"context"
	"errors"
	"github.com/d97arkslayer/twitter-go/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/**
 * StoreRelation
 * Use to store in database a relation between users
 */
func StoreRelation(relation Models.Relation)(bool, error){
	collection, ctx, cancel := setupConnection("twitter-go", "relation")
	defer cancel()
	exists,_ := findRelation(relation, *collection, ctx)
	if exists == true {
		return false, errors.New("the relation already exist")
	}
	_, err := collection.InsertOne(ctx, relation)
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
 * DeleteRelation
 * Use to delete a relation between users
 */
func DeleteRelation(relation Models.Relation)(bool, error){
	collection, ctx, cancel := setupConnection("twitter-go", "relation")
	defer cancel()
	exists, err := findRelation(relation, *collection, ctx)
	if err != nil {
		return false, err
	}
	if exists == false {
		return false, errors.New("the relation not exists")
	}
	_, err = collection.DeleteOne(ctx, relation)
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
 * findRelation
 * use to find an existent relation
 */
func findRelation(relation Models.Relation, collection mongo.Collection,ctx context.Context)(bool, error){
	filter := bson.M{"userId": bson.M{"$eq": relation.UserId}, "userRelationId": bson.M{"$eq": relation.UserRelationId}}
	var existRelation Models.Relation
	err := collection.FindOne(ctx,filter).Decode(&existRelation)
	if err != nil {
		return false, err
	}
	if existRelation.UserRelationId != "" {
		return true, nil
	}
	return false, nil
}