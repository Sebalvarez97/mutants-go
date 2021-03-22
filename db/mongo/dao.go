package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

type MongoDao interface {
	connect(ctx context.Context) (*mongo.Client, error)
	disconnect(ctx context.Context, client *mongo.Client) error
	FindAll(ctx context.Context, filter bson.D, collectionName string) (*mongo.Cursor, error)
	CountForCollection(ctx context.Context, filter bson.D, collectionName string) (int64, error)
	FindMany(ctx context.Context, limit int64, filter bson.D, collectionName string) (*mongo.Cursor, error)
	FindOne(ctx context.Context, filter bson.D, collectionName string) (*mongo.SingleResult, error)
	InsertOne(ctx context.Context, document interface{}, collectionName string) error
	UpdateOne(ctx context.Context, filter bson.D, update bson.D, collectionName string) error
}

type dao struct {
	db      string
	Options *options.ClientOptions
}

func NewMongoDao(config Config) MongoDao {
	uri := config.Uri
	db := config.Db
	clientOptions := options.Client().ApplyURI(uri)
	if t, err := time.ParseDuration(config.Timeout); err == nil {
		clientOptions.SetConnectTimeout(t)
	}
	if minPool, err := strconv.ParseUint(config.MinPool, 10, 64); err == nil {
		clientOptions.SetMinPoolSize(minPool)
	}
	if maxPool, err := strconv.ParseUint(config.MaxPool, 10, 64); err == nil {
		clientOptions.SetMaxPoolSize(maxPool)
	}
	return &dao{
		db:      db,
		Options: clientOptions,
	}
}

func (d *dao) connect(ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, d.Options)
	return client, err
}

func (d *dao) disconnect(ctx context.Context, client *mongo.Client) error {
	err := client.Disconnect(ctx)
	return err
}

func (d *dao) FindMany(ctx context.Context, limit int64, filter bson.D, collectionName string) (*mongo.Cursor, error) {
	client, conErr := d.connect(ctx)
	if conErr != nil {
		return nil, conErr
	} else {
		collection := client.Database(d.db).Collection(collectionName)
		findOptions := options.Find()
		findOptions.SetLimit(limit)
		cur, err := collection.Find(context.TODO(), filter, findOptions)
		if err != nil {
			disconErr := d.disconnect(ctx, client)
			if disconErr != nil {
				return cur, disconErr
			}
			return cur, err
		}
		curErr := cur.Err()
		if curErr != nil {
			disconErr := d.disconnect(ctx, client)
			if disconErr != nil {
				return cur, disconErr
			}
			return cur, curErr
		}
		err = d.disconnect(ctx, client)
		return cur, err
	}
}

func (d *dao) FindAll(ctx context.Context, filter bson.D, collectionName string) (*mongo.Cursor, error) {
	client, conErr := d.connect(ctx)
	if conErr != nil {
		return nil, conErr
	} else {
		collection := client.Database(d.db).Collection(collectionName)
		cur, err := collection.Find(context.TODO(), filter)
		if err != nil {
			disconErr := d.disconnect(ctx, client)
			if disconErr != nil {
				return cur, disconErr
			}
			return cur, err
		}
		curErr := cur.Err()
		if curErr != nil {
			disconErr := d.disconnect(ctx, client)
			if disconErr != nil {
				return cur, disconErr
			}
			return cur, curErr
		}
		err = d.disconnect(ctx, client)
		return cur, err
	}
}

func (d *dao) CountForCollection(ctx context.Context, filter bson.D, collectionName string) (int64, error) {
	client, conErr := d.connect(ctx)
	if conErr != nil {
		return 0, conErr
	} else {
		collection := client.Database(d.db).Collection(collectionName)
		count, err := collection.CountDocuments(context.TODO(), filter)
		if err != nil {
			disconErr := d.disconnect(ctx, client)
			if disconErr != nil {
				return count, disconErr
			}
			return count, err
		}
		err = d.disconnect(ctx, client)
		return count, err
	}

}

func (d *dao) InsertOne(ctx context.Context, document interface{}, collectionName string) error {
	client, conErr := d.connect(ctx)
	if conErr != nil {
		return conErr
	} else {
		collection := client.Database(d.db).Collection(collectionName)
		_, err := collection.InsertOne(context.TODO(), document)
		if err != nil {
			disconErr := d.disconnect(ctx, client)
			if disconErr != nil {
				return disconErr
			}
			return err
		}
		disconErr := d.disconnect(ctx, client)
		return disconErr
	}
}

func (d *dao) InsertMany(ctx context.Context, document []interface{}, collectionName string) error {
	client, conErr := d.connect(ctx)
	if conErr != nil {
		return conErr
	} else {
		collection := client.Database(d.db).Collection(collectionName)
		_, err := collection.InsertMany(context.TODO(), document)
		if err != nil {
			disconErr := d.disconnect(ctx, client)
			if disconErr != nil {
				return disconErr
			}
		}
		return err
	}
}

func (d *dao) UpdateOne(ctx context.Context, filter bson.D, update bson.D, collectionName string) error {
	client, conErr := d.connect(ctx)
	if conErr != nil {
		return conErr
	} else {
		collection := client.Database(d.db).Collection(collectionName)
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			disconErr := d.disconnect(ctx, client)
			if disconErr != nil {
				return disconErr
			}
		}
		return err
	}
}

func (d *dao) FindOne(ctx context.Context, filter bson.D, collectionName string) (*mongo.SingleResult, error) {
	client, conErr := d.connect(ctx)
	if conErr != nil {
		return nil, conErr
	} else {
		collection := client.Database(d.db).Collection(collectionName)
		result := collection.FindOne(context.TODO(), filter)
		resultErr := result.Err()
		if resultErr != nil {
			disconErr := d.disconnect(ctx, client)
			if disconErr != nil {
				return result, disconErr
			}
		}
		return result, resultErr
	}
}

func (d *dao) Delete(ctx context.Context, filter bson.D, collectionName string) error {
	client, conErr := d.connect(ctx)
	if conErr != nil {
		return conErr
	} else {
		collection := client.Database(d.db).Collection(collectionName)
		_, err := collection.DeleteMany(context.TODO(), filter)
		if err != nil {
			disconErr := d.disconnect(ctx, client)
			if disconErr != nil {
				return disconErr
			}
		}
		return err
	}
}
