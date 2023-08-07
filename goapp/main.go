package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func hello(rw http.ResponseWriter, r *http.Request) {
// 	type helloMsg struct {
// 		Message string `json:message`
// 	}
// 	message := helloMsg{Message: "hello from go app"}

// 	jsonRes, jsonErr := json.Marshal(message)
// 	if jsonErr != nil {
// 		fmt.Println("Unable to encode JSON")
// 	}

// 	rw.Header().Set("Content-Type", "application/json")
// 	rw.WriteHeader(http.StatusOK)
// 	rw.Write(jsonRes)

// 	// json.NewEncoder(w).Encode(map[string]string{"message": "hello from go app"})
// }

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello from GO!",
	})
}

func main() {
	// r := mux.NewRouter()
	// r.HandleFunc("/", hello)
	// log.Fatal(http.ListenAndServe("localhost:70", handlers.LoggingHandler(os.Stdout, r)))

	r := gin.Default()
	r.GET("/", hello)
	r.Run(":70")
}

// import (
// 	"gapp/router"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/cors"
// )

// func main() {
// 	app := fiber.New()
// 	router.SetupRoutes(app)
// 	app.Use(cors.New(cors.Config{
// 		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
// 		AllowOrigins:     "*",
// 		AllowCredentials: true,
// 		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
// 	}))
// 	app.Listen("localhost:70")
// }
