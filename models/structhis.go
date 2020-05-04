package models

type His struct {
	Date     string
	Diseases string
}

type PatHis struct {
	ID      string `bson:"_id"`
	Age     string
	City    string
	History []His
}
