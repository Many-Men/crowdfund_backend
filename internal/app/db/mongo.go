package db

import (
	"context"
	"github.com/Many-Men/crowdfund_backend/config"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB(cfg *config.Config) *mongo.Database {
	//uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=Cluster0", cfg.MongoDB.User, cfg.MongoDB.Password, cfg.MongoDB.Host)

	uri := "mongodb://localhost:27017/?appName=MyApp"
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	db := client.Database(cfg.MongoDB.Database)

	return db
}
