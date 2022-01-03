package datacontext

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	usr 		= "chessusr"
	pwd 		= "chess00usr11"
	host 		= "localhost"
	port 		= 27017
	database 	= "ChessMoveDB"
)

func GetDBContext(collection string) *mongo.Collection {

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", usr, pwd, host, port)
	fmt.Println("uri: ",uri)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database(database).Collection(collection)
}