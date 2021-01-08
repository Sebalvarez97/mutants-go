package dao

import (
	"context"
	"github.com/Sebalvarez97/mutants/api/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strconv"
	"time"
)

type MongoDaoImpl struct {
	db          string
	uri         string
	timeout     time.Duration
	minPoolSize uint64
	maxPoolSize uint64
}

func NewMongoDao(db string) interfaces.MongoDao {
	uri := "mongodb://localhost:27017"
	t, _ := time.ParseDuration("30s")
	minPool, _ := strconv.ParseUint("100", 10, 64)
	maxPool, _ := strconv.ParseUint("100", 10, 64)
	if db := os.Getenv(mongoDbEnv); db != "" {
		db = os.Getenv(mongoDbEnv)
	}
	if u := os.Getenv(mongoUriEnv); u != "" {
		uri = os.Getenv(mongoUriEnv)
	}
	if to := os.Getenv(mongoTimeOutEnv); to != "" {
		if dur, err := time.ParseDuration(to); err == nil {
			t = dur
		}
	}
	if mp := os.Getenv(mongoMinConnectionPool); mp != "" {
		if ui64, err := strconv.ParseUint(mp, 10, 64); err == nil {
			minPool = ui64
		}
	}
	if mp := os.Getenv(mongoMaxConnectionPool); mp != "" {
		if ui64, err := strconv.ParseUint(mp, 10, 64); err == nil {
			maxPool = ui64
		}
	}
	return MongoDaoImpl{
		db:          db,
		uri:         uri,
		timeout:     t,
		minPoolSize: minPool,
		maxPoolSize: maxPool,
	}
}

const mongoUriEnv = "MONGO_URI"
const mongoDbEnv = "MONGO_DB"
const mongoTimeOutEnv = "MONGO_TIMEOUT"
const mongoMinConnectionPool = "MONGO_MIN_POOL"
const mongoMaxConnectionPool = "MONGO_MAX_POOL"

func (i MongoDaoImpl) connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(i.uri)
	clientOptions.SetMinPoolSize(i.minPoolSize)
	clientOptions.SetMaxPoolSize(i.maxPoolSize)
	clientOptions.SetConnectTimeout(i.timeout)
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
