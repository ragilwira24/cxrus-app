package main

import (
	"cxrus-app/routes"
	"cxrus-app/util/db"
	"log"
	"net/http"
)

func main() {

	db.InitDB()
	log.Println("Initializing DB Success")

	router := routes.InitBaseRoutes()
	log.Println("All Routes Has Been Initialized")
	err := http.ListenAndServe(":"+"8080", router)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Server is start using port 3000")
	}

}
