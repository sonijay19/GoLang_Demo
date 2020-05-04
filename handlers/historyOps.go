package handlers

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	md "../models"

	"go.mongodb.org/mongo-driver/bson"
)

func AddHis(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/AddHis/" {
		t, _ := template.ParseFiles("./views/error.html")
		myvar := map[string]interface{}{"MyVar": "404 Page was not found"}
		t.Execute(w, myvar)
	} else {
		t, _ := template.ParseFiles("./views/helth/Doctor/table.gohtml")
		myvar := map[string]interface{}{"MyVar": "Hello World"}
		t.Execute(w, myvar)
	}
}

func ShowHis(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ShowHis/" {
		t, _ := template.ParseFiles("./views/error.html")
		myvar := map[string]interface{}{"MyVar": "404 Page was not found"}
		t.Execute(w, myvar)
	} else {
		collection := md.GetCollection("doctor")
		filter := bson.D{{"_id", "sonijay1"}}
		var result md.PatHis
		err := collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
		myvar := map[string]interface{}{"history": result.History}
		t, _ := template.ParseFiles("./views/helth/Doctor/history.gohtml")
		t.Execute(w, myvar)
	}
}
