package db

import (
    "fmt"
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/mr-amirfazel/chisai/config"
    "log"
    "time"
)

var client *mongo.Client

func ConnectDB() error {
    conf, e := config.LoadConfig()
	if e != nil {
        fmt.Println("Error loading config: %s\n", e)
        return e
    }

    clientOptions := options.Client().ApplyURI(conf.Database.ConnectionString)
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

    return nil
}

func GetDBInstance() *mongo.Client {
    return client
}