package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Firstname string		`json:"firstname" bson:"firstname"`
	Lastname  string		`json:"lastname" bson:"lastname"`
	EmailId   string		`json:"email_id" bson:"email_id"`
	ContactNo string		`json:"contact_no" bson:"contact_no"`
	Dob       string        `json:"dob" bson:"dob"` // Date of bitrh
}