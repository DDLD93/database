package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"strings"

	"github.com/ddld93/database/controller"
	"github.com/ddld93/database/model"
	"github.com/ddld93/database/utils"
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
		form,err := ur.FormCtrl.GetForm("tuygjh")
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

