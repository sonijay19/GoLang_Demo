package handlers

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	md "../models"
)

func DoctorAdd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/SignUp/" {
		t, _ := template.ParseFiles("./views/error.gohtml")
		myvar := map[string]interface{}{"MyVar": "404 Page was not found"}
		t.Execute(w, myvar)
	} else {
		collection := md.GetCollection("doctor")
		drInfo := md.DoctorInfo{
			DrName:     r.FormValue("drName"),
			DrmobNo:    r.FormValue("drmobNo"),
			DrEmail:    r.FormValue("drEmail"),
			DrPassword: r.FormValue("drPassword"),
			DrLicense:  r.FormValue("drLicense"),
			DrCity:     r.FormValue("drCity"),
			DrPincode:  r.FormValue("drPincode"),
			Domain:     r.FormValue("domain"),
		}
		_, err := collection.InsertOne(context.TODO(), drInfo)
		if err != nil {
			fmt.Fprint(w, err)
		}
		fmt.Fprint(w, "1")
	}
}
