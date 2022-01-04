package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

const uri = "mongodb://localhost:27017"

func main() {

	credential := options.Credential{
		Username: "jack",
		Password: "12345678",
	}

	clientOptions := options.Client()
	clientOptions.ApplyURI(uri).SetAuth(credential)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	// coll := client.Database("db_enigma").Collection("students")
	productColl := client.Database("db_enigma").Collection("products")

	log.Println("Successfully connected and ping")

	//const shortForm = "2006-01-02"
	//dt, _ := time.Parse(shortForm, "2022-01-05")
	//newStudent := Student{
	//	Id:       primitive.NewObjectID(),
	//	Name:     "John",
	//	Gender:   "M",
	//	Age:      28,
	//	JoinDate: dt,
	//	IdCard:   "206",
	//	Senior:   false,
	//}
	//
	//InsertOneStudent(ctx, coll, newStudent)

	// FindAllStudent(ctx, coll)
	// FindStudentByGenderAndAge(ctx, coll, "F", 22)
	// FindStudentByGenderAndAge2(ctx, coll, "M", 26)
	// CountStudentByAge(ctx, coll, 26)
	CountProductByCategory(ctx, productColl, "handphone")

}
