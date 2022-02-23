package routes

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"net/http"

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
	Payload interface{} `json:"payload"`
	Error error `json:"error"`
}



func (ur *FormRoute) GetFormByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
		params:= mux.Vars(r)
		email := params["email"]
		form,err := ur.FormCtrl.GetForm(email)
		if err != nil{
			resp := CustomResponse{
				Status: "failed",
				Message: err.Error(),
				Error: err,
				
			}
			   json.NewEncoder(w).Encode(resp)
			   return
		}
	
		w.WriteHeader(http.StatusOK)
		resp := CustomResponse{
			Status: "success",
			Message: "Forms Retrieved successifully",
			Payload: form,
		}
		   json.NewEncoder(w).Encode(resp)
	
}

func (ur *FormRoute) GetAllForms(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	
		forms,err := ur.FormCtrl.GetForms()
		if err != nil{
			resp := CustomResponse{
				Status: "failed",
				Message: "Error Retrieving forms",
				Error: err,
				
			}
			   json.NewEncoder(w).Encode(resp)
			return
		}
		w.WriteHeader(http.StatusOK)
	 resp := CustomResponse{
		 Status: "success",
		 Message: "Forms Retrieved successifully",
		 Payload: forms,
	 }
		json.NewEncoder(w).Encode(resp)
	
}


func (ur *FormRoute) CreateForm(w http.ResponseWriter, r *http.Request) {
	
	
    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)
    // FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    file, _, err := r.FormFile("image")
    if err != nil {
		resp := CustomResponse{
			Status: "failed",
			Message: "Error getting images from form",
			Error: err,
		}
		   json.NewEncoder(w).Encode(resp)
		return
    }
    defer file.Close()
   
    userEmail := r.Form.Get("userEmail")
    // Create a temporary file within our images directory that follows
    // a particular naming pattern
    tempFile, err := ioutil.TempFile("images", userEmail+"*"+".png")
    if err != nil {
        resp := CustomResponse{
			Status: "failed",
			Message: "Error error creating temp image file",
			Error: err,
		}
		   json.NewEncoder(w).Encode(resp)
		return
	}
	name := tempFile.Name()
    defer tempFile.Close()

    // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        resp := CustomResponse{
			Status: "failed",
			Message: "Error processing image bytes",
			Error: err,
		}
		   json.NewEncoder(w).Encode(resp)
		return
    }
    // write this byte array to our temporary file
   _,err = tempFile.Write(fileBytes)
   if err != nil {
	resp := CustomResponse{
		Status: "failed",
		Message: "Error writing byte image",
		Error: err,
	}
	   json.NewEncoder(w).Encode(resp)
	return
   }
    // return that we have successfully uploaded our file!
	form := model.Form{
		UserEmail: r.Form.Get("userEmail"),
		FullName: r.Form.Get("fullName"),
		Program: r.Form.Get("program"),
		Source: r.Form.Get("source"),
		ProfilePic: name,
		CreatedAt: time.Now(),
	}
	
	resp,err := ur.FormCtrl.NewEntry(&form)
	if err != nil {
		resp := CustomResponse{
			Status: "failed",
			Message: "Error adding form data",
			Error: err,
		}
		   json.NewEncoder(w).Encode(resp)
		return
	}
	err1 := utilities.FormFlagToggle(userEmail)
	if err1 != nil {
		resp := CustomResponse{
			Status: "failed",
			Message: "Error updating user state",
			Error: err,
		}
		   json.NewEncoder(w).Encode(resp)
		return
	}

	response:= CustomResponse{
		Status: "success",
		Message: resp,
		Payload: form,
	}
	json.NewEncoder(w).Encode(response)	
}
