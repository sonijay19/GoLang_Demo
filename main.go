package main

import (
	"log"
	"net/http"

	hdl "./handlers"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", hdl.ViewHandle)
	http.HandleFunc("/dashboard/", hdl.DashboardView)
	http.HandleFunc("/user/", hdl.UserView)
	http.HandleFunc("/SignUp/", hdl.DoctorAdd)
	http.HandleFunc("/LogIn/", hdl.LogIn)
	http.HandleFunc("/PatDetails/", hdl.PatDetails)
	http.HandleFunc("/AddHis/", hdl.AddHis)
	http.HandleFunc("/ShowHis/", hdl.ShowHis)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
