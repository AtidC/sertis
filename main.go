package main

import (
	"blog/app/auth"
	"blog/app/card"
	"blog/db"
	"blog/log"
	"os"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	config "github.com/spf13/viper"
	"github.com/tylerb/graceful"
)

func init() {
	// Get environment parameter :: dev, uat, prd
	env := os.Args[1]

	// Set log level
	config.Set("env", env)
	if env == "dev" || env == "uat" {
		log.SetLogLevel("Debug")
	} else {
		log.SetLogLevel("Info")
	}

	// API Start
	log.Info("Server start running on %s environment configuration", env)

	// Get env config
	config.SetConfigFile("config/" + env + ".yml")
	if err := config.ReadInConfig(); err != nil {
		log.Error("Fatal error env config (file): %s", err.Error())
	}
}

func main() {
	// Initial Grobal Database connection
	conn := db.InitPostgresPool()
	defer conn.Close()

	// Initial ECHO Framework
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	// Auth
	e.POST("/login", auth.Login)
	r := e.Group("/blog")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &auth.JwtCustomClaims{},
		SigningKey: []byte(config.GetString("auth.secretkey")),
	}))
	// Restricted Routes
	r.POST("/card", card.Create)
	r.GET("/card/list", card.List)
	r.PUT("/card", card.Update)
	r.DELETE("/card/:cardId", card.Delete)

	// Start Server, Graceful Shutdown with in 5sec.
	e.Server.Addr = ":" + config.GetString("service.port")
	graceful.ListenAndServe(e.Server, 5*time.Second)
}
