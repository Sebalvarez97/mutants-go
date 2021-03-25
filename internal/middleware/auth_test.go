package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type AuthenticationResponse struct {
	Code   int    `json:"code"`
	Expire string `json:"expire"`
	Token  string `json:"token"`
}

func TestGetAuthMiddleWare(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("TestAuthorized", func(t *testing.T) {
		rr := httptest.NewRecorder()
		router := gin.Default()
		authMiddleWare := GetAuthMiddleWare(router)
		router.Use(authMiddleWare)

		body := []byte("{\"username\":\"admin\",\"password\":\"admin\"}")
		authRequest, authErr := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(body))
		ar := httptest.NewRecorder()
		assert.NoError(t, authErr)
		authRequest.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(ar, authRequest)
		assert.NoError(t, authErr)
		assert.Equal(t, 200, ar.Code)

		b, readErr := ioutil.ReadAll(ar.Body)
		assert.NoError(t, readErr)

		var authBody AuthenticationResponse
		unMarshalErr := json.Unmarshal(b, &authBody)
		assert.NoError(t, unMarshalErr)
		assert.Equal(t, 200, authBody.Code)

		token := authBody.Token

		router.Handle(http.MethodGet, "/test", func(context *gin.Context) {
			context.Status(http.StatusOK)
		})

		request, err := http.NewRequest(http.MethodGet, "/test", nil)
		assert.NoError(t, err)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+token)

		expected := bytes.NewBuffer([]byte(nil))

		router.ServeHTTP(rr, request)
		assert.NoError(t, err)
		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, expected, rr.Body)
	})
}
