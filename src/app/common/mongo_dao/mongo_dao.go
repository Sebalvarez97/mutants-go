package mongo_dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

const MongoNotFoundErr = "mongo: no documents in result"
const mongoUriEnv = "MONGO_URI"
const defaultMongoUri = "mongodb://localhost:27017"

type MongoDao struct {
	Db string
}

func NewMongoDao(db string) *MongoDao {
	return &MongoDao{Db: db}
}

func (i MongoDao) connect() (*mongo.Client, error) {
	uri := defaultMongoUri
	if u := os.Getenv(mongoUriEnv); u != "" {
		uri = os.Getenv(mongoUriEnv)
	}
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	return client, err
}

func disconnect(client *mongo.Client) error {
	err := client.Disconnect(context.TODO())
	return err
}

func (i MongoDao) FindMany(limit int64, filter bson.D, collectionName string) (*mongo.Cursor, error) {
	client, conErr := i.connect()
	if conErr != nil {
		return nil, conErr
	} else {
		collection := client.Database(i.Db).Collection(collectionName)
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

func (i MongoDao) InsertOne(document interface{}, collectionName string) error {
	client, conErr := i.connect()
	if conErr != nil {
		return conErr
	} else {
		collection := client.Database(i.Db).Collection(collectionName)
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

func (i MongoDao) InsertMany(document []interface{}, collectionName string) error {
	client, conErr := i.connect()
	if conErr != nil {
		return conErr
	} else {
		collection := client.Database(i.Db).Collection(collectionName)
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

func (i MongoDao) UpdateOne(filter bson.D, update bson.D, db string, collectionName string) error {
	client, conErr := i.connect()
	if conErr != nil {
		return conErr
	} else {
		collection := client.Database(i.Db).Collection(collectionName)
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

func (i MongoDao) FindOne(filter bson.D, collectionName string) (*mongo.SingleResult, error) {
	client, conErr := i.connect()
	if conErr != nil {
		return nil, conErr
	} else {
		collection := client.Database(i.Db).Collection(collectionName)
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

func (i MongoDao) Delete(filter bson.D, collectionName string) error {
	client, conErr := i.connect()
	if conErr != nil {
		return conErr
	} else {
		collection := client.Database(i.Db).Collection(collectionName)
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
