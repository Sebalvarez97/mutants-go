package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Sebalvarez97/mutants/api/errors"
	. "github.com/Sebalvarez97/mutants/api/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MutantServiceImplMock struct{}

var mutantServiceIsMutantMock func(body IsMutantRequestBody) bool

func (i MutantServiceImplMock) IsMutant(body IsMutantRequestBody) bool {
	return mutantServiceIsMutantMock(body)
}

var mutantServiceGetMutantStats func() (*Stats, *errors.ApiErrorImpl)

func (i MutantServiceImplMock) GetMutantStats() (*Stats, *errors.ApiErrorImpl) {
	return mutantServiceGetMutantStats()
}

func TestGetIsMutant(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("Success Mutant", func(t *testing.T) {

		input := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
		values := map[string][]string{"dna": input}
		jsonValue, _ := json.Marshal(values)

		service := MutantServiceImplMock{}
		mutantServiceIsMutantMock = func(body IsMutantRequestBody) bool {
			return true
		}

		controller := NewMutantController(service)

		rr := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/mutant", controller.IsMutantHandler)

		request, err := http.NewRequest(http.MethodPost, "/mutant", bytes.NewReader(jsonValue))
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		assert.NoError(t, err)

		assert.Equal(t, 200, rr.Code)
	})

	t.Run("Success No Mutant", func(t *testing.T) {

		input := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
		values := map[string][]string{"dna": input}
		jsonValue, _ := json.Marshal(values)

		service := MutantServiceImplMock{}
		mutantServiceIsMutantMock = func(body IsMutantRequestBody) bool {
			return false
		}

		controller := NewMutantController(service)

		rr := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/mutant", controller.IsMutantHandler)

		request, err := http.NewRequest(http.MethodPost, "/mutant", bytes.NewReader(jsonValue))
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		assert.NoError(t, err)

		assert.Equal(t, 403, rr.Code)
	})
}

func TestGetIsMutantFail(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("Bad Request Binding Error", func(t *testing.T) {

		input := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
		values := map[string][]string{"dna_chain": input}
		jsonValue, _ := json.Marshal(values)

		service := MutantServiceImplMock{}
		mutantServiceIsMutantMock = func(body IsMutantRequestBody) bool {
			return true
		}

		controller := NewMutantController(service)

		rr := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/mutant", controller.IsMutantHandler)

		request, err := http.NewRequest(http.MethodPost, "/mutant", bytes.NewReader(jsonValue))
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		assert.NoError(t, err)

		assert.Equal(t, 400, rr.Code)
	})
}

func TestIsMutantBadRequest(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("BadRequest Smaller Matrix", func(t *testing.T) {

		input := []string{"AGA", "CAG", "TTA"}
		values := map[string][]string{"dna": input}
		jsonValue, _ := json.Marshal(values)

		service := MutantServiceImplMock{}

		controller := NewMutantController(service)

		rr := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/mutant", controller.IsMutantHandler)

		request, err := http.NewRequest(http.MethodPost, "/mutant", bytes.NewReader(jsonValue))
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		assert.NoError(t, err)

		expected := bytes.NewBuffer([]byte("{\"Code\":400,\"Message\":\"Invalid value entered: invalid input, the matrix is to short, has to be 4x4 or bigger\",\"Cause\":{}}"))
		assert.Equal(t, 400, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})

	t.Run("BadRequest Not NxN Matrix", func(t *testing.T) {

		input := []string{"AGA", "CAG", "TTA", "TTT"}
		values := map[string][]string{"dna": input}
		jsonValue, _ := json.Marshal(values)

		service := MutantServiceImplMock{}

		controller := NewMutantController(service)

		rr := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/mutant", controller.IsMutantHandler)

		request, err := http.NewRequest(http.MethodPost, "/mutant", bytes.NewReader(jsonValue))
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		assert.NoError(t, err)

		expected := bytes.NewBuffer([]byte("{\"Code\":400,\"Message\":\"Invalid value entered: invalid input, it isn't a NxN matrix, this could cause an Internal Error\",\"Cause\":{}}"))
		assert.Equal(t, 400, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})

	t.Run("BadRequest Nitrogen Base", func(t *testing.T) {

		input := []string{"AGAA", "CAGA", "TTAZ", "TTTA"}
		values := map[string][]string{"dna": input}
		jsonValue, _ := json.Marshal(values)

		service := MutantServiceImplMock{}

		controller := NewMutantController(service)

		rr := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/mutant", controller.IsMutantHandler)

		request, err := http.NewRequest(http.MethodPost, "/mutant", bytes.NewReader(jsonValue))
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		assert.NoError(t, err)

		expected := bytes.NewBuffer([]byte("{\"Code\":400,\"Message\":\"Invalid value entered: invalid nitrogen base found: Z\",\"Cause\":{}}"))
		assert.Equal(t, 400, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})
}

func TestGetStatsOk(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success Stats More Humans", func(t *testing.T) {

		stats := NewStats(40, 100, 0.4)

		service := MutantServiceImplMock{}
		mutantServiceGetMutantStats = func() (*Stats, *errors.ApiErrorImpl) {
			return stats, nil
		}

		controller := NewMutantController(service)

		rr := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/mutant/stats", controller.GetStatsHandler)

		request, err := http.NewRequest(http.MethodGet, "/mutant/stats", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		assert.NoError(t, err)

		expected := bytes.NewBuffer([]byte("{\"count_mutant_dna\":40,\"count_human_dna\":100,\"ratio\":0.4}"))
		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})

	t.Run("Success Stats More Mutants", func(t *testing.T) {

		stats := NewStats(100, 25, 4.0)

		service := MutantServiceImplMock{}
		mutantServiceGetMutantStats = func() (*Stats, *errors.ApiErrorImpl) {
			return stats, nil
		}

		controller := NewMutantController(service)

		rr := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/mutant/stats", controller.GetStatsHandler)

		request, err := http.NewRequest(http.MethodGet, "/mutant/stats", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		assert.NoError(t, err)

		expected := bytes.NewBuffer([]byte("{\"count_mutant_dna\":100,\"count_human_dna\":25,\"ratio\":4}"))
		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})

	t.Run("Success Stats Cero Humans", func(t *testing.T) {

		stats := NewStats(100, 0, 1.0)

		service := MutantServiceImplMock{}
		mutantServiceGetMutantStats = func() (*Stats, *errors.ApiErrorImpl) {
			return stats, nil
		}

		controller := NewMutantController(service)

		rr := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/mutant/stats", controller.GetStatsHandler)

		request, err := http.NewRequest(http.MethodGet, "/mutant/stats", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		assert.NoError(t, err)

		expected := bytes.NewBuffer([]byte("{\"count_mutant_dna\":100,\"count_human_dna\":0,\"ratio\":1}"))
		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})
}

func TestGetStatsFailure(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Internal Server Error db", func(t *testing.T) {

		service := MutantServiceImplMock{}
		mutantServiceGetMutantStats = func() (*Stats, *errors.ApiErrorImpl) {
			apiErr := errors.GenericError(fmt.Errorf("db error"))
			return &Stats{}, &apiErr
		}

		controller := NewMutantController(service)

		rr := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/mutant/stats", controller.GetStatsHandler)

		request, err := http.NewRequest(http.MethodGet, "/mutant/stats", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		assert.NoError(t, err)

		assert.Equal(t, 500, rr.Code)
	})
}
