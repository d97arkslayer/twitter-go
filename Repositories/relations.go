package Repositories

import (
	"errors"
	"github.com/d97arkslayer/twitter-go/Models"
	"go.mongodb.org/mongo-driver/bson"
)

/**
 * StoreRelation
 * Use to store in database a relation between users
 */
func StoreRelation(relation Models.Relation)(bool, error){
	collection, ctx, cancel := setupConnection("twitter-go", "relation")
	defer cancel()
	exists,_ := FindRelation(relation)
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
	exists, err := FindRelation(relation)
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
 * FindRelation
 * use to find an existent relation
 */
func FindRelation(relation Models.Relation)(bool, error){
	collection, ctx, cancel := setupConnection("twitter-go", "relation")
	defer cancel()
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