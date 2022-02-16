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
	Status     string `json:"status"`
	Message string `json:"message"`
}



func (ur *FormRoute) GetFormById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	reqToken := r.Header.Get("Authorization")
	// checking if request carries a valid token
	if reqToken == "" {
		resp := CustomResponse{
			Status: "Token not Found",
			Message: "Bearer token not included in request",}
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
	
		json.NewEncoder(w).Encode(form)
	
}

func (ur *FormRoute) GetAllForms(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	reqToken := r.Header.Get("Authorization")
	// checking if request carries a valid token
	if reqToken == "" {
		resp := CustomResponse{
			Status: "failed",
			Message: "Bearer token not included in request",}
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


func (ur *FormRoute) CreateForm(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	
	// checking if request carries a valid token
	if reqToken == "" {
		resp := CustomResponse{
			Status:     "Token not Found",
			Message: "Bearer token not included in request"}
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
	fmt.Println("hiit")

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
    tempFile, err := ioutil.TempFile("images", payload.Username+"*"+".png")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	name := tempFile.Name()
    defer tempFile.Close()

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
		UserEmail: payload.Username,
		FullName: r.Form.Get("fullName"),
		Program: r.Form.Get("program"),
		Source: r.Form.Get("source"),
		ProfilePic: name,
	
	}
	resp,err := ur.FormCtrl.NewEntry(&form)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	err = utilities.FormFlagToggle(payload.Username)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	
	response:= CustomResponse{
		Status: "success",
		Message: resp,
	}
	json.NewEncoder(w).Encode(response)	
}
