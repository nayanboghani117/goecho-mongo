package main

import (

	"github.com/labstack/echo"
	"go-echo-mongo/go-echo-mongo/apis/handler"
	"go-echo-mongo/go-echo-mongo/apis/users"
	"go-echo-mongo/go-echo-mongo/mware"


	echoSwagger "github.com/swaggo/echo-swagger"
	//"go-echo-mongo/go-echo-mongo/apis/handler"
	//"go-echo-mongo/go-echo-mongo/apis/users"
	_ "go-echo-mongo/go-echo-mongo/docs"

	"go-echo-mongo/go-echo-mongo/config"

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

	config := config.LoadConfig()

	// Database
	db := mware.NewDB(config.App.DSN)
	db.Init()
	h := &handler.Handler{DB:db}
	e.POST("/users", users.Post(h))

	//h := &handler.Handler{DB: db}

	//
	//v1 := e.Group("/v1")
	//{
	//	user := v1.Group("/users")
	//	{
	//		user.GET("/:id", users.GetUserByID(h))
	//		user.GET("/", users.GetUsers(h))
	//		user.POST("/", users.CreateUser(h))
	//		user.POST("/login", users.SignIn(h))
	//		user.PUT("/:id", users.UpdateUser(h))
	//		user.DELETE("/:id", users.DeleteUser(h))
	//	}
	//
	//
	//	auth := v1.Group("/auth")
	//	{
	//		auth.GET("/userauth", users.Private, mware.IsLoggedIn)
	//		auth.GET("/admin", users.Private, mware.IsLoggedIn,mware.IsAdmin)
	//	}
	//}

	e.GET("/swagger/*any", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
