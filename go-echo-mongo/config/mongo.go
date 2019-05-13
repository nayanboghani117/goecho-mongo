package config

import (
	"gopkg.in/mgo.v2"
	"log"
)

func GetMongoDb() (*mgo.Collection,error){
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	db:= session.DB("demo").C("user")

	return db,nil
}