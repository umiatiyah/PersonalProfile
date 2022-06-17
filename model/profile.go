package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	NickName  string             `json:"nickname,omitempty" bson:"nickname,omitempty"`
	FullName  string             `json:"fullname,omitempty" bson:"fullname,omitempty"`
	Freelance bool               `json:"freelance" bson:"freelance,omitempty"`
	Email     string             `json:"email" bson:"email,omitempty"`
	Address1  string             `json:"address1" bson:"address1,omitempty"`
	Address2  string             `json:"address2" bson:"address2,omitempty"`
}
