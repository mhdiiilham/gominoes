package mongodb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSuccessMongoConnection(t *testing.T) {
	client, err := NewMongoDBConnection(os.Getenv("MONGO_DB_USER"), os.Getenv("MONGO_DB_PASS"), os.Getenv("MONGO_DB"))

	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, client, "Client should not be nill")
}
