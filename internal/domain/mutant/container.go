package mutant

import (
	"context"
	"github.com/Sebalvarez97/mutants-go/config"
	"github.com/Sebalvarez97/mutants-go/db/mongo"
	"github.com/Sebalvarez97/mutants-go/internal/domain/model"
	"github.com/Sebalvarez97/mutants-go/internal/repository/dna"
)

type Container struct {
	DnaRepository DnaRepository
}

type DnaRepository interface {
	FindByDnaHash(ctx context.Context, hash string) (*model.Dna, error)
	Upsert(ctx context.Context, dna *model.Dna) error
	FindNumberOfHumans(ctx context.Context) (int, error)
	FindNumberOfMutants(ctx context.Context) (int, error)
}

func InitializeContainer(config config.Config) Container {
	return Container{dna.NewDnaRepository(mongo.NewMongoDao(config.Mongo))}
}
