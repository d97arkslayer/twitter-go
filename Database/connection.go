package Database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)
/**
 * MongoConnection
 * Var to export connection
 */
var MongoConnection = ConnectDataBase()
var clientOptions = options.Client().ApplyURI("mongodb+srv://twitter:XnCpDAWAFio1lLrI@darkd-cluster.drsev.mongodb.net/twitter-go?retryWrites=true&w=majority")

/**
  * ConnectDataBase
  * Use to connect with database
 */
func ConnectDataBase() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Connection successful")
	return client
}

/**
 * CheckConnection
 * Check connection
 */
func CheckConnection() int {
	err := MongoConnection.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}

