package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Record struct {
	Email        string `bson:"email"`
	PriceToCheck int64  `bson:"priceCheck"`
	Status       string `bson:"stat"`
}

const URI string = "mongodb+srv://alertdb.ydzzrth.mongodb.net/?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority&tlsCertificateKeyFile=/home/retr0/Github-repos/btc-alert-api/mongoAccess.pem"

func AddAlert(email string, priceCheck int64) error {
	newAlert := Record{
		Email:        email,
		PriceToCheck: priceCheck,
		Status:       "active",
	}
	ctx := context.TODO()
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(URI).SetServerAPIOptions(serverAPIOptions)
	cli, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	defer cli.Disconnect(ctx)

	collection := cli.Database("alertDB").Collection("users")
	res, err := collection.InsertOne(ctx, newAlert)
	if err != nil {
		return err
	}
	log.Printf("added alert with ID %v", res.InsertedID)
	return nil
}

func DelAlert(email string) error {
	filter := bson.D{{"email", email}}

	ctx := context.TODO()
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(URI).SetServerAPIOptions(serverAPIOptions)
	cli, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	defer cli.Disconnect(ctx)

	collection := cli.Database("alertDB").Collection("users")

	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 1 {
		log.Printf("deleted alert for email %v", email)
	} else {
		log.Printf("no document found with email %v", email)
	}
	return nil
}
