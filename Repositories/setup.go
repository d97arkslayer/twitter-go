package Repositories

import (
	"context"
	"github.com/d97arkslayer/twitter-go/Database"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

/**
 * setupConnection
 * Use this to do operations over collection Users
 */
func setupConnection(db string, coll string) (*mongo.Collection, context.Context, context.CancelFunc){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	database := Database.MongoConnection.Database(db)
	collection := database.Collection(coll)
	return collection, ctx, cancel
}