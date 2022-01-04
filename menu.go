package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mongodb-course/model"
)

func InsertOneStudent(ctx context.Context, coll *mongo.Collection, student model.Student) {
	newId, err := coll.InsertOne(ctx, student)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID baru : %v \n", (*newId).InsertedID)
}

func FindAllStudent(ctx context.Context, coll *mongo.Collection) {
	var results []bson.M
	allDocumentCursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		log.Fatalln(err)
	}

	defer func(allDocumentCursor *mongo.Cursor, ctx context.Context) {
		err := allDocumentCursor.Close(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}(allDocumentCursor, ctx)

	err = allDocumentCursor.All(ctx, &results)
	if err != nil {
		log.Fatalln(err)
	}

	for _, doc := range results {
		fmt.Printf("_id:%v, name: %v, age: %v \n", doc["_id"], doc["name"], doc["age"])
	}
}

func FindStudentByGenderAndAge(ctx context.Context, coll *mongo.Collection, gender string, age int) {
	filterMaleAndAge20 := bson.D{
		{"$and", bson.A{
			bson.D{{"gender", gender}},
			bson.D{{"age", age}},
		}},
	}
	projection := bson.D{
		{"_id", 0},
		{"name", 1},
	}
	findOpts := options.Find().SetProjection(projection)

	var maleAndAge20Result []bson.M
	maleAndAge20Cursor, err := coll.Find(ctx, filterMaleAndAge20, findOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer maleAndAge20Cursor.Close(ctx)
	if err = maleAndAge20Cursor.All(ctx, &maleAndAge20Result); err != nil {
		panic(err)
	}
	for _, doc := range maleAndAge20Result {
		fmt.Printf("name: %v\n", doc["name"])
	}
}

func FindStudentByGenderAndAge2(ctx context.Context, coll *mongo.Collection, gender string, age int) {
	filterMaleAndAge20 := bson.D{
		{"$and", bson.A{
			bson.D{{"gender", gender}},
			bson.D{{"age", age}},
		}},
	}
	projection := bson.D{
		{"_id", 1},
		{"name", 1},
		{"joinDate", 1},
	}
	findOpts := options.Find().SetProjection(projection)

	maleAndAge20Result := make([]*model.Student, 0)
	maleAndAge20Cursor, err := coll.Find(ctx, filterMaleAndAge20, findOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer maleAndAge20Cursor.Close(ctx)
	for maleAndAge20Cursor.Next(ctx) {
		var row model.Student
		err := maleAndAge20Cursor.Decode(&row)
		if err != nil {
			log.Fatal(err)
		}
		maleAndAge20Result = append(maleAndAge20Result, &row)
	}
	layoutISO := "2006-01-02"

	for _, doc := range maleAndAge20Result {
		t := doc.JoinDate
		fmt.Printf("ID: %v, name: %v, join date: %v\n", doc.Id, doc.Name, t.Format(layoutISO))
	}
}

func CountStudentByAge(ctx context.Context, coll *mongo.Collection, age int) {
	count, err := coll.CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Student total: %d\n", count)

	filter := bson.D{{"age", age}}
	count, err = coll.CountDocuments(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Student total with age %v: %d\n", age, count)
}

func CountProductByCategory(ctx context.Context, productColl *mongo.Collection, category string) {
	matchStage := bson.D{{"$match", bson.D{{"category", category}}}}
	groupStage := bson.D{{"$group", bson.D{
		{"_id", "$category"},
		{"count", bson.D{{"$sum", 1}}},
	},
	}}

	aggCursor, err := productColl.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		log.Fatal(err)
	}
	defer aggCursor.Close(ctx)
	var aggInfo []bson.M
	if err = aggCursor.All(ctx, &aggInfo); err != nil {
		log.Fatal(err)
	}
	for _, info := range aggInfo {
		fmt.Println()
		fmt.Printf("Group: %v, Total: %v\n", info["_id"], info["count"])
	}

}
