package auth

import (
	"blog/log"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
	config "github.com/spf13/viper"
)

func Login(c echo.Context) error {
	startProcess := time.Now()
	requestId := uuid.New().String()
	log.StartAPI(requestId, c)

	// Get username & password
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Check username & password
	valid, err := checkPassWord(requestId, username, password)
	if err != nil {
		log.EndAPI(requestId, startProcess, http.StatusInternalServerError, err.Error())
		return c.JSON(http.StatusBadRequest, Response{
			Code:      "E000",
			Message:   fmt.Sprintf("Can't login, %s", err.Error()),
			RequestId: requestId,
		})
	}
	if !valid {
		err := errors.New("Invalid username or password")
		log.EndAPI(requestId, startProcess, http.StatusUnauthorized, err.Error())
		return c.JSON(http.StatusBadRequest, Response{
			Code:      "E001",
			Message:   err.Error(),
			RequestId: requestId,
		})
	}

	// Get user data
	user, err := selectUserData(requestId, username)
	if err != nil {
		log.EndAPI(requestId, startProcess, http.StatusInternalServerError, err.Error())
		return c.JSON(http.StatusBadRequest, Response{
			Code:      "E100",
			Message:   fmt.Sprintf("Can't login, %s", err.Error()),
			RequestId: requestId,
		})
	}

	// Set JWT Token
	claims := &JwtCustomClaims{
		ID:   user.ID,
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.GetString("auth.secretkey")))
	if err != nil {
		log.EndAPI(requestId, startProcess, http.StatusInternalServerError, err.Error())
		return c.JSON(http.StatusBadRequest, Response{
			Code:      "E100",
			Message:   fmt.Sprintf("Can't login, %s", err.Error()),
			RequestId: requestId,
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func TokenInfo(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims.ID
}
