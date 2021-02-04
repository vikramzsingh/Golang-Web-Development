package main

import (
	"net/http"
	"gopkg.in/mgo.v2"
	"fmt"
	"Golang-Web-Development/CRUD-application/controllers"
	"html/template"
)

var TPL *template.Template

func  init() {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	uc := controllers.NewUserController(getSession(TPL))
	http.HandleFunc("/", uc.CreateUser)
	http.HandleFunc("/getuser", uc.GetUser)
	http.HandleFunc("/updateuser", uc.UpdateUser)
	http.HandleFunc("/deleteuser", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", nil) // nil means using defaut servemux
}

// connection to mongodb
func getSession(TPL *template.Template) (*mgo.Session, *template.Template) {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		fmt.Println(err)
	}
	return s, TPL
}
