package main

import (
	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine) {

	// Home
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"title": "Home Page",
		})
	})

}
