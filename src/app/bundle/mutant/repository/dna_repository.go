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
	dna, err := dao.FindAll(bson.D{}, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return nil, &apiErr
	}
	return mapToObjects(dna)
}

func (i DnaRepositoryImpl) FindMany() ([]Dna, *errors.ApiErrorImpl) {
	dna, err := dao.FindMany(10, bson.D{}, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return nil, &apiErr
	}
	return mapToObjects(dna)
}

func (i DnaRepositoryImpl) FindAllMutants() ([]Dna, *errors.ApiErrorImpl) {
	filter := bson.D{{"is_mutant", true}}
	dna, err := dao.FindAll(filter, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return nil, &apiErr
	}
	return mapToObjects(dna)
}

func (i DnaRepositoryImpl) FindAllHumans() ([]Dna, *errors.ApiErrorImpl) {
	filter := bson.D{{"is_mutant", false}}
	dna, err := dao.FindAll(filter, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return nil, &apiErr
	}
	return mapToObjects(dna)
}

func (i DnaRepositoryImpl) FindById(id string) (Dna, *errors.ApiErrorImpl) {
	idObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Print(err)
		apiErr := errors.GenericError(err)
		return Dna{}, &apiErr
	}
	dna, err := dao.FindOne(bson.D{{"_id", idObject}}, collection)
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

func (i DnaRepositoryImpl) FindByDnaHash(hash string) (Dna, *errors.ApiErrorImpl) {
	dna, err := dao.FindOne(bson.D{{"dna_hash", hash}}, collection)
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

func (i DnaRepositoryImpl) Insert(dna *Dna) *errors.ApiErrorImpl {
	if err := dao.InsertOne(dna, collection); err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return &apiErr
	}
	return nil
}

func (i DnaRepositoryImpl) Upsert(dna *Dna) *errors.ApiErrorImpl {
	err := i.Update(dna.DnaHash, dna)
	if err != nil {
		if err.Code == errors.NotFoundCode {
			return i.Insert(dna)
		}
	}
	return err
}

func (i DnaRepositoryImpl) Update(hash string, dna *Dna) *errors.ApiErrorImpl {
	_, findErr := i.FindByDnaHash(hash)
	if findErr != nil {
		return findErr
	}
	updateErr := dao.UpdateOne(bson.D{{"dna_hash", hash}}, bson.D{{"$set", bson.D{{"chain", dna.Chain}, {"is_mutant", dna.IsMutant}, {"mutant_sequences", dna.MutantSequences}}}}, collection)
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
