package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_WhenTokenNotInRequestHeader_ThenReturnErrRequestContainsInvalidToken(t *testing.T) {
	handlerExecuted := false
	r := gin.Default()
	r.Use(Auth("secret"))
	{
		r.GET("/ping", func(c *gin.Context) {
			handlerExecuted = true
		})
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert.False(t, handlerExecuted)
	assert.Equal(t, 401, w.Code)
	assert.Equal(t, "{\"error\":\"error invalid token\",\"success\":false}", w.Body.String())
}

func Test_WhenTokenIsInvalid_ThenReturnErrRequestContainsInvalidToken(t *testing.T) {
	handlerExecuted := false
	r := gin.Default()
	r.Use(Auth("secret"))
	{
		r.GET("/ping", func(c *gin.Context) {
			handlerExecuted = true
		})
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	req.Header.Add("Authorization", "invalid_token")
	r.ServeHTTP(w, req)

	assert.False(t, handlerExecuted)
	assert.Equal(t, 401, w.Code)
	assert.Equal(t, "{\"error\":\"error invalid token\",\"success\":false}", w.Body.String())
}

func Test_WhenTokenIsValid_ThenExecuteHandler(t *testing.T) {
	handlerExecuted := false
	r := gin.Default()
	r.Use(Auth("secret"))
	{
		r.GET("/ping", func(c *gin.Context) {
			handlerExecuted = true
		})
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	token, _ := GenerateToken("email@email.com", "username", "secret")
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)

	assert.True(t, handlerExecuted)
}
