package main

import (
	"log"
	"net/http"
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
		// Load config data
		homepage, err := ParseConfigFile("config.json")
		if err != nil {
			log.Printf("Error loading config: %v", err)
			c.HTML(500, "index.tmpl", gin.H{
				"error": "Failed to load configuration",
			})
			return
		}

		// Render template with config data
		c.HTML(200, "index.tmpl", homepage)
	})

	//GM-tools rerouted to old website for now
	r.GET("/gm", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://dmtools.undeadtoken.com")
	})

	// RASA rerouted to old website for now
	r.GET("/rasa", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://undeadtoken.com/projects/RASA.html")
	})

}
