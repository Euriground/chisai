package db

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "time"
)

var client *mongo.Client

func ConnectDB() error {
    clientOptions := options.Client().ApplyURI("mongodb+srv://amirfazel45:Chisai@sample-cluster.087lnjs.mongodb.net/?retryWrites=true&w=majority")
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