package main

import (
	"fmt"
	"log"
	"os"
	"docker-multi-stage-build-101/route"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// PORT := ":9000"

func SetupRouter() *gin.Engine {
	// Disable Console Color, you don't need console color when writing the logs to file.
	// gin.DisableConsoleColor()

	// Logging to a file
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// Default With the Logger and Recovery middleware already attached
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	router.Use(cors.New(config))

	// Simple group: api
	api := router.Group("/api")
	{
		// /api/user/naoto/kick
		api.GET("/user/:name/*action", route.GetAction)
		// /api/welcome?firstname=Naoto&lastname=Imamachi
		api.GET("/welcome", route.GetWelcome)
	}
	return router
}

func main() {
	router := SetupRouter()
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":9000"
		log.Printf("Defaulting to port %s", port)
	}
	router.Run(port) // listen and serve on 0.0.0.0:8080
}
