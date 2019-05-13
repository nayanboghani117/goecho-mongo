package main

import (
	"github.com/labstack/echo"
	"github.com/swaggo/echo-swagger"
	"go-echo-mongo/api/handler"
	_ "go-echo-mongo/docs"
	"go-echo-mongo/mware"
	"gopkg.in/mgo.v2"
	"log"
)

func main() {
	e := echo.New()
	mware.Mainmiddleware(e)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	db, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Create indices
	if err = db.Copy().DB("demo").C("user").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	// Initialize handler
	h := &handler.Handler{DB: db}


	e.GET("/getuserbyid/:id", handler.GetUserByID(h))
	e.POST("/createuser", handler.CreateUser(h))
	e.PUT("/updateuser/:id", handler.UpdateUser(h))
	e.DELETE("/deleteuser/:id", handler.DeleteUser(h))
	e.POST("/signin",handler.SignIn(h))
	e.GET("/getusers",handler.GetUsers(h))
	e.GET("/private", handler.Private, mware.IsLoggedIn)
	e.GET("/admin", handler.Private, mware.IsLoggedIn,mware.IsAdmin)

	e.Logger.Fatal(e.Start(":8000"))

}
