package dna

import (
	"context"
	"github.com/Sebalvarez97/mutants-go/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

type Dao interface {
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
	Db      string
	Options *options.ClientOptions
}

func NewMongoDao(config db.Config) Dao {
	uri := config.Uri
	database := config.Name
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
		Db:      database,
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
	}
	defer func() {
		_ = d.disconnect(ctx, client)
	}()
	collection := client.Database(d.Db).Collection(collectionName)
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	err = cur.Err()
	return cur, err

}

func (d *dao) FindAll(ctx context.Context, filter bson.D, collectionName string) (*mongo.Cursor, error) {
	client, conErr := d.connect(ctx)
	if conErr != nil {
		return nil, conErr
	}
	defer func() {
		_ = d.disconnect(ctx, client)
	}()
	collection := client.Database(d.Db).Collection(collectionName)
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	err = cur.Err()
	return cur, err

}

func (d *dao) CountForCollection(ctx context.Context, filter bson.D, collectionName string) (int64, error) {
	client, conErr := d.connect(ctx)
	if conErr != nil {
		return 0, conErr
	}
	defer func() {
		_ = d.disconnect(ctx, client)
	}()
	collection := client.Database(d.Db).Collection(collectionName)
	count, err := collection.CountDocuments(context.TODO(), filter)
	return count, err

}

func (d *dao) InsertOne(ctx context.Context, document interface{}, collectionName string) error {
	client, conErr := d.connect(ctx)
	if conErr != nil {
		return conErr
	}
	defer func() {
		_ = d.disconnect(ctx, client)
	}()
	collection := client.Database(d.Db).Collection(collectionName)
	_, err := collection.InsertOne(context.TODO(), document)
	return err

}

func (d *dao) InsertMany(ctx context.Context, document []interface{}, collectionName string) error {
	client, conErr := d.connect(ctx)
	if conErr != nil {
		return conErr
	}
	defer func() {
		_ = d.disconnect(ctx, client)
	}()
	collection := client.Database(d.Db).Collection(collectionName)
	_, err := collection.InsertMany(context.TODO(), document)
	return err

}

func (d *dao) UpdateOne(ctx context.Context, filter bson.D, update bson.D, collectionName string) error {
	client, conErr := d.connect(ctx)
	if conErr != nil {
		return conErr
	}
	defer func() {
		_ = d.disconnect(ctx, client)
	}()
	collection := client.Database(d.Db).Collection(collectionName)
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (d *dao) FindOne(ctx context.Context, filter bson.D, collectionName string) (*mongo.SingleResult, error) {
	client, conErr := d.connect(ctx)
	if conErr != nil {
		return nil, conErr
	}
	defer func() {
		_ = d.disconnect(ctx, client)
	}()
	collection := client.Database(d.Db).Collection(collectionName)
	result := collection.FindOne(context.TODO(), filter)
	resultErr := result.Err()
	return result, resultErr

}

func (d *dao) Delete(ctx context.Context, filter bson.D, collectionName string) error {
	client, conErr := d.connect(ctx)
	if conErr != nil {
		return conErr
	}
	defer func() {
		_ = d.disconnect(ctx, client)
	}()
	collection := client.Database(d.Db).Collection(collectionName)
	_, err := collection.DeleteMany(context.TODO(), filter)
	return err
}
