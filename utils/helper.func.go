package utilities

import (
	"errors"
	"github.com/ddld93/database/model"
)

func FormModelValidate(user *model.Form)  (*model.Form, error){
	// checking required fields 
	if user.FullName == "" {
		return user, errors.New("full name field cannot be empty")
	}
	if user.Email == "" {
		return user, errors.New("email field cannot be empty")
	}
	if user.Password == "" {
		return user, errors.New("password field cannot be empty")
	}
	
	if user.Phone == "" {
		return user, errors.New("phone field cannot be empty")
	}
	
	// assigning default value

	user.Role = "client"
	return user, nil
}