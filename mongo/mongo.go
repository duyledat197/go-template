package mongo

import (
	"context"
	"errors"
	"log"
	"time"

	mongoI "github.com/duyledat197/go-template/mongo/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	timeout             = 10
	errNilClientMsg     = "Client must not be nil"
	errEmptyDbNameMsg   = "Database name must not be empty"
	errEmptyCollNameMsg = "Collection name must not be empty"
)

var (
	errNilClient     = errors.New(errNilClientMsg)
	errEmptyDbName   = errors.New(errEmptyDbNameMsg)
	errEmptyCollName = errors.New(errEmptyCollNameMsg)
)

func getContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), timeout*time.Second)
	return ctx
}

type mongoDbCtxCollection struct {
	collection mongoI.Collection
}

func (m *mongoDbCtxCollection) Database() *mongo.Database {
	return m.collection.Database()
}

func (m *mongoDbCtxCollection) InsertOne(document interface{}) (*mongo.InsertOneResult, error) {
	return m.collection.InsertOne(getContext(), document)
}

func (m *mongoDbCtxCollection) InsertMany(documents []interface{}) (*mongo.InsertManyResult, error) {
	return m.collection.InsertMany(getContext(), documents)
}

func (m *mongoDbCtxCollection) FindOne(filter interface{}) *mongo.SingleResult {
	return m.collection.FindOne(getContext(), filter)
}

func (m *mongoDbCtxCollection) Find(filter interface{}) (*mongo.Cursor, error) {
	return m.collection.Find(getContext(), filter)
}

func (m *mongoDbCtxCollection) UpdateOne(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return m.collection.UpdateOne(getContext(), filter, update)
}

func (m *mongoDbCtxCollection) UpdateMany(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return m.collection.UpdateMany(getContext(), filter, update)
}

func (m *mongoDbCtxCollection) DeleteOne(filter interface{}) (*mongo.DeleteResult, error) {
	return m.collection.DeleteOne(getContext(), filter)
}

func (m *mongoDbCtxCollection) DeleteMany(filter interface{}) (*mongo.DeleteResult, error) {
	return m.collection.DeleteMany(getContext(), filter)
}

// MustGetNewMongoDbClient ensures the same things as NewMongoDbClient,
// but fatals when an error is encountered.
func MustGetNewMongoDbClient(dbUri string) *mongo.Client {
	client, err := NewMongoDbClient(dbUri)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// NewMongoDbClient takes a database uri and ensures the client
// can connect to an existing server. It then leaves the client
// in a disconnected state.
//
// If there are errors with any of the above states,
// then an error is returned.
func NewMongoDbClient(dbUri string) (*mongo.Client, error) {
	cOptions := options.Client().ApplyURI(dbUri)
	client, err := mongo.NewClient(cOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(getContext())
	if err != nil {
		return nil, err
	}

	// Note, Connect does not indicate the server exists
	// or is running. So, this line is meant to do that.
	err = client.Ping(getContext(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}

// MustMakeNewMongoDbCtxCollectionWithClient does the same thing as
// NewMongoCtxCollectionWithClient, but fatals on an error.
func MustMakeNewMongoDbCtxCollectionWithClient(client *mongo.Client, dbName string, collName string) mongoI.CtxCollection {
	coll, err := NewMongoDbCtxCollectionWithClient(client, dbName, collName)
	if err != nil {
		log.Fatal(err)
	}

	return coll
}

// NewMongoCtxCollectionWithClient creates a CtxCollection to the given collection in
// the database.
//
// Errors out if any of the inputs are either nil or empty.
func NewMongoDbCtxCollectionWithClient(client *mongo.Client, dbName string, collName string) (mongoI.CtxCollection, error) {
	if client == nil {
		return nil, errNilClient
	}

	if dbName == "" {
		return nil, errEmptyDbName
	}

	if collName == "" {
		return nil, errEmptyCollName
	}

	collection := client.Database(dbName).Collection(collName)
	return NewMongoDbCtxCollection(collection), nil
}

// NewMongoCtxCollection wraps pdv's Collection interface with the CtxCollection
// interface.
func NewMongoDbCtxCollection(coll mongoI.Collection) mongoI.CtxCollection {
	return &mongoDbCtxCollection{
		collection: coll,
	}
}
