package model

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Form struct{
	Id    bson.ObjectId 	`json:"id,omitempty" bson:"_id,omitempty"`
	FullName string        	`json:"fullName" bson:"fullName"`				
	Program string        	`json:"program" bson:"program"`	//midwifery or nursing
	Source string        	`json:"source" bson:"source"`	// friends relatives peer-group alumnii media wordOfMouth website educationFair collegeStaff socialMedia
	ProfilePic string       `json:"profilePic" bson:"profilePic"`		
	CreatedAt time.Time	`json:"createAt" bson:"creatAt"`
}