package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"product/handler"
	"net/http"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/",handler.Home)
	r.HandleFunc("/token", handler.SetToken).Methods("GET")
	r.HandleFunc("/products", handler.GetProducts).Methods("GET")


	r.HandleFunc("/products", handler.PostProduct).Methods("POST")

	handler := cors.Default().Handler(r)
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}