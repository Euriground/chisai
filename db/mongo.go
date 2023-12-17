package db

import (
	"context"
	"fmt"
	"github.com/mr-amirfazel/chisai/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
)

func ConnectDB() error {
	conf, e := config.LoadConfig()
	if e != nil {
		fmt.Println("Error loading config: %s\n", e)
		return e
	}

    dbConf := conf.Database
	clientOptions := options.Client().ApplyURI(fmt.Sprintf(dbConf.ConnectionString, dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	database = client.Database("chisai")
	collection = database.Collection("URL")

	return nil
}

func GetDBInstance() *mongo.Client {
	return client
}

func GetCollection() *mongo.Collection {
	return collection
}
