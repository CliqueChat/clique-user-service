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

type MongoRepository struct {
	Client *mongo.Client
}

var MongoRepo MongoRepository

func Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		panic(err)
	}

	MongoRepo.Client = client
}

func (m *MongoRepository) GetDatabase(dbName string) mongo.Database {
	return *m.Client.Database(dbName)
}

func (m *MongoRepository) GetCollection(colName string) mongo.Collection {
	return *m.Client.Database(mongoDBName).Collection(colName)
}
