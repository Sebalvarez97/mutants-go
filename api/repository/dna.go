package repository

import (
	"context"
	"github.com/Sebalvarez97/mutants/api/errors"
	"github.com/Sebalvarez97/mutants/api/interfaces"
	"github.com/Sebalvarez97/mutants/api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type DnaRepositoryImpl struct {
	dao interfaces.MongoDao
}

func NewDnaRepository(dao interfaces.MongoDao) interfaces.DnaRepository {
	return DnaRepositoryImpl{dao: dao}
}

const collection = "dna"
const MongoNotFoundErr = "mongo: no documents in result"

func (i DnaRepositoryImpl) FindAll() ([]model.Dna, *errors.ApiErrorImpl) {
	dna, err := i.dao.FindAll(bson.D{}, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return nil, &apiErr
	}
	return mapToObjects(dna)
}

func (i DnaRepositoryImpl) FindMany() ([]model.Dna, *errors.ApiErrorImpl) {
	dna, err := i.dao.FindMany(10, bson.D{}, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return nil, &apiErr
	}
	return mapToObjects(dna)
}

func (i DnaRepositoryImpl) FindAllMutants() ([]model.Dna, *errors.ApiErrorImpl) {
	filter := bson.D{{"is_mutant", true}}
	dna, err := i.dao.FindAll(filter, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return nil, &apiErr
	}
	return mapToObjects(dna)
}

func (i DnaRepositoryImpl) FindNumberOfMutants() (int, *errors.ApiErrorImpl) {
	filter := bson.D{{"is_mutant", true}}
	count, err := i.dao.CountForCollection(filter, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return 0, &apiErr
	}
	return int(count), nil
}

func (i DnaRepositoryImpl) FindById(id string) (model.Dna, *errors.ApiErrorImpl) {
	idObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Print(err)
		apiErr := errors.GenericError(err)
		return model.Dna{}, &apiErr
	}
	dna, err := i.dao.FindOne(bson.D{{"_id", idObject}}, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		if err.Error() == MongoNotFoundErr {
			apiErr = errors.NotFoundError(err)
		}
		log.Print(apiErr.Error())
		return model.Dna{}, &apiErr
	}
	return mapToObject(dna)
}

func (i DnaRepositoryImpl) FindAllHumans() ([]model.Dna, *errors.ApiErrorImpl) {
	filter := bson.D{{"is_mutant", false}}
	dna, err := i.dao.FindAll(filter, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return nil, &apiErr
	}
	return mapToObjects(dna)
}

func (i DnaRepositoryImpl) FindNumberOfHumans() (int, *errors.ApiErrorImpl) {
	filter := bson.D{{"is_mutant", false}}
	count, err := i.dao.CountForCollection(filter, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return 0, &apiErr
	}
	return int(count), nil
}

func (i DnaRepositoryImpl) FindByDnaHash(hash string) (model.Dna, *errors.ApiErrorImpl) {
	dna, err := i.dao.FindOne(bson.D{{"dna_hash", hash}}, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		if err.Error() == MongoNotFoundErr {
			apiErr = errors.NotFoundError(err)
		}
		log.Print(apiErr.Error())
		return model.Dna{}, &apiErr
	}
	return mapToObject(dna)
}

func (i DnaRepositoryImpl) Insert(dna *model.Dna) *errors.ApiErrorImpl {
	if err := i.dao.InsertOne(dna, collection); err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return &apiErr
	}
	return nil
}

func (i DnaRepositoryImpl) Upsert(dna *model.Dna) {
	err := i.Update(dna.DnaHash, dna)
	if err != nil {
		if err.Code == errors.NotFoundCode {
			err = i.Insert(dna)
			log.Print(err.Error())
		}
	}
}

func (i DnaRepositoryImpl) Update(hash string, dna *model.Dna) *errors.ApiErrorImpl {
	_, findErr := i.FindByDnaHash(hash)
	if findErr != nil {
		return findErr
	}
	updateErr := i.dao.UpdateOne(bson.D{{"dna_hash", hash}}, bson.D{{"$set", bson.D{{"is_mutant", dna.IsMutant}, {"mutant_sequences", dna.MutantSequences}}}}, collection)
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

func mapToObjects(cur *mongo.Cursor) ([]model.Dna, *errors.ApiErrorImpl) {
	var results []model.Dna
	err := cur.All(context.TODO(), &results)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return results, &apiErr
	}
	return results, nil
}

func mapToObject(document *mongo.SingleResult) (model.Dna, *errors.ApiErrorImpl) {
	var dna model.Dna
	err := document.Decode(&dna)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return dna, &apiErr
	}
	return dna, nil
}
