package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

const mongoUriEnv = "MONGO_URI"
const defaultMongoUri = "mongodb://localhost:27017"
const mongoDbEnv = "MONGO_DB"
const defaultMongoDb = "test"

var db string
var uri string

func init() {
	db = defaultMongoDb
	if db := os.Getenv(mongoDbEnv); db != "" {
		db = os.Getenv(mongoDbEnv)
	}
	uri = defaultMongoUri
	if u := os.Getenv(mongoUriEnv); u != "" {
		uri = os.Getenv(mongoUriEnv)
	}
}

func connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	return client, err
}

func disconnect(client *mongo.Client) error {
	err := client.Disconnect(context.TODO())
	return err
}

func FindMany(limit int64, filter bson.D, collectionName string) (*mongo.Cursor, error) {
	client, conErr := connect()
	if conErr != nil {
		return nil, conErr
	} else {
		collection := client.Database(db).Collection(collectionName)
		findOptions := options.Find()
		findOptions.SetLimit(limit)
		cur, err := collection.Find(context.TODO(), filter, findOptions)
		if err != nil {
			disconErr := disconnect(client)
			if disconErr != nil {
				return cur, disconErr
			}
			return cur, err
		}
		curErr := cur.Err()
		if curErr != nil {
			disconErr := disconnect(client)
			if disconErr != nil {
				return cur, disconErr
			}
			return cur, curErr
		}
		err = disconnect(client)
		return cur, err
	}
}

func FindAll(filter bson.D, collectionName string) (*mongo.Cursor, error) {
	client, conErr := connect()
	if conErr != nil {
		return nil, conErr
	} else {
		collection := client.Database(db).Collection(collectionName)
		cur, err := collection.Find(context.TODO(), filter)
		if err != nil {
			disconErr := disconnect(client)
			if disconErr != nil {
				return cur, disconErr
			}
			return cur, err
		}
		curErr := cur.Err()
		if curErr != nil {
			disconErr := disconnect(client)
			if disconErr != nil {
				return cur, disconErr
			}
			return cur, curErr
		}
		err = disconnect(client)
		return cur, err
	}
}

func InsertOne(document interface{}, collectionName string) error {
	client, conErr := connect()
	if conErr != nil {
		return conErr
	} else {
		collection := client.Database(db).Collection(collectionName)
		_, err := collection.InsertOne(context.TODO(), document)
		if err != nil {
			disconErr := disconnect(client)
			if disconErr != nil {
				return disconErr
			}
			return err
		}
		disconErr := disconnect(client)
		return disconErr
	}
}

func InsertMany(document []interface{}, collectionName string) error {
	client, conErr := connect()
	if conErr != nil {
		return conErr
	} else {
		collection := client.Database(db).Collection(collectionName)
		_, err := collection.InsertMany(context.TODO(), document)
		if err != nil {
			disconErr := disconnect(client)
			if disconErr != nil {
				return disconErr
			}
		}
		return err
	}
}

func UpdateOne(filter bson.D, update bson.D, collectionName string) error {
	client, conErr := connect()
	if conErr != nil {
		return conErr
	} else {
		collection := client.Database(db).Collection(collectionName)
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			disconErr := disconnect(client)
			if disconErr != nil {
				return disconErr
			}
		}
		return err
	}
}

func FindOne(filter bson.D, collectionName string) (*mongo.SingleResult, error) {
	client, conErr := connect()
	if conErr != nil {
		return nil, conErr
	} else {
		collection := client.Database(db).Collection(collectionName)
		result := collection.FindOne(context.TODO(), filter)
		resultErr := result.Err()
		if resultErr != nil {
			disconErr := disconnect(client)
			if disconErr != nil {
				return result, disconErr
			}
		}
		return result, resultErr
	}
}

func Delete(filter bson.D, collectionName string) error {
	client, conErr := connect()
	if conErr != nil {
		return conErr
	} else {
		collection := client.Database(db).Collection(collectionName)
		_, err := collection.DeleteMany(context.TODO(), filter)
		if err != nil {
			disconErr := disconnect(client)
			if disconErr != nil {
				return disconErr
			}
		}
		return err
	}
}
