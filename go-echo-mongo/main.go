package main

import (
	"github.com/labstack/echo"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go-echo-mongo/go-echo-mongo/api/handler"
	_ "go-echo-mongo/go-echo-mongo/docs"
	"go-echo-mongo/go-echo-mongo/mware"
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

	v1 := e.Group("/v1")
	{
		user := v1.Group("/users")
		{
			user.GET("/:id", handler.GetUserByID(h))
			user.GET("/", handler.GetUsers(h))
			user.POST("/", handler.CreateUser(h))
			user.POST("/login", handler.SignIn(h))
			user.PUT("/:id", handler.UpdateUser(h))
			user.DELETE("/:id", handler.DeleteUser(h))
		}


		auth := v1.Group("/auth")
		{
			auth.GET("/userauth", handler.Private, mware.IsLoggedIn)
			auth.GET("/admin", handler.Private, mware.IsLoggedIn,mware.IsAdmin)
		}
	}

	e.GET("/swagger/*any", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
