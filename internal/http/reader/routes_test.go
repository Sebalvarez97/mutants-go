package reader

import (
	"bytes"
	"fmt"
	"github.com/Sebalvarez97/mutants-go/errors"
	"github.com/Sebalvarez97/mutants-go/internal/domain/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMutantReaderHandler_GetStatsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetStatsMoreHumans", func(t *testing.T) {
		readerHandler := newFakeMutantReaderHandler()
		routerHandler := NewRouterHandler(readerHandler.MutantReaderHandler)

		stats := model.NewStats(40, 100, 0.4)
		readerHandler.MutantSrv.Mock.On("GetMutantStats", mock.Anything).Return(stats, nil)

		rr := httptest.NewRecorder()
		router := gin.Default()
		routerHandler.RouteURLs(router)

		request, err := http.NewRequest(http.MethodGet, "/mutant/stats", nil)
		expected := bytes.NewBuffer([]byte("{\"count_mutant_dna\":40,\"count_human_dna\":100,\"ratio\":0.4}"))
		assert.NoError(t, err)
		router.ServeHTTP(rr, request)
		assert.NoError(t, err)
		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})
	t.Run("GetStatsMoreMutants", func(t *testing.T) {
		readerHandler := newFakeMutantReaderHandler()
		routerHandler := NewRouterHandler(readerHandler.MutantReaderHandler)

		stats := model.NewStats(100, 40, 2.5)
		readerHandler.MutantSrv.Mock.On("GetMutantStats", mock.Anything).Return(stats, nil)

		rr := httptest.NewRecorder()
		router := gin.Default()
		routerHandler.RouteURLs(router)

		request, err := http.NewRequest(http.MethodGet, "/mutant/stats", nil)
		expected := bytes.NewBuffer([]byte("{\"count_mutant_dna\":100,\"count_human_dna\":40,\"ratio\":2.5}"))
		assert.NoError(t, err)
		router.ServeHTTP(rr, request)
		assert.NoError(t, err)
		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})
	t.Run("GetStatsZeroHumans", func(t *testing.T) {
		readerHandler := newFakeMutantReaderHandler()
		routerHandler := NewRouterHandler(readerHandler.MutantReaderHandler)

		stats := model.NewStats(100, 0, 1.0)
		readerHandler.MutantSrv.Mock.On("GetMutantStats", mock.Anything).Return(stats, nil)

		rr := httptest.NewRecorder()
		router := gin.Default()
		routerHandler.RouteURLs(router)

		request, err := http.NewRequest(http.MethodGet, "/mutant/stats", nil)
		expected := bytes.NewBuffer([]byte("{\"count_mutant_dna\":100,\"count_human_dna\":0,\"ratio\":1}"))
		assert.NoError(t, err)
		router.ServeHTTP(rr, request)
		assert.NoError(t, err)
		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})
	t.Run("GetStatsDbError", func(t *testing.T) {
		readerHandler := newFakeMutantReaderHandler()
		routerHandler := NewRouterHandler(readerHandler.MutantReaderHandler)

		readerHandler.MutantSrv.Mock.On("GetMutantStats", mock.Anything).Return(nil, errors.DbConnectionError(fmt.Errorf("db connection error")))

		rr := httptest.NewRecorder()
		router := gin.Default()
		routerHandler.RouteURLs(router)

		request, err := http.NewRequest(http.MethodGet, "/mutant/stats", nil)
		expected := bytes.NewBuffer([]byte("{\"Code\":500,\"Message\":\"Server failed to connect/disconnect from db because of db connection error\",\"Cause\":{}}"))
		assert.NoError(t, err)
		router.ServeHTTP(rr, request)
		assert.NoError(t, err)
		assert.Equal(t, 500, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})
}
