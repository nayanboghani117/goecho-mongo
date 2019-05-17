package authapi

import (
	"time"

	"appScrip/app"
	"appScrip/errors"
	"appScrip/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/auth"
	"gopkg.in/mgo.v2/bson"
)

// Credential Structure
type (
	// userService specifies the interface for the users service needed by usersResource.
	userService interface {
		Get(rs app.RequestScope) ([]models.User, error)
		GetOne(rs app.RequestScope, id bson.ObjectId) (*models.User, error)
		Create(rs app.RequestScope, model *models.User) (*models.User, error)
		Update(rs app.RequestScope, id bson.ObjectId, model *models.User) (*models.User, error)
		Delete(rs app.RequestScope, id bson.ObjectId) (*models.User, error)
		Authenticate(rs app.RequestScope, email string, pass string) (*models.User, error)
	}

	// usersResource defines the handlers for the CRUD APIs.
	userResource struct {
		service userService
	}

	Credential struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

// ServerAuthResource sets up the routing of auth handlers.
func ServerAuthResource(rg *routing.RouteGroup, service userService) {
	r := &userResource{service}

	rg.Post("/auth", r.Auth)
}

// Auth function
func (r *userResource) Auth(c *routing.Context) error {
	var credential Credential
	if err := c.Read(&credential); err != nil {
		return errors.Unauthorized(err.Error())
	}

	response, err := r.service.Authenticate(app.GetRequestScope(c), credential.Email, credential.Password)

	if err != nil {
		return errors.Unauthorized("invalid credential")
	}

	token, err := auth.NewJWT(jwt.MapClaims{
		"id":   response.ID,
		"name": response.Name,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}, app.Config.JWTSigningKey)
	if err != nil {
		return errors.Unauthorized(err.Error())
	}

	return c.Write(map[string]string{
		"token": token,
	})
}
