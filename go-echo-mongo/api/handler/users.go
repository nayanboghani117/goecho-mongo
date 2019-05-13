package handler

import (

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go-echo-mongo/model"
	"net/http"
)

func CreateUser(h *Handler) echo.HandlerFunc {
	return func(c echo.Context) error {

		u := &model.User{ID: bson.NewObjectId()}
		if err := c.Bind(u); err != nil {
			return err
		}
		if u.Email == "" || u.PassWord == "" {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
		}
		// Save user
		db := h.DB.Clone()
		defer db.Close()
		if err := db.DB("demo").C("user").Insert(u); err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, u)
	}
}

func UpdateUser(h *Handler) echo.HandlerFunc {

	return func(c echo.Context) error {
		id := c.Param("id")
		if id == "" || !bson.IsObjectIdHex(id) {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid ID"}
		}
		u := &model.User{}
		if err := c.Bind(u); err != nil {
			return err
		}
		db := h.DB.Clone()
		defer db.Close()
		u.ID = bson.ObjectIdHex(id)
		if err := db.DB("demo").C("user").UpdateId(u.ID,u); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, u)
	}

}

func DeleteUser(h *Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if id == "" || !bson.IsObjectIdHex(id) {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid ID"}
		}
		db := h.DB.Clone()
		defer db.Close()
		u := &model.User{}
		u.ID = bson.ObjectIdHex(id)
		if err:= db.DB("demo").C("user").RemoveId(u.ID); err != nil {
			return err
		}
		return  c.JSON(http.StatusOK,id)
	}
}

func GetUsers(h *Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		db := h.DB.Clone()
		defer db.Close()
		u := &model.Users{}
		if err:= db.DB("demo").C("user").Find(bson.M{}).All(&u.Users); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
	}

}


func GetUserByID(h *Handler) echo.HandlerFunc{
	return func(c echo.Context) error {
		id := c.Param("id")
		u := &model.User{}
		if id == "" || !bson.IsObjectIdHex(id) {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid ID"}
		}
		db := h.DB.Clone()
		defer db.Close()
		u.ID = bson.ObjectIdHex(id)
		if err:= db.DB("demo").C("user").FindId(u.ID).One(&u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated,u)
	}
}

func SignIn(h *Handler) echo.HandlerFunc {

	return func(c echo.Context) (err error) {

		u := new(model.User)
		if err = c.Bind(u); err != nil {
			return
		}

		db := h.DB.Clone()
		defer db.Close()
		if err = db.DB("demo").C("user").
			Find(bson.M{"email": u.Email, "password": u.PassWord}).One(u); err != nil {
			if err == mgo.ErrNotFound {
				return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
			}
			return
		}

		// JWT token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = u.ID
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response
		u.Token, err = token.SignedString([]byte(Key))
		if err != nil {
			return err
		}

		u.PassWord = ""
		return c.JSON(http.StatusOK, u)
	}
}

func Private(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["id"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}