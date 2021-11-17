package connectionhelper

import (
	"context"
	"fmt"
	"sync"

	"github.com/stclaird/greenmap/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Used to create a singleton object of MongoDB client.
Initialized and exposed through  GetMongoClient().*/
var clientInstance *mongo.Client

//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

//Used to execute client creation procedure only once.
var mongoOnce sync.Once

var mongoconf = config.GetMongoConfig()

var credential = options.Credential{
	Username: mongoconf.User,
	Password: mongoconf.Password,
}

//I have used below constants just to hold required database config's.

var CONNECTIONSTRING = mongoconf.Connection_String
var DB = mongoconf.Db
var LOCATIONS = mongoconf.Collection

//GetMongoClient - Return mongodb connection to work with
func GetMongoClient() (*mongo.Client, error) {
	//Perform connection creation operation only once.

	fmt.Printf("CONNECTIONSTRING %s", CONNECTIONSTRING)

	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING).SetAuth(credential)
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}
