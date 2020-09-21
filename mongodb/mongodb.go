package mongodb

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mjyaqubi/gin-helpers/pagination"
)

// Client - MongodbClient struct
type Client struct {
	Database *mongo.Database
}

// ClientOptions - Mongodb Client Options
type ClientOptions struct {
	Host string
	Port string
	Name string
}

// NewClient - New Mongodb Client
func NewClient(clientOptions *ClientOptions) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := "mongodb://" + clientOptions.Host + ":" + clientOptions.Port
	options := options.Client().ApplyURI(uri)
	options.SetMaxPoolSize(10)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}

	return &Client{
		Database: client.Database(clientOptions.Name),
	}, nil
}

// Aggregate - aggregate query
func (db *Client) Aggregate(collection string, pipeline interface{}) (*mongo.Cursor, error) {
	return db.Database.Collection(collection).Aggregate(context.TODO(), pipeline)
}

// Find - find all documents
func (db *Client) Find(collection string, filters *bson.M, pagination pagination.Query) (*mongo.Cursor, error) {
	opt := options.Find()
	opt.SetSkip(pagination.Skip)
	opt.SetLimit(pagination.Limit)
	return db.Database.Collection(collection).Find(context.TODO(), filters, opt)
}

// FindOne - find one document
func (db *Client) FindOne(collection string, filters *bson.M) *mongo.SingleResult {
	return db.Database.Collection(collection).FindOne(context.TODO(), filters)
}

// InsertOne - insert one document
func (db *Client) InsertOne(collection string, document *bson.M) (*mongo.InsertOneResult, error) {
	return db.Database.Collection(collection).InsertOne(context.TODO(), document)
}

// InsertOneIfNotExist - insert one document if not already exist
func (db *Client) InsertOneIfNotExist(
	collection string,
	filters *bson.M,
	document *bson.M,
) (*mongo.InsertOneResult, error) {
	count, err := db.Database.Collection(collection).CountDocuments(context.TODO(), filters)
	if err != nil {
		return nil, err
	} else if count > 0 {
		return nil, errors.New("the document already exist")
	}

	return db.Database.Collection(collection).InsertOne(context.TODO(), document)
}

// FindOneAndUpdate - find one document and update
func (db *Client) FindOneAndUpdate(collection string, filters *bson.M, document *bson.M) *mongo.SingleResult {
	return db.Database.Collection(collection).FindOneAndUpdate(context.TODO(), filters, document)
}
