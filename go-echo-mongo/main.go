package main

import (
	"github.com/labstack/echo"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go-echo-mongo/api/handler"
	_ "go-echo-mongo/docs"
	"go-echo-mongo/mware"
	"gopkg.in/mgo.v2"
	"log"
)

func main() {


	// @title Swagger Example API
	// @version 1.0
	// @description This is a sample server user server.
	// @termsOfService http://swagger.io/terms/

	// @contact.name API Support
	// @contact.url http://www.swagger.io/support
	// @contact.email support@swagger.io

	// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

	// @host localhost:8080
	// @BasePath /v1

	e := echo.New()
	mware.Mainmiddleware(e)

	db, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		e.Logger.Fatal(err)
	}

	if err = db.Copy().DB("demo").C("user").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}


	h := &handler.Handler{DB: db}

	g := e.Group("/v1")
	user := g.Group("/user")
	user.GET("/getuserbyid/:id", handler.GetUserByID(h))
	user.POST("/createuser", handler.CreateUser(h))
	user.PUT("/updateuser/:id", handler.UpdateUser(h))
	user.DELETE("/deleteuser/:id", handler.DeleteUser(h))
	user.POST("/signin",handler.SignIn(h))
	user.GET("/getusers",handler.GetUsers(h))
	e.GET("/private", handler.Private, mware.IsLoggedIn)
	e.GET("/admin", handler.Private, mware.IsLoggedIn,mware.IsAdmin)

	e.GET("/swagger/*any", echoSwagger.WrapHandler)

	//url := echoSwagger.URL("http://localhost:8080/swagger/doc.json") //The url pointing to API definition
	//e.GET("/swagger/*", echoSwagger.EchoWrapHandler(url))

	e.Logger.Fatal(e.Start(":8080"))

}
