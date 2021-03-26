package middleware

import (
	"bytes"
	"fmt"
	"github.com/Sebalvarez97/mutants-go/errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCustomErrorHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("TestHandleValidationError", func(t *testing.T) {
		router := gin.Default()
		errorHandler := GetCustomErrorHandler()
		router.Use(errorHandler)

		router.Handle(http.MethodGet, "/test", func(context *gin.Context) {
			context.Error(errors.NewValidationError("error validating something"))
		})

		request, err := http.NewRequest(http.MethodGet, "/test", nil)
		assert.NoError(t, err)

		expected := bytes.NewBuffer([]byte("{\n    \"StatusCode\": 400,\n    \"Error\": {\n        \"message\": \"error validating something\",\n        \"error\": \"bad_request\",\n        \"status\": 400,\n        \"cause\": null\n    },\n    \"Notify\": false\n}"))

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, request)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})
	t.Run("TestHandleCommunicationError", func(t *testing.T) {
		router := gin.Default()
		errorHandler := GetCustomErrorHandler()
		router.Use(errorHandler)

		router.Handle(http.MethodGet, "/test", func(context *gin.Context) {
			context.Error(errors.NewCommunicationError("error communicating something", context.Request.RequestURI, 412))
		})

		request, err := http.NewRequest(http.MethodGet, "/test", nil)
		assert.NoError(t, err)

		expected := bytes.NewBuffer([]byte("{\n    \"StatusCode\": 500,\n    \"Error\": {\n        \"message\": \"error communicating something\",\n        \"error\": \"internal_server_error\",\n        \"status\": 500,\n        \"cause\": null\n    },\n    \"Notify\": false\n}"))

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, request)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})
	t.Run("TestHandleNotFoundError", func(t *testing.T) {
		router := gin.Default()
		errorHandler := GetCustomErrorHandler()
		router.Use(errorHandler)

		router.Handle(http.MethodGet, "/test/:id", func(context *gin.Context) {
			id := context.Param("id")
			context.Error(errors.NewNotFoundError("test_resource", id))
		})

		request, err := http.NewRequest(http.MethodGet, "/test/1234", nil)
		assert.NoError(t, err)

		expected := bytes.NewBuffer([]byte("{\n    \"StatusCode\": 404,\n    \"Error\": {\n        \"message\": \"The test_resource with identifier 1234 was not found\",\n        \"error\": \"not_found\",\n        \"status\": 404,\n        \"cause\": null\n    },\n    \"Notify\": false\n}"))

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, request)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})
	t.Run("TestHandleNotFoundWithoutIdError", func(t *testing.T) {
		router := gin.Default()
		errorHandler := GetCustomErrorHandler()
		router.Use(errorHandler)

		router.Handle(http.MethodGet, "/test", func(context *gin.Context) {
			context.Error(errors.NewNotFoundErrorWithoutId("test_resource"))
		})

		request, err := http.NewRequest(http.MethodGet, "/test", nil)
		assert.NoError(t, err)

		expected := bytes.NewBuffer([]byte("{\n    \"StatusCode\": 404,\n    \"Error\": {\n        \"message\": \"The test_resource was not found\",\n        \"error\": \"not_found\",\n        \"status\": 404,\n        \"cause\": null\n    },\n    \"Notify\": false\n}"))

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, request)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})
	t.Run("TestHandleAuthError", func(t *testing.T) {
		router := gin.Default()
		errorHandler := GetCustomErrorHandler()
		router.Use(errorHandler)

		router.Handle(http.MethodGet, "/test", func(context *gin.Context) {
			context.Error(errors.NewAuthorizationError("fail to authorize", "auth_fail"))
		})

		request, err := http.NewRequest(http.MethodGet, "/test", nil)
		assert.NoError(t, err)

		expected := bytes.NewBuffer([]byte("{\n    \"StatusCode\": 401,\n    \"Error\": {\n        \"message\": \"fail to authorize\",\n        \"error\": \"unauthorized\",\n        \"status\": 401,\n        \"cause\": null\n    },\n    \"Notify\": false\n}"))

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, request)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})
	t.Run("TestHandleOtherError", func(t *testing.T) {
		router := gin.Default()
		errorHandler := GetCustomErrorHandler()
		router.Use(errorHandler)

		router.Handle(http.MethodGet, "/test", func(context *gin.Context) {
			context.Error(fmt.Errorf("generic error"))
		})

		request, err := http.NewRequest(http.MethodGet, "/test", nil)
		assert.NoError(t, err)

		expected := bytes.NewBuffer([]byte("{\n    \"StatusCode\": 500,\n    \"Error\": {\n        \"message\": \"generic error\",\n        \"error\": \"internal_server_error\",\n        \"status\": 500,\n        \"cause\": null\n    },\n    \"Notify\": false\n}"))

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, request)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})
}
