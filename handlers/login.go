package handlers

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	md "../models"
	"go.mongodb.org/mongo-driver/bson"
)

func LogIn(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/LogIn/" {
		t, _ := template.ParseFiles("./views/error.gohtml")
		myvar := map[string]interface{}{"MyVar": "404 Page was not found"}
		t.Execute(w, myvar)
	} else {
		fmt.Println("In Login go")
		collection := md.GetCollection("doctor")
		filter := bson.D{{"_id", r.FormValue("drEmail")}, {"drpassword", r.FormValue("drPassWord")}}
		fmt.Println(filter)
		var result md.DoctorInfo
		err := collection.FindOne(context.TODO(), filter).Decode(&result)
		fmt.Println(result)
		if err != nil {
			fmt.Println("sonijay")
			fmt.Fprint(w, "Password or Email is not match")
		} else {
			fmt.Fprint(w, "1")
		}
	}
}
