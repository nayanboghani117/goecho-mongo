package main

import (
	"fmt"
	"github.com/labstack/echo"
	"goechoProject/config"
	"gopkg.in/mgo.v2"
	"github.com/Sirupsen/logrus"
)

type (

	Handler struct {
		DB *mgo.Session
	}
)

func main() {
	e := echo.New()
	conf := config.LoadConfig()
	db, err := mgo.Dial(conf.App.DSN) //dbx.MustOpen("postgres", app.Config.DSN)
	if err != nil {
		panic(err)
	}
	h := &Handler{DB: db}
	logger := logrus.New()
	fmt.Println(h.DB.DatabaseNames())
}