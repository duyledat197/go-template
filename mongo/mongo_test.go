package mongo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestNewMongoDbCtxCollection(t *testing.T) {
	testMongoColl := &mongo.Collection{}
	testMongoCtxColl := NewMongoDbCtxCollection(testMongoColl)
	mongoCtxColl := testMongoCtxColl.(*mongoDbCtxCollection)
	assert.NotNil(t, testMongoCtxColl)
	assert.Equal(t, testMongoColl, mongoCtxColl.collection)
}
