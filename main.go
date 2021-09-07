package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	projectId = "test-admin-firebase-f196c"
)

type notif struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

//var posts []Post

func init() {
	ctx := context.Background()
	config := &firebase.Config{
		DatabaseURL: "https://test-admin-firebase-f196c-default-rtdb.europe-west1.firebasedatabase.app",
	}
	app, err := firebase.NewApp(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}
	var n notif
	err = client.NewRef("notificationTemplates/long/1").Get(ctx, &n)
	if err != nil {
		log.Panic("after new ref: ", err)
	}
	log.Println("body:", n.Body)
}

func main() {
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
