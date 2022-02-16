package utilities

import (
	"errors"
	"log"
	"net/http"

	"github.com/ddld93/database/model"
)

func FormModelValidate(user *model.Form)  (*model.Form, error){
	// checking required fields 
	if user.FullName == "" {
		return user, errors.New("full name field cannot be empty")
	}
	// if user.Email == "" {
	// 	return user, errors.New("email field cannot be empty")
	// }
	// if user.Password == "" {
	// 	return user, errors.New("password field cannot be empty")
	// }
	
	// if user.Phone == "" {
	// 	return user, errors.New("phone field cannot be empty")
	//}
	
	// assigning default value

	//user.Role = "client"
	return user, nil
}
func FormFlagToggle(email string) error{
	
	url := "https://api.paystack.co/transaction/verify/" + email

    // Create a new request using http
    req, _ := http.NewRequest("GET", url, nil)

    // Send req using http Client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERROR] -", err)
		return err
    }
    defer resp.Body.Close()
	return nil
}