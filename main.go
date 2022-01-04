package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"mongodb-course/db"
	"mongodb-course/utils"
)

func main() {
	mdb, err := db.InitResource()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := utils.InitContext()
	defer cancel()
	defer func() {
		if err = mdb.Db.Client().Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	//Read Preference = Default Mode. All operation read from current replica set Primary
	if err := mdb.Db.Client().Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

}
