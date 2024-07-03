package models


type Product struct {
	ID string `bson:"_id" json:"id,omitempty"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}