package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ichtrojan/go-todo/routes"
	"github.com/ichtrojan/thoth"
)

func main() {
	logger, _ := thoth.Init("log")

	// if err := godotenv.Load(); err != nil {
	// 	logger.Log(errors.New("no .env file found"))
	// 	log.Fatal("No .env file found")
	// }

	//port, exist := os.LookupEnv("PORT")

	// if !exist {
	// 	logger.Log(errors.New("PORT not set in .env"))
	// 	log.Fatal("PORT not set in .env")
	// }

	// Azure App Service sets the port as an Environment Variable
	// This can be random, so needs to be loaded at startup
	port := os.Getenv("HTTP_PLATFORM_PORT")

	// default back to 8080 for local dev
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, routes.Init())

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	}
}
