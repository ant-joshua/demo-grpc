package models

type Product struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Price int32  `json:"price"`
}
