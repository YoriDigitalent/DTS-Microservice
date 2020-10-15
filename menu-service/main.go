package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/YoriDigitalent/DTS-Microservice/menu-service-handler"
	"log"
	"net/http"
)
	

func main()  {
	
	router := mux.NewRouter()

	router.Handle("/add-menu", http.HandlerFunc(handler.AddMenu))
	fmt.Println("Menu service listen on port: 8000")
	log.Panic(http.ListenAndServe(":8000", router))
}