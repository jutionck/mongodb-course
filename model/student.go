package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Student struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `json:"name" binding:"required"`
	Gender   string             `json:"gender"`
	Age      int                `json:"age"`
	JoinDate time.Time          `json:"joinDate"`
	IdCard   string             `json:"idCard"`
	Senior   bool               `json:"senior"`
}
