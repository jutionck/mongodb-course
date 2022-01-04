package usecase

import (
	"mongodb-course/model"
	"mongodb-course/repository"
)

type IStudentUseCase interface {
	NewRegistration(student model.Student) (*model.Student, error)
	FindStudentInfoByName(name string) (*model.Student, error)
}

type StudentUseCase struct {
	repo repository.IStudentRepository
}

func NewStudentUseCase(studentRepository repository.IStudentRepository) IStudentUseCase {
	return &StudentUseCase{studentRepository}
}

func (s *StudentUseCase) NewRegistration(student model.Student) (*model.Student, error) {
	return s.repo.CreateOne(student)
}

func (s *StudentUseCase) FindStudentInfoByName(name string) (*model.Student, error) {
	return s.repo.GetOneByUsername(name)
}
