package repository

import (
	"context"
	. "github.com/Sebalvarez97/mutants/src/app/bundle/mutant/domain"
	"github.com/Sebalvarez97/mutants/src/app/common/errors"
	. "github.com/Sebalvarez97/mutants/src/app/common/mongo_dao"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type DnaRepositoryImpl struct{}

var dao = NewMongoDao("mutants")

const collection = "dna"

func (i DnaRepositoryImpl) FindAll() ([]Dna, *errors.ApiErrorImpl) {
	trainers, err := dao.FindMany(10, bson.D{}, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return nil, &apiErr
	}
	return mapToObjects(trainers)
}

func (i DnaRepositoryImpl) FindById(id string) (Dna, *errors.ApiErrorImpl) {
	idObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Print(err)
		apiErr := errors.GenericError(err)
		return Dna{}, &apiErr
	}
	trainer, err := dao.FindOne(bson.D{{"_id", idObject}}, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		if err.Error() == MongoNotFoundErr {
			apiErr = errors.NotFoundError(err)
		}
		log.Print(apiErr.Error())
		return Dna{}, &apiErr
	}
	return mapToObject(trainer)
}

func (i DnaRepositoryImpl) Insert(dna *Dna) *errors.ApiErrorImpl {
	if err := dao.InsertOne(dna, collection); err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return &apiErr
	}
	return nil
}

func mapToObjects(cur *mongo.Cursor) ([]Dna, *errors.ApiErrorImpl) {
	var results []Dna
	err := cur.All(context.TODO(), &results)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return results, &apiErr
	}
	return results, nil
}

func mapToObject(document *mongo.SingleResult) (Dna, *errors.ApiErrorImpl) {
	var trainer Dna
	err := document.Decode(&trainer)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return trainer, &apiErr
	}
	return trainer, nil
}
