package location

import (
	"context"
	"fmt"

	"github.com/stclaird/greenmap/connectionhelper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Location struct {
	Id                    primitive.ObjectID `bson:"_id"`
	LocationType          string             `bson:"LocationType"`
	LocationName          string             `bson:"LocationName"`
	AddressFirstLine      string             `bson:"AddressFirstLine"`
	AddressTown           string             `bson:"AddressTown"`
	AddressPostCode       string             `bson:"AddressPostCode"`
	AddressState          string             `bson:"AddressState"`
	AddressCountry        string             `bson:"AddressCountry"`
	AddressCountryCodeISO string             `bson:"AddressCountryCodeISO"`
	LocationPoint         LocationPoint      `bson:"LocationPoint"`

	Residential int `bson:"Residential"`
}

type LocationPoint struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

func GetAllLocations() ([]Location, error) {
	//Define filter query for fetching specific document from collection
	filter := bson.D{{}} //bson.D{{}} specifies 'all documents'
	locations := []Location{}
	//Get MongoDB connection using connectionhelper.
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		fmt.Printf("err %i", err)
		return locations, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(connectionhelper.DB).Collection(connectionhelper.LOCATIONS)
	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return locations, findError
	}
	//Map result to slice
	for cur.Next(context.TODO()) {
		t := Location{}
		err := cur.Decode(&t)
		if err != nil {
			return locations, err
		}
		locations = append(locations, t)
	}
	// once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(locations) == 0 {
		return locations, mongo.ErrNoDocuments
	}
	return locations, nil
}
