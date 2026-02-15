package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ryoito218/blog-go-next.js/internal/domain/model"
	"github.com/ryoito218/blog-go-next.js/internal/infra/db"
)

func main() {
	conn, err := db.Open()
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	if err := conn.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.PostImage{},
	); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
