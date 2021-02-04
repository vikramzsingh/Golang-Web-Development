package controllers

import (
	"Golang-Web-Development/CRUD-application/models"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/* or you can also access template by this technique
var TPL *template.Template

func  init() {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
}
*/
type UserController struct {
	session *mgo.Session
	TPL     *template.Template
}

// var oid bson.ObjectId
var ObjID = map[string]bson.ObjectId{}

func NewUserController(s *mgo.Session, TPL *template.Template) *UserController {
	return &UserController{s, TPL}
}

// CREATE USER
func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request) {

	// process form
	if req.Method == http.MethodPost { // checking method

		// process form values
		fname := req.FormValue("fname")
		lname := req.FormValue("lname")
		email := req.FormValue("email")
		contactno := req.FormValue("contactno")
		dob := req.FormValue("dob")
		// creating new ObjectId for mongoDB
		oid := bson.NewObjectId()

		ObjID["id"] = oid // inseting into map

		 u := models.User{oid, fname, lname, email, contactno, dob}



		if err := uc.session.DB("store").C("dogs").Insert(u); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // 500
			return
		}

		// Marshal into JSON
		uj, err := json.Marshal(u)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 200
		fmt.Fprintf(w, "Inserted Document DATA:\n %s\n", uj)
		return
	} // end of form process

	err := uc.TPL.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) // 404
		return
	}
}

// RETRIVE USER
func (uc UserController) GetUser(w http.ResponseWriter, req *http.Request) {
	var u []models.User

	oid := ObjID["id"]
	// Fetching/Retriving doc. data
	if err := uc.session.DB("store").C("dogs").Find(bson.M{"_id":oid}).All(&u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
		return
	}

	// Marshal into JSON
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)

	fmt.Fprintln(w, "\nGo data User struct type: ")
	for _, v := range u {
		fmt.Fprintln(w, v)
	}

}

// UPDATE USER
func (uc UserController) UpdateUser(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		fname := req.FormValue("fname")
		lname := req.FormValue("lname")
		//id := bson.ObjectIdHex("5fcd711411c37848a4901931")

		//u := models.User{}
		//var user models.User

		oid := ObjID["id"]

		err := uc.session.DB("store").C("dogs").Update(bson.M{"_id": oid}, bson.D{
				{"$set", bson.D{{"firstname", fname},{"lastname", lname}}},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // 500
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 200
		fmt.Fprintln(w, "data Updated")
		return
	}
	err := uc.TPL.ExecuteTemplate(w, "update.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) //404
	}
}

// DELETE data
func (uc UserController) DeleteUser(w http.ResponseWriter, req *http.Request) {
	//id := bson.ObjectIdHex("5fcd711411c37848a4901931")

	//u := models.User{}
	//var user models.User

	oid := ObjID["id"]

	err := uc.session.DB("store").C("dogs").RemoveId(oid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintln(w, "Data Deleted")
	return
}
