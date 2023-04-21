package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	connectionString = "mongodb://zeroPass:9674Ephzx&T7@127.0.0.1:27017/?retryWrites=true&w=majority"
)

func TestNewMongoClient(t *testing.T) {
	// Act
	mc, err := NewMongoClient(connectionString)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, mc)
}

func TestInsertDocument(t *testing.T) {
	// Arrange
	mc, _ := NewMongoClient(connectionString)

	// Act
	result, err := mc.InsertDocument("test", "test", map[string]string{
		"test": "test",
	})

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
