package main

import (
	"os"
	"gokit-practice/abilities"
	"github.com/joho/godotenv"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"net/http"
	"github.com/davecgh/go-spew/spew"
	//"github.com/satori/go.uuid"
	"github.com/go-kit/kit/log"
	_ "github.com/auth0/go-jwt-middleware"
	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/codegangsta/negroni"
	"gokit-practice/auth"
)

func main() {
	// DB
	env := GetEnv()
	db, _ := gorm.Open(
		"postgres",
		"host="+env.dbHost+" port="+env.dbPort+" user="+env.dbUser+" dbname="+env.dbName+" sslmode=disable",
	)
	defer db.Close()

	// logger
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	httpLogger := log.With(logger, "component", "http")

	// auth mw
	authMw := auth.NewMiddleware(auth.Env{
		Aud: env.authAud,
		Iss: env.authIss,
		JwksEndpoint: env.authJwksEndpoint,
	})

	// Services
	abilitiesService := abilities.NewService(db)
	abilitiesService = abilities.WithLogger(logger, abilitiesService)

	// Router
	r := abilities.MakeHTTPHandler(abilitiesService, httpLogger, authMw)

	// Start
	spew.Dump("Starting server")
	err := http.ListenAndServe(":"+env.httpPort, r)
	spew.Dump(err)
}

type Env struct {
	dbHost           string
	dbPort           string
	dbUser           string
	dbName           string
	httpPort         string
	authAud          string
	authIss          string
	authJwksEndpoint string
}

func GetEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		//log.("Error loading .env file")
	}

	return &Env{
		dbHost:           os.Getenv("DB_HOST"),
		dbPort:           os.Getenv("DB_PORT"),
		dbUser:           os.Getenv("DB_USER"),
		dbName:           os.Getenv("DB_NAME"),
		httpPort:         os.Getenv("HTTP_PORT"),
		authAud:          os.Getenv("AUTH_AUD"),
		authIss:          os.Getenv("AUTH_ISS"),
		authJwksEndpoint: os.Getenv("AUTH_JWKS_ENDPOINT"),
	}
}
