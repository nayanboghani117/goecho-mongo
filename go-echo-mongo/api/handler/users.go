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

// createUser godoc
// @Summary Add a user
// @Description add by json user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body model.User true "Add user"
// @Success 200 {object} model.User
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /user/createuser [post]
func CreateUser(h *Handler) echo.HandlerFunc {
	return func(c echo.Context) error {

		u := &model.User{ID: bson.NewObjectId()}
		if err := c.Bind(u); err != nil {
			return err
		}
		if u.Email == "" || u.PassWord == "" {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
		}
		db := h.DB.Clone()
		defer db.Close()
		if err := db.DB("demo").C("user").Insert(u); err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, u)
	}
}

// updateUser godoc
// @Summary update a user
// @Description  update user by json
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /user/updateuser/{id} [put]
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

// deleteuserbyid godoc
// @Summary Delete a user
// @Description delete user by ID
// @Tags user
// @Accept  json,xml
// @Produce  json,xml
// @Param id path string true "User ID"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /user/deleteuser/{id} [delete]
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


// getuserbyid godoc
// @Summary Show a user
// @Description get user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /user/getuserbyid/{id} [get]
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

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = u.FirstName
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

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
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}