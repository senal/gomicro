package main

import (
	"context"
	"fmt"
	"log"
	"logger/data"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRpcPort = "50001"
)

// mongo drivers
// go.mongodb.org/mongo-driver/mongo
// go.mongodb.org/mongo-driver/mongo/options

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {

	// connect to mongo
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	client = mongoClient

	// create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	go app.serve()

}

func (app *Config) serve() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func connectToMongo() (*mongo.Client, error) {
	// create options
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURL).SetServerAPIOptions(serverApi)
	opts.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}

	return client, err
}
