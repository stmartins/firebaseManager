package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	fmt.Println("demarrage du init")
	router := gin.Default()
	router.LoadHTMLGlob("index.html")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Vincent",
		})
	})
	router.Run(":8080")
}

func Start() {
	fmt.Println("demarrage du server")
}
