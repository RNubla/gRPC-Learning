package main

import (
	"fmt"
	"log"
	"os"

	"gRPC-Learning/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	// case "postgres":
	// 	// URL for PostgreSQL connection.
	// 	url = fmt.Sprintf(
	// 		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	// 		os.Getenv("DB_HOST"),
	// 		os.Getenv("DB_PORT"),
	// 		os.Getenv("DB_USER"),
	// 		os.Getenv("DB_PASSWORD"),
	// 		os.Getenv("DB_NAME"),
	// 		os.Getenv("DB_SSL_MODE"),
	// 	)
	// case "mysql":
	// 	// URL for Mysql connection.
	// 	url = fmt.Sprintf(
	// 		"%s:%s@tcp(%s:%s)/%s",
	// 		os.Getenv("DB_USER"),
	// 		os.Getenv("DB_PASSWORD"),
	// 		os.Getenv("DB_HOST"),
	// 		os.Getenv("DB_PORT"),
	// 		os.Getenv("DB_NAME"),
	// 	)
	// case "redis":
	// 	// URL for Redis connection.
	// 	url = fmt.Sprintf(
	// 		"%s:%s",
	// 		os.Getenv("REDIS_HOST"),
	// 		os.Getenv("REDIS_PORT"),
	// 	)
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {

	app := fiber.New()
	// app.Use(cors.New(cors.Config{
	// 	AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
	// 	AllowOrigins:     "*",
	// 	AllowCredentials: true,
	// 	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	// }))
	router.SetupRoutes(app)

	// serverHost := goDotEnvVariable("SERVER_HOST")
	port := goDotEnvVariable("SERVER_PORT")
	// url := fmt.Sprint(serverHost, ":", port)
	url := fmt.Sprint(":", port)

	log.Fatal(app.Listen(url))
	// log.Fatal(app.Listen("localhost:5000"))

}
