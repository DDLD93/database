package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Form struct {
	Id         bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	UserEmail  string		 `json:"userEmail" bson:"userEmail"`
	FullName   string        `json:"fullName" bson:"fullName"`
	Program    string        `json:"program" bson:"program"` //midwifery or nursing
	Source     string        `json:"source" bson:"source"`   // friends relatives peer-group alumnii media wordOfMouth website educationFair collegeStaff socialMedia
	ProfilePic string        `json:"profilePic" bson:"profilePic"`
	CreatedAt  time.Time     `json:"createAt" bson:"creatAt"`
}
