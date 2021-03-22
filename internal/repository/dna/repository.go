package dna

import (
	"context"
	"github.com/Sebalvarez97/mutants-go/errors"
	"github.com/Sebalvarez97/mutants-go/internal/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const (
	collection  = "dna"
	notFoundErr = "mongo: no documents in result"
)

type DnaRepository interface {
	FindByDnaHash(ctx context.Context, hash string) (*model.Dna, error)
	Upsert(ctx context.Context, dna *model.Dna) error
	insert(ctx context.Context, dna *model.Dna) error
	update(ctx context.Context, hash string, dna *model.Dna) error
	FindNumberOfHumans(ctx context.Context) (int, error)
	FindNumberOfMutants(ctx context.Context) (int, error)
}

type MongoDao interface {
	FindAll(ctx context.Context, filter bson.D, collectionName string) (*mongo.Cursor, error)
	CountForCollection(ctx context.Context, filter bson.D, collectionName string) (int64, error)
	FindMany(ctx context.Context, limit int64, filter bson.D, collectionName string) (*mongo.Cursor, error)
	FindOne(ctx context.Context, filter bson.D, collectionName string) (*mongo.SingleResult, error)
	InsertOne(ctx context.Context, document interface{}, collectionName string) error
	UpdateOne(ctx context.Context, filter bson.D, update bson.D, collectionName string) error
}

func NewDnaRepository(dao MongoDao) DnaRepository {
	return &repository{dao: dao}
}

type repository struct {
	dao MongoDao
}

func (r *repository) FindByDnaHash(ctx context.Context, hash string) (*model.Dna, error) {
	dna, err := r.dao.FindOne(ctx, bson.D{{"dna_hash", hash}}, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		if err.Error() == notFoundErr {
			apiErr = errors.NotFoundError(err)
		}
		log.Print(apiErr.Error())
		return &model.Dna{}, &apiErr
	}
	return mapToObject(ctx, dna)
}

func (r *repository) Upsert(ctx context.Context, dna *model.Dna) error {
	err := r.update(ctx, dna.DnaHash, dna)
	if err != nil {
		if err.(errors.ApiError).Code == errors.NotFoundCode {
			err = r.insert(ctx, dna)
			log.Print(err.Error())
		}
	}
	return nil
}

func (r *repository) insert(ctx context.Context, dna *model.Dna) error {
	if err := r.dao.InsertOne(ctx, dna, collection); err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return &apiErr
	}
	return nil
}

func (r *repository) update(ctx context.Context, hash string, dna *model.Dna) error {
	_, findErr := r.FindByDnaHash(ctx, hash)
	if findErr != nil {
		return findErr
	}
	updateErr := r.dao.UpdateOne(ctx, bson.D{{"dna_hash", hash}}, bson.D{{"$set", bson.D{{"is_mutant", dna.IsMutant}, {"mutant_sequences", dna.MutantSequences}}}}, collection)
	if updateErr != nil {
		apiErr := errors.GenericError(updateErr)
		if updateErr.Error() == notFoundErr {
			apiErr = errors.NotFoundError(updateErr)
		}
		log.Print(apiErr.Error())
		return &apiErr
	}
	return nil
}

func (r *repository) FindNumberOfHumans(ctx context.Context) (int, error) {
	filter := bson.D{{"is_mutant", false}}
	count, err := r.dao.CountForCollection(ctx, filter, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return 0, &apiErr
	}
	return int(count), nil
}

func (r *repository) FindNumberOfMutants(ctx context.Context) (int, error) {
	filter := bson.D{{"is_mutant", true}}
	count, err := r.dao.CountForCollection(ctx, filter, collection)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return 0, &apiErr
	}
	return int(count), nil
}

func mapToObjects(ctx context.Context, cur *mongo.Cursor) ([]model.Dna, error) {
	var results []model.Dna
	err := cur.All(context.TODO(), &results)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return results, &apiErr
	}
	return results, nil
}

func mapToObject(ctx context.Context, document *mongo.SingleResult) (*model.Dna, error) {
	var dna model.Dna
	err := document.Decode(&dna)
	if err != nil {
		apiErr := errors.GenericError(err)
		log.Print(apiErr.Error())
		return &dna, &apiErr
	}
	return &dna, nil
}
