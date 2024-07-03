package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"product/authentication"
	"product/database"
	"product/models"
	"product/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client = database.ConnectDB()
var data *mongo.Collection =  client.Database("Product").Collection("product")

// var SECRET_KEY = "secret"

// func GenerateJWT(id string) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"id": id,
// 	})
// 	return token.SignedString([]byte(SECRET_KEY))
// }

// func ValidateJWT(tokenStr string) (string, error) {
// 	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(SECRET_KEY), nil
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !token.Valid {
// 		return "", err
// 	}
// 	return claims["id"].(string), nil
// }

func Home (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func SetToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") 
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	token, err := authentication.GenerateJWT(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(token)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	token := r.Header.Get("Authorization")
	const Bearer = "Bearer "
	tokenString := token[len(Bearer):]

	_ , err := authentication.ValidateJWT(tokenString)
	if err != nil { 
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var products []models.Product

	cursor, err := data.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var product models.Product
		cursor.Decode(&product)
		products = append(products, product)
	}
	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(products) == 0 {
		http.Error(w, "No products found", http.StatusNotFound)
		return
	} 
	json.NewEncoder(w).Encode(products)	
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("Authorization")
	const Bearer = "Bearer "
	tokenString := token[len(Bearer):] 

	_ , err := authentication.ValidateJWT(tokenString)
	if err != nil { 
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)	
		return
	}
	product.ID = utils.GenerateID()

	_, err = data.InsertOne(context.TODO(), product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(product)  
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("Authorization")
	const Bearer = "Bearer "
	tokenString := token[len(Bearer):]

	_ , err := authentication.ValidateJWT(tokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id) 
	_, err = data.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}