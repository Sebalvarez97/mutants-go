package mutant

import (
	"context"
	"fmt"
	"github.com/Sebalvarez97/mutants-go/internal/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func initTest() (context.Context, *fakeMutantContainer) {
	ctx := context.Background()
	fc := newFakeContainer()
	return ctx, fc
}

func TestService_GetMutantStats(t *testing.T) {
	t.Run("GetMutantStatsMoreHumansThanMutants", func(t *testing.T) {
		ctx, fc := initTest()
		s := NewService(fc.Container)

		fc.DnaRepository.Mock.On("FindNumberOfMutants", mock.Anything).Return(40, nil)
		fc.DnaRepository.Mock.On("FindNumberOfHumans", mock.Anything).Return(100, nil)

		stats, err := s.GetMutantStats(ctx)
		assert.Nil(t, err)
		assert.NotNil(t, stats)
		assert.Equal(t, 40, stats.Mutants)
		assert.Equal(t, 100, stats.Humans)
		assert.Equal(t, 0.4, stats.Ratio)
	})
	t.Run("GetMutantStatsMoreMutantsThanHumans", func(t *testing.T) {
		ctx, fc := initTest()
		s := NewService(fc.Container)

		fc.DnaRepository.Mock.On("FindNumberOfMutants", mock.Anything).Return(100, nil)
		fc.DnaRepository.Mock.On("FindNumberOfHumans", mock.Anything).Return(40, nil)

		stats, err := s.GetMutantStats(ctx)
		assert.Nil(t, err)
		assert.NotNil(t, stats)
		assert.Equal(t, 100, stats.Mutants)
		assert.Equal(t, 40, stats.Humans)
		assert.Equal(t, 2.5, stats.Ratio)
	})
	t.Run("GetMutantStatsWithZeroHumans", func(t *testing.T) {
		ctx, fc := initTest()
		s := NewService(fc.Container)

		fc.DnaRepository.Mock.On("FindNumberOfMutants", mock.Anything).Return(100, nil)
		fc.DnaRepository.Mock.On("FindNumberOfHumans", mock.Anything).Return(0, nil)

		stats, err := s.GetMutantStats(ctx)
		assert.Nil(t, err)
		assert.NotNil(t, stats)
		assert.Equal(t, 100, stats.Mutants)
		assert.Equal(t, 0, stats.Humans)
		assert.Equal(t, 1.0, stats.Ratio)
	})
	t.Run("GetMutantStatsWithSomeError", func(t *testing.T) {
		ctx, fc := initTest()
		s := NewService(fc.Container)

		fc.DnaRepository.Mock.On("FindNumberOfMutants", mock.Anything).Return(100, fmt.Errorf("database error getting mutants"))
		fc.DnaRepository.Mock.On("FindNumberOfHumans", mock.Anything).Return(0, nil)

		stats, err := s.GetMutantStats(ctx)
		assert.NotNil(t, err)
		assert.NotNil(t, stats)
		assert.Equal(t, 0, stats.Mutants)
		assert.Equal(t, 0, stats.Humans)
		assert.Equal(t, 0.0, stats.Ratio)
	})
}

func TestService_IsMutant(t *testing.T) {
	t.Run("IsMutantOk", func(t *testing.T) {
		ctx, fc := initTest()
		s := NewService(fc.Container)

		fc.DnaRepository.Mock.On("Upsert", mock.Anything).Return(nil)

		sa := []string{
			"ATGCGA",
			"CAGAGC",
			"TTATGT",
			"AGAAGG",
			"CCCCTA",
			"TCACTG"}

		request := model.IsMutantRequestBody{Dna: sa}

		mut, err := s.IsMutant(ctx, request)

		assert.Nil(t, err)
		assert.NotNil(t, mut)
		assert.True(t, mut)
	})
	t.Run("IsNotMutantOk", func(t *testing.T) {
		ctx, fc := initTest()
		s := NewService(fc.Container)

		fc.DnaRepository.Mock.On("Upsert", mock.Anything).Return(nil)

		sa := []string{
			"ATGCGA",
			"CAGTGC",
			"TTAGTG",
			"AGAGGG",
			"CCCTTA",
			"TCACTG"}

		request := model.IsMutantRequestBody{Dna: sa}

		mut, err := s.IsMutant(ctx, request)

		assert.Nil(t, err)
		assert.NotNil(t, mut)
		assert.False(t, mut)
	})
}
