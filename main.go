package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ddld93/database/controller"
	"github.com/ddld93/database/routes"
	"github.com/gorilla/mux"

)

// func init() {

// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// }
// }

func main()  {
	port := "3000"
	FormCtrl := controller.NewConnCtrl("mongo", 27017)
	route := routes.FormRoute{FormCtrl: FormCtrl}

	r := mux.NewRouter()
	
    r.HandleFunc("/api/forms/newform",route.Form).Methods("POST")
    r.HandleFunc("/api/forms/getform/{id}",route.GetFormById ).Methods("GET")
	r.HandleFunc("/api/forms/getforms", route.GetAllForms).Methods("GET") 



    http.Handle("/", r)

	fmt.Printf("Server listening on port %v", port)
	if err := http.ListenAndServe(":"+ port, r); err != nil {
		log.Fatal("Error starting server !! ", err)
	}


}