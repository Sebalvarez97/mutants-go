package writer

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMutantWriterHandler_IsMutantHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("IsMutantTrue", func(t *testing.T) {
		input := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
		values := map[string][]string{"dna": input}
		jsonValue, _ := json.Marshal(values)

		writerHandler := newFakeMutantWriterHandler()
		routerHandler := NewRouterHandler(writerHandler.MutantWriterHandler)

		writerHandler.MutantSrv.Mock.On("IsMutant", mock.Anything).Return(true, nil)

		rr := httptest.NewRecorder()
		router := gin.Default()
		routerHandler.RouteURLs(router)

		request, err := http.NewRequest(http.MethodPost, "/mutant", bytes.NewReader(jsonValue))
		router.ServeHTTP(rr, request)
		assert.NoError(t, err)
		assert.Equal(t, 200, rr.Code)
	})
	t.Run("IsMutantFalse", func(t *testing.T) {
		input := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
		values := map[string][]string{"dna": input}
		jsonValue, _ := json.Marshal(values)

		writerHandler := newFakeMutantWriterHandler()
		routerHandler := NewRouterHandler(writerHandler.MutantWriterHandler)

		writerHandler.MutantSrv.Mock.On("IsMutant", mock.Anything).Return(false, nil)

		rr := httptest.NewRecorder()
		router := gin.Default()
		routerHandler.RouteURLs(router)

		request, err := http.NewRequest(http.MethodPost, "/mutant", bytes.NewReader(jsonValue))
		assert.NoError(t, err)
		router.ServeHTTP(rr, request)
		assert.NoError(t, err)
		assert.Equal(t, 403, rr.Code)
	})
	t.Run("IsMutantBindingError", func(t *testing.T) {
		input := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
		values := map[string][]string{"dna_chain": input}
		jsonValue, _ := json.Marshal(values)

		writerHandler := newFakeMutantWriterHandler()
		routerHandler := NewRouterHandler(writerHandler.MutantWriterHandler)

		rr := httptest.NewRecorder()
		router := gin.Default()
		routerHandler.RouteURLs(router)

		request, err := http.NewRequest(http.MethodPost, "/mutant", bytes.NewReader(jsonValue))
		expected := bytes.NewBuffer([]byte(nil))
		assert.NoError(t, err)
		router.ServeHTTP(rr, request)
		assert.NoError(t, err)
		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})
	t.Run("IsMutantInvalidErrorToSmallMatrix", func(t *testing.T) {
		input := []string{"AGA", "CAG", "TTA"}
		values := map[string][]string{"dna": input}
		jsonValue, _ := json.Marshal(values)

		writerHandler := newFakeMutantWriterHandler()
		routerHandler := NewRouterHandler(writerHandler.MutantWriterHandler)

		rr := httptest.NewRecorder()
		router := gin.Default()
		routerHandler.RouteURLs(router)

		request, err := http.NewRequest(http.MethodPost, "/mutant", bytes.NewReader(jsonValue))
		expected := bytes.NewBuffer([]byte(nil))
		assert.NoError(t, err)
		router.ServeHTTP(rr, request)
		assert.NoError(t, err)
		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})
	t.Run("IsMutantInvalidErrorNotNxNMatrix", func(t *testing.T) {
		input := []string{"AGA", "CAG", "TTA", "TTT"}
		values := map[string][]string{"dna": input}
		jsonValue, _ := json.Marshal(values)

		writerHandler := newFakeMutantWriterHandler()
		routerHandler := NewRouterHandler(writerHandler.MutantWriterHandler)

		rr := httptest.NewRecorder()
		router := gin.Default()
		routerHandler.RouteURLs(router)

		request, err := http.NewRequest(http.MethodPost, "/mutant", bytes.NewReader(jsonValue))
		expected := bytes.NewBuffer([]byte(nil))
		assert.NoError(t, err)
		router.ServeHTTP(rr, request)
		assert.NoError(t, err)
		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})
	t.Run("IsMutantInvalidErrorInvalidDnaNitrogenBase", func(t *testing.T) {
		input := []string{"AGAA", "CAGA", "TTAZ", "TTTA"}
		values := map[string][]string{"dna": input}
		jsonValue, _ := json.Marshal(values)

		writerHandler := newFakeMutantWriterHandler()
		routerHandler := NewRouterHandler(writerHandler.MutantWriterHandler)

		rr := httptest.NewRecorder()
		router := gin.Default()
		routerHandler.RouteURLs(router)

		request, err := http.NewRequest(http.MethodPost, "/mutant", bytes.NewReader(jsonValue))
		expected := bytes.NewBuffer([]byte(nil))
		assert.NoError(t, err)
		router.ServeHTTP(rr, request)
		assert.NoError(t, err)
		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})
}
