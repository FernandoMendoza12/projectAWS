package db

import (
	"aws/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN *mongo.Client
var DatabaseName string

func ConnectDB(ctx context.Context) error {
	user := ctx.Value(models.Key("user")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)
	connString := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)

	var clientOptions = options.Client().ApplyURI(connString)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("> Error al connectar a mongo Atlas" + err.Error())
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Error al durante el Ping:" + err.Error())
	}

	fmt.Println("> Conexion exitosa a MongoDB")
	MongoCN = client
	DatabaseName = ctx.Value(models.Key("database")).(string)

	return nil

}

func ConnectedDB() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}
