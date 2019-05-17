package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"go-echo-mongo/store/app"
	"go-echo-mongo/store/errors"
	"gopkg.in/mgo.v2"
	"net/http"
)


func main() {

	//load config path
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	// load error messages
	if err := errors.LoadMessages(app.Config.ErrorFile); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	logger := logrus.New()
	db, err := mgo.Dial(app.Config.DSN) //dbx.MustOpen("postgres", app.Config.DSN)
	if err != nil {
		panic(err)
	}else {
		fmt.Println(db)
	}
	e := echo.New()
	address := fmt.Sprintf(":%v", app.Config.ServerPort)
	logger.Infof("server %v is started at %v", app.Version, address)
	panic(http.ListenAndServe(address, nil))

	e.Logger.Fatal(e.Start(":8080"))

}
