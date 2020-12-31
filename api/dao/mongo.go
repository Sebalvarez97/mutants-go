package dao

import (
	"context"
	"github.com/Sebalvarez97/mutants/api/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type MongoDaoImpl struct {
	db  string
	uri string
}

func NewMongoDao(db string) interfaces.MongoDao {
	return MongoDaoImpl{
		db:  db,
		uri: "mongodb://localhost:27017",
	}
}

const mongoUriEnv = "MONGO_URI"
const mongoDbEnv = "MONGO_DB"

func (i MongoDaoImpl) init() {
	if db := os.Getenv(mongoDbEnv); db != "" {
		i.db = os.Getenv(mongoDbEnv)
	}
	if u := os.Getenv(mongoUriEnv); u != "" {
		i.uri = os.Getenv(mongoUriEnv)
	}
}

func (i MongoDaoImpl) connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(i.uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	return client, err
}

func (i MongoDaoImpl) disconnect(client *mongo.Client) error {
	err := client.Disconnect(context.TODO())
	return err
}

func (i MongoDaoImpl) FindMany(limit int64, filter bson.D, collectionName string) (*mongo.Cursor, error) {
	client, conErr := i.connect()
	if conErr != nil {
		return nil, conErr
	} else {
		collection := client.Database(i.db).Collection(collectionName)
		findOptions := options.Find()
		findOptions.SetLimit(limit)
		cur, err := collection.Find(context.TODO(), filter, findOptions)
		if err != nil {
			disconErr := i.disconnect(client)
			if disconErr != nil {
				return cur, disconErr
			}
			return cur, err
		}
		curErr := cur.Err()
		if curErr != nil {
			disconErr := i.disconnect(client)
			if disconErr != nil {
				return cur, disconErr
			}
			return cur, curErr
		}
		err = i.disconnect(client)
		return cur, err
	}
}

func (i MongoDaoImpl) FindAll(filter bson.D, collectionName string) (*mongo.Cursor, error) {
	client, conErr := i.connect()
	if conErr != nil {
		return nil, conErr
	} else {
		collection := client.Database(i.db).Collection(collectionName)
		cur, err := collection.Find(context.TODO(), filter)
		if err != nil {
			disconErr := i.disconnect(client)
			if disconErr != nil {
				return cur, disconErr
			}
			return cur, err
		}
		curErr := cur.Err()
		if curErr != nil {
			disconErr := i.disconnect(client)
			if disconErr != nil {
				return cur, disconErr
			}
			return cur, curErr
		}
		err = i.disconnect(client)
		return cur, err
	}
}

func (i MongoDaoImpl) InsertOne(document interface{}, collectionName string) error {
	client, conErr := i.connect()
	if conErr != nil {
		return conErr
	} else {
		collection := client.Database(i.db).Collection(collectionName)
		_, err := collection.InsertOne(context.TODO(), document)
		if err != nil {
			disconErr := i.disconnect(client)
			if disconErr != nil {
				return disconErr
			}
			return err
		}
		disconErr := i.disconnect(client)
		return disconErr
	}
}

func (i MongoDaoImpl) InsertMany(document []interface{}, collectionName string) error {
	client, conErr := i.connect()
	if conErr != nil {
		return conErr
	} else {
		collection := client.Database(i.db).Collection(collectionName)
		_, err := collection.InsertMany(context.TODO(), document)
		if err != nil {
			disconErr := i.disconnect(client)
			if disconErr != nil {
				return disconErr
			}
		}
		return err
	}
}

func (i MongoDaoImpl) UpdateOne(filter bson.D, update bson.D, collectionName string) error {
	client, conErr := i.connect()
	if conErr != nil {
		return conErr
	} else {
		collection := client.Database(i.db).Collection(collectionName)
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			disconErr := i.disconnect(client)
			if disconErr != nil {
				return disconErr
			}
		}
		return err
	}
}

func (i MongoDaoImpl) FindOne(filter bson.D, collectionName string) (*mongo.SingleResult, error) {
	client, conErr := i.connect()
	if conErr != nil {
		return nil, conErr
	} else {
		collection := client.Database(i.db).Collection(collectionName)
		result := collection.FindOne(context.TODO(), filter)
		resultErr := result.Err()
		if resultErr != nil {
			disconErr := i.disconnect(client)
			if disconErr != nil {
				return result, disconErr
			}
		}
		return result, resultErr
	}
}

func (i MongoDaoImpl) Delete(filter bson.D, collectionName string) error {
	client, conErr := i.connect()
	if conErr != nil {
		return conErr
	} else {
		collection := client.Database(i.db).Collection(collectionName)
		_, err := collection.DeleteMany(context.TODO(), filter)
		if err != nil {
			disconErr := i.disconnect(client)
			if disconErr != nil {
				return disconErr
			}
		}
		return err
	}
}
