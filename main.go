package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ddld93/database/controller"
	"github.com/ddld93/database/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
}
}

func main()  {
	port := os.Getenv("PORT")
	FormCtrl := controller.NewConnCtrl("localhost", 27017)
	route := routes.FormRoute{FormCtrl: FormCtrl}

	r := mux.NewRouter()
	
    r.HandleFunc("/form",route.Form).Methods("POST")
    r.HandleFunc("/GetForm",route.GetFormById ).Methods("GET")
	r.HandleFunc("/GetForms", route.GetAllForms).Methods("GET") 



    http.Handle("/", r)

	fmt.Printf("Server listening on port %v", port)
	if err := http.ListenAndServe(":"+ port, r); err != nil {
		log.Fatal("Error starting server !! ", err)
	}


}