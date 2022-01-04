package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"mongodb-course/db"
	"mongodb-course/model"
	"mongodb-course/repository"
	"mongodb-course/usecase"
	"mongodb-course/utils"
	"time"
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

	studentRepo := repository.NewStudentRepository(mdb)
	studentUseCase := usecase.NewStudentUseCase(studentRepo)

	student, err := studentUseCase.FindStudentInfoByName("Jution")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*student)

	const shortForm = "2006-01-02"
	dt, _ := time.Parse(shortForm, "2017-05-24")
	newStudent := model.Student{
		Id:       primitive.NewObjectID(),
		Name:     "Dika",
		Gender:   "M",
		Age:      24,
		JoinDate: dt,
		IdCard:   "208",
		Senior:   false,
	}

	registeredStudent, err := studentUseCase.NewRegistration(newStudent)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*registeredStudent)

}
