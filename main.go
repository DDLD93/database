package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/rs/cors"


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
	
    r.HandleFunc("/newform",route.CreateForm).Methods("POST")
    r.HandleFunc("/getform/{email}",route.GetFormByEmail).Methods("GET")
	r.HandleFunc("/getforms", route.GetAllForms).Methods("GET") 
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
		
	})

	handler := c.Handler(r)
    
	fmt.Printf("Server listening on port %v", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal("Error starting server !! ", err)
	}


}