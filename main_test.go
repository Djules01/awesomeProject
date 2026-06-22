package main

import (
	mongodbadapter "awesomeProject/adapters/mongodb"
	"awesomeProject/configuration"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunReturnsDBError(t *testing.T) {
	config := configuration.Config{
		APIKey:   "test-key",
		MongoURI: "://uri-invalide",
		MongoDB:  "test",
		Port:     "8080",
	}

	err := run(config, func(addr string, handler http.Handler) error {
		t.Fatal("listenAndServe ne doit pas etre appele si InitClient retourne une erreur")
		return nil
	})

	require.Error(t, err)
}

func TestRunReturnsListenAndServeError(t *testing.T) {
	if _, err := mongodbadapter.InitClient("mongodb://localhost:27017"); err != nil {
		t.Skip("MongoDB requis sur localhost:27017 pour ce test")
	}

	expectedErr := errors.New("server error")
	config := configuration.Config{
		APIKey:   "test-key",
		MongoURI: "mongodb://localhost:27017",
		MongoDB:  "test_todolist",
		Port:     "9090",
	}

	err := run(config, func(addr string, handler http.Handler) error {
		assert.Equal(t, ":9090", addr)
		require.NotNil(t, handler)
		return expectedErr
	})

	assert.ErrorIs(t, err, expectedErr)
}
