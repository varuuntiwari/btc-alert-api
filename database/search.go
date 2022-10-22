package database

import (
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SearchAlerts(email string) ([]Record, error) {
	ctx := context.TODO()
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(URI).SetServerAPIOptions(serverAPIOptions)
	cli, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	collection := cli.Database("alertDB").Collection("users")
	model := mongo.IndexModel{
		Keys: bson.M{"email": "text"},
	}
	idx, err := collection.Indexes().CreateOne(ctx, model)
	if err != nil {
		return nil, err
	} else {
		log.Println("Index name", idx)
	}
	cursor, err := collection.Find(ctx, bson.D{
		{"$text", bson.D{
			{"$search", email},
		}},
	})
	if err != nil {
		return nil, err
	}
	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	var res []Record
	for _, result := range results {
		var tmp map[string]interface{}
		var tmpBytes []byte
		tmpBytes, _ = bson.MarshalExtJSON(result, true, true)
		_ = json.Unmarshal(tmpBytes, &tmp)
		b, err := json.Marshal(tmp["priceCheck"])
		if err != nil {
			return nil, err
		}
		price := interface{}(string(b))
		log.Printf("%+v\n", price)
		res = append(res, Record{
			Email:        tmp["email"].(string),
			PriceToCheck: 0,
			Status:       tmp["stat"].(string),
		})
		log.Println(tmp)
	}

	return res, nil
}

func SearchPrice(price float64) (emails []string) {
	ctx := context.TODO()
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(URI).SetServerAPIOptions(serverAPIOptions)
	cli, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil
	}
	collection := cli.Database("alertDB").Collection("users")
	filter := bson.D{{ "priceCheck", price }}
	data, err := collection.Find(ctx, filter)
	log.Println(data)
	return
}