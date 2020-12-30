package repository

import (
	"context"
	. "github.com/Sebalvarez97/mutants/src/main/bundle/mutant/domain"
	"github.com/Sebalvarez97/mutants/src/main/common/errors"
	dao "github.com/Sebalvarez97/mutants/src/main/common/mongo_dao"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type MongoDao interface {
	FindAll(filter bson.D, collectionName string) (*mongo.Cursor, error)
	FindMany(limit int64, filter bson.D, collectionName string) (*mongo.Cursor, error)
	FindOne(filter bson.D, collectionName string) (*mongo.SingleResult, error)
	InsertOne(document interface{}, collectionName string) error
	UpdateOne(filter bson.D, update bson.D, collectionName string) error
}

type MongoDaoImpl struct{}

func (m MongoDaoImpl) FindAll(filter bson.D, collectionName string) (*mongo.Cursor, error) {
	return dao.FindAll(filter, collectionName)
}

func (m MongoDaoImpl) FindMany(limit int64, filter bson.D, collectionName string) (*mongo.Cursor, error) {
	return dao.FindMany(limit, filter, collectionName)
}

func (m MongoDaoImpl) FindOne(filter bson.D, collectionName string) (*mongo.SingleResult, error) {
	return dao.FindOne(filter, collectionName)
}

func (m MongoDaoImpl) InsertOne(document interface{}, collectionName string) error {
	return dao.InsertOne(document, collectionName)
}

func (m MongoDaoImpl) UpdateOne(filter bson.D, update bson.D, collectionName string) error {
	return dao.UpdateOne(filter, update, collectionName)
}

var MongoMutantsDao MongoDao

func init() {
	MongoMutantsDao = MongoDaoImpl{}
}

const collection = "dna"
const MongoNotFoundErr = "mongo: no documents in result"

func FindAll() ([]Dna, *errors.ApiErrorImpl) {
	dna, err := MongoMutantsDao.FindAll(bson.D{}, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return nil, &apiErr
	}
	return mapToObjects(dna)
}

func FindMany() ([]Dna, *errors.ApiErrorImpl) {
	dna, err := MongoMutantsDao.FindMany(10, bson.D{}, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return nil, &apiErr
	}
	return mapToObjects(dna)
}

func FindAllMutants() ([]Dna, *errors.ApiErrorImpl) {
	filter := bson.D{{"is_mutant", true}}
	dna, err := MongoMutantsDao.FindAll(filter, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return nil, &apiErr
	}
	return mapToObjects(dna)
}

func FindById(id string) (Dna, *errors.ApiErrorImpl) {
	idObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Print(err)
		apiErr := errors.GenericError(err)
		return Dna{}, &apiErr
	}
	dna, err := MongoMutantsDao.FindOne(bson.D{{"_id", idObject}}, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		if err.Error() == MongoNotFoundErr {
			apiErr = errors.NotFoundError(err)
		}
		log.Print(apiErr.Error())
		return Dna{}, &apiErr
	}
	return mapToObject(dna)
}

func FindAllHumans() ([]Dna, *errors.ApiErrorImpl) {
	filter := bson.D{{"is_mutant", false}}
	dna, err := MongoMutantsDao.FindAll(filter, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return nil, &apiErr
	}
	return mapToObjects(dna)
}

func FindByDnaHash(hash string) (Dna, *errors.ApiErrorImpl) {
	dna, err := MongoMutantsDao.FindOne(bson.D{{"dna_hash", hash}}, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		if err.Error() == MongoNotFoundErr {
			apiErr = errors.NotFoundError(err)
		}
		log.Print(apiErr.Error())
		return Dna{}, &apiErr
	}
	return mapToObject(dna)
}

func Insert(dna *Dna) *errors.ApiErrorImpl {
	if err := MongoMutantsDao.InsertOne(dna, collection); err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return &apiErr
	}
	return nil
}

func Upsert(dna *Dna) *errors.ApiErrorImpl {
	err := Update(dna.DnaHash, dna)
	if err != nil {
		if err.Code == errors.NotFoundCode {
			return Insert(dna)
		}
	}
	return err
}

func Update(hash string, dna *Dna) *errors.ApiErrorImpl {
	_, findErr := FindByDnaHash(hash)
	if findErr != nil {
		return findErr
	}
	updateErr := MongoMutantsDao.UpdateOne(bson.D{{"dna_hash", hash}}, bson.D{{"$set", bson.D{{"chain", dna.Chain}, {"is_mutant", dna.IsMutant}, {"mutant_sequences", dna.MutantSequences}}}}, collection)
	if updateErr != nil {
		apiErr := errors.GenericError(updateErr)
		if updateErr.Error() == MongoNotFoundErr {
			apiErr = errors.NotFoundError(updateErr)
		}
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
	var dna Dna
	err := document.Decode(&dna)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return dna, &apiErr
	}
	return dna, nil
}
