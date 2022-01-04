package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"mongodb-course/db"
	"mongodb-course/model"
	"mongodb-course/utils"
)

type IStudentRepository interface {
	GetAll() ([]*model.Student, error)
	GetOneByUsername(name string) (*model.Student, error)
	CreateOne(student model.Student) (*model.Student, error)
}

type StudentRepository struct {
	repo *mongo.Collection
}

func NewStudentRepository(resource *db.Resource) IStudentRepository {
	studentCollection := resource.Db.Collection("students")
	studentRepository := &StudentRepository{repo: studentCollection}
	return studentRepository
}

func (s *StudentRepository) GetAll() ([]*model.Student, error) {
	var students []*model.Student
	ctx, cancel := utils.InitContext()
	defer cancel()
	cursor, err := s.repo.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var student model.Student
		err = cursor.Decode(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}
	return students, nil
}

func (s *StudentRepository) GetOneByUsername(name string) (*model.Student, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()
	var student model.Student
	err := s.repo.FindOne(ctx, bson.M{"name": name}).Decode(&student)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &student, nil
}

func (s *StudentRepository) CreateOne(student model.Student) (*model.Student, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	student.Id = primitive.NewObjectID()
	_, err := s.repo.InsertOne(ctx, student)

	if err != nil {
		return nil, err
	}

	return &student, nil
}
