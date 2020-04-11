// mongo package provides the interfaces to either simplify or to enable mock testing
// with Mongo's entities.

package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection is an interface containing a subset of the methods used in Mongo's
// Collection. This is mostly for mock testing purposes.
// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection
type Collection interface {
	// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.Database
	Database() *mongo.Database
	// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.InsertOne
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.InsertMany
	InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.FindOne
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.Find
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.UpdateOne
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.UpdateMany
	UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.DeleteOne
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.DeleteMany
	DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}

// CtxCollection is an interface containing the same methods as Collection, but provides
// a default context and options for each method to standardize the way the client
// makes calls.
type CtxCollection interface {
	// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.Database
	Database() *mongo.Database
	// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.InsertOne
	InsertOne(document interface{}) (*mongo.InsertOneResult, error)
	// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.InsertMany
	InsertMany(documents []interface{}) (*mongo.InsertManyResult, error)
	// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.FindOne
	FindOne(filter interface{}) *mongo.SingleResult
	// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.Find
	Find(filter interface{}) (*mongo.Cursor, error)
	// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.UpdateOne
	UpdateOne(filter interface{}, update interface{}) (*mongo.UpdateResult, error)
	// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.UpdateMany
	UpdateMany(filter interface{}, update interface{}) (*mongo.UpdateResult, error)
	// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.DeleteOne
	DeleteOne(filter interface{}) (*mongo.DeleteResult, error)
	// See: https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.DeleteMany
	DeleteMany(filter interface{}) (*mongo.DeleteResult, error)
}
