package middleware

import (
	"appScrip/app"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-ozzo/ozzo-routing"
)

// JWTHandler function
func JWTHandler(c *routing.Context, j *jwt.Token) error {
	userID := j.Claims.(jwt.MapClaims)["id"].(string)
	app.GetRequestScope(c).SetUserID(userID)
	userName := j.Claims.(jwt.MapClaims)["name"].(string)
	app.GetRequestScope(c).SetUserName(userName)
	return nil
}
