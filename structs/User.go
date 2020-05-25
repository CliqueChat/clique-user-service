package structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/genproto/googleapis/type/date"
)

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	FName    string             `json:"fname" bson:"fname,omitempty"`
	LName    string             `json:"lname" bson:"lname,omitempty"`
	UserName string             `json:"username" bson:"username,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	BirthDay date.Date          `json:"birthday" bson:"birthday,omitempty"`
	Bio      string             `json:"bio" bson:"bio,omitempty"`
}
