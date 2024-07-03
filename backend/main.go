package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"product/handler"
	"net/http"
	"product/middleware"
)

func main() { 
	r := mux.NewRouter()

	r.HandleFunc("/",handler.Home)
	r.HandleFunc("/token", handler.SetToken).Methods("GET")
	r.HandleFunc("/products", handler.GetProducts).Methods("GET")
	r.HandleFunc("/products", handler.PostProduct).Methods("POST")
	r.HandleFunc("/products/{id}", handler.DeleteProduct).Methods("DELETE")

	fmt.Println("Listening on port 8080") 
	log.Fatal(http.ListenAndServe(":8080", middleware.GetCorsConfig().Handler(r)))
}