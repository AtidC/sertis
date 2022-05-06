package auth

import "github.com/golang-jwt/jwt"

type Response struct {
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	RequestId string      `json:"requestId"`
	Result    interface{} `json:"result,omitempty"`
}

type JwtCustomClaims struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
