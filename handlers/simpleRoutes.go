package handlers

import (
	"html/template"
	"net/http"
)

func ViewHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		t, _ := template.ParseFiles("./views/error.gohtml")
		myvar := map[string]interface{}{"MyVar": "404 Page was not found"}
		t.Execute(w, myvar)
	} else {
		t, _ := template.ParseFiles("./views/index.gohtml")
		myvar := map[string]interface{}{"MyVar": "Hello World"}
		t.Execute(w, myvar)
	}

}

func DashboardView(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dashboard/" {
		t, _ := template.ParseFiles("./views/error.gohtml")
		myvar := map[string]interface{}{"MyVar": "404 Page was not found"}
		t.Execute(w, myvar)
	} else {
		t, _ := template.ParseFiles("./views/helth/Doctor/dashboard.gohtml")
		myvar := map[string]interface{}{"MyVar": "Hello World"}
		t.Execute(w, myvar)
	}

}

func UserView(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user/" {
		t, _ := template.ParseFiles("./views/error.gohtml")
		myvar := map[string]interface{}{"MyVar": "404 Page was not found"}
		t.Execute(w, myvar)
	} else {
		t, _ := template.ParseFiles("./views/helth/Doctor/user.gohtml")
		myvar := map[string]interface{}{"MyVar": "Hello World"}
		t.Execute(w, myvar)
	}
}

func PatDetails(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/PatDetails/" {
		t, _ := template.ParseFiles("./views/error.gohtml")
		myvar := map[string]interface{}{"MyVar": "404 Page was not found"}
		t.Execute(w, myvar)
	} else {
		t, _ := template.ParseFiles("./views/helth/Doctor/patient.gohtml")
		myvar := map[string]interface{}{"MyVar": "Hello World"}
		t.Execute(w, myvar)
	}
}

/*func AddHis(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/AddHis/" {
		t, _ := template.ParseFiles("./views/error.html")
		myvar := map[string]interface{}{"MyVar": "404 Page was not found"}
		t.Execute(w, myvar)
	} else {
		t, _ := template.ParseFiles("./views/helth/Doctor/table.html")
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
		t, _ := template.ParseFiles("./views/helth/Doctor/history.html")
		myvar := map[string]interface{}{"MyVar": "Hello World"}
		t.Execute(w, myvar)
	}
}
*/
