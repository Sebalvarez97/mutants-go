package interfaces

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDao interface {
	FindAll(filter bson.D, collectionName string) (*mongo.Cursor, error)
	CountForCollection(filter bson.D, collectionName string) (int64, error)
	FindMany(limit int64, filter bson.D, collectionName string) (*mongo.Cursor, error)
	FindOne(filter bson.D, collectionName string) (*mongo.SingleResult, error)
	InsertOne(document interface{}, collectionName string) error
	UpdateOne(filter bson.D, update bson.D, collectionName string) error
}
