package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
	"strings"

	"github.com/ddld93/database/controller"
	"github.com/ddld93/database/model"
	utilities "github.com/ddld93/database/utils"
	"github.com/gorilla/mux"

)

type FormRoute struct {
	FormCtrl *controller.DB_Connect
}
type CustomResponse struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}


func (ur *FormRoute) NewForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	reqToken := r.Header.Get("Authorization")
	// checking if request carries a valid token
	if reqToken == "" {
		resp := CustomResponse{
			Message: "Token not Found",
			Description: "Bearer token not included in request",}
			json.NewEncoder(w).Encode(resp)
		return
	}
		splitToken := strings.Split(reqToken, "Bearer ")
		token := splitToken[1]
		 _,err := utilities.TokenMaker.VerifyToken(token)
		if err != nil{
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Println(err)
			json.NewEncoder(w).Encode(err)
			return
		}
		fmt.Println("token valid")
		form := model.Form{}
		err2 := json.NewDecoder(r.Body).Decode(&form)
	if err2 != nil {
		resp := CustomResponse{Message: err.Error(), Description: "Error Decoding request body"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return
	}

		resp,err4 := ur.FormCtrl.NewEntry(&form)
		if err4 != nil{
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(resp)

}


func (ur *FormRoute) GetFormById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	reqToken := r.Header.Get("Authorization")
	// checking if request carries a valid token
	if reqToken == "" {
		resp := CustomResponse{
			Message: "Token not Found",
			Description: "Bearer token not included in request",}
			json.NewEncoder(w).Encode(resp)
		return
	}
	
		splitToken := strings.Split(reqToken, "Bearer ")
		token := splitToken[1]
		 _,err := utilities.TokenMaker.VerifyToken(token)
		if err != nil{
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Println(err)
			json.NewEncoder(w).Encode(err)
			return
		}
		idNo:= mux.Vars(r)
		id := idNo["id"]
		form,err := ur.FormCtrl.GetForm(id)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
		}
		content, err := ioutil.ReadFile(form.ProfilePic)
		
		form.ProfilePic = string(content)
		if err != nil {
			fmt.Println(err,"\n")
		}
		

		json.NewEncoder(w).Encode(form)
	
}

func (ur *FormRoute) GetAllForms(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	reqToken := r.Header.Get("Authorization")
	// checking if request carries a valid token
	if reqToken == "" {
		resp := CustomResponse{
			Message: "Token not Found",
			Description: "Bearer token not included in request",}
			json.NewEncoder(w).Encode(resp)
		return
	}
		splitToken := strings.Split(reqToken, "Bearer ")
		token := splitToken[1]
		 _,err := utilities.TokenMaker.VerifyToken(token)
		if err != nil{
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Println(err)
			json.NewEncoder(w).Encode(err)
			return
		}
		forms,err := ur.FormCtrl.GetForms()
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.WriteHeader(http.StatusOK)
	
		json.NewEncoder(w).Encode(forms)
	
}


func (ur *FormRoute) Form(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	// checking if request carries a valid token
	if reqToken == "" {
		resp := CustomResponse{
			Message:     "Token not Found",
			Description: "Bearer token not included in request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	token := splitToken[1]
	payload, err := utilities.TokenMaker.VerifyToken(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}


    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)
    // FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    file, _, err := r.FormFile("image")
    if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
    }
    defer file.Close()
   

    // Create a temporary file within our images directory that follows
    // a particular naming pattern
    tempFile, err := ioutil.TempFile("images/", payload.Username+"*"+".png")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	name := tempFile.Name()
    defer tempFile.Close()
	fmt.Println(name)

    // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
    }
    // write this byte array to our temporary file
   _,err = tempFile.Write(fileBytes)
   if err != nil {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err)
	return
   }
    // return that we have successfully uploaded our file!
	form := model.Form{
		FullName: r.Form.Get("fullName"),
		Program: r.Form.Get("program"),
		Source: r.Form.Get("source"),
		ProfilePic: name,
	
	}
	resp,err := ur.FormCtrl.NewEntry(&form)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	
	json.NewEncoder(w).Encode(resp)	
}
