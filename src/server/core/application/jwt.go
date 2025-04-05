package application

import (
	"fmt"
	"template/core/log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	jwt.RegisteredClaims
}

var Config = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(JwtCustomClaims)
	},
	SigningKey: []byte("secret"),
}

func GenerateJwtToken(username string, email string, role int) (token string, err error) {
	claims := &JwtCustomClaims{
		email,
		username,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte("secret"))
	if err != nil {
		log.Errorf("[GenerateJwtToken]: %s", err)
		return
	}

	return
}

func ProtectedWithJwtHandler(r *Router, handler HandlerRouterFunc) func(c echo.Context) (err error) {
	return func(c echo.Context) (err error) {
		username, err := retrieveJwtInformation(c)
		if err != nil {
			return
		}
		log.Debugf("[ProtectedWithJwtHandler]: Request from: %s", username)
		return HandleRequest(r, c, handler)
	}
}

func retrieveJwtInformation(c echo.Context) (username string, err error) {
	user := c.Get("user")
	if user == nil {
		err = fmt.Errorf("jwt token is not present")
		return
	}

	token, ok := user.(*jwt.Token)
	if !ok {
		err = fmt.Errorf("jwt is not valid")
		return
	}

	claims := token.Claims.(*JwtCustomClaims)
	username = claims.Username
	return
}
