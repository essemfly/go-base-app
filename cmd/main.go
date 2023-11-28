package main

import (
	"essemfly/go_base_app/config"
	"essemfly/go_base_app/setup"
	"fmt"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server on port:", cfg.Port)

	_, err = setup.InitializeServices(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%d", cfg.Port))
}
