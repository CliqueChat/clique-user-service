package repositories

import (
	"context"
	"github.com/CliqueChat/clique-user-service/helpers"
	"github.com/CliqueChat/clique-user-service/resources"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var prop = resources.GetApplicationProfile()
var mongoURI, _ = prop.Get(helpers.MongoURI)
var mongoDBName, _ = prop.Get(helpers.MongoDBName)

type MongoClient struct {
	Client *mongo.Client
	Ctx    context.Context
}

func (m *MongoClient) Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		panic(err)
	}
	m.Client = client
	m.Ctx = ctx
}

func (m *MongoClient) GetDatabase(dbName string) mongo.Database {
	return *m.Client.Database(dbName)
}

func (m *MongoClient) GetCollection(colName string) mongo.Collection {
	return *m.Client.Database(mongoDBName).Collection(colName)
}
