package controller

import (
	"errors"
	"fmt"
	"log"

	"github.com/ddld93/database/model"
	"github.com/ddld93/database/utils"
	

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DB_Connect struct {
	Session *mgo.Session
}
var (
	database = "forms"
	collection= "forms_collection"
)

func NewConnCtrl(host string, port int) *DB_Connect {
	url := fmt.Sprintf("mongodb://%s:%d", host, port)
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal("Error connecting to mogo database", err)
	}
	return &DB_Connect{Session: session}
}


func (sess *DB_Connect) NewEntry(form *model.Form) (string, error) {
	// validating user inputs
	forms, err  := utilities.FormModelValidate(form)
	if err != nil{
		return "",err
	}
	
	err2 := sess.Session.DB(database).C(collection).Insert(forms)
	if err2 != nil {
		fmt.Println("Error inserting new form ", err2)
		return "", err2
	}
	return "form submitted successfull", nil
}

func (sess *DB_Connect) GetForm(email string) (*model.Form, error) {
	form := model.Form{}
	err := sess.Session.DB(database).C(collection).Find(bson.M{"userEmail":email}).One(&form)
	if err != nil {
		return &form, errors.New("error getting form information from database")
	}
	return &form, nil
}

func (sess *DB_Connect) GetForms() ([]model.Form, error) {
	form := []model.Form{}
	err := sess.Session.DB(database).C(collection).Find(bson.M{}).All(&form)
	if err != nil {
		return form, errors.New("error getting forms ")
	}	
	return form, nil
}

