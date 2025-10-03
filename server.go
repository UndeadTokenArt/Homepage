package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func runServer() {
	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Templates from templates folder
	r.LoadHTMLGlob("templates/*")

	// Static files from static folder
	r.Static("/static", "static")

	// Routes
	registerRoutes(r)

	// Server Port set by environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Run server
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

func registerRoutes(r *gin.Engine) {
	// Home
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"title": "Home Page",
		})
	})

}
