package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHelloWorld(t *testing.T) {
	// setup http routes
	r := SetUpRouter()
	r.GET("/", HelloWorld)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// testing
	expected := `{"Hello":"World"}`
	actual, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, expected, string(actual))
	assert.Equal(t, http.StatusOK, w.Code)
}

