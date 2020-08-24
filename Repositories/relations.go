package Repositories

import "github.com/d97arkslayer/twitter-go/Models"

/**
 * StoreRelation
 * Use to store in database a relation between users
 */
func StoreRelation(relation Models.Relation)(bool, error){
	collection, ctx, cancel := setupConnection("twitter-go", "relation")
	defer cancel()
	_, err := collection.InsertOne(ctx, relation)
	if err != nil {
		return false, err
	}
	return true, nil
}
