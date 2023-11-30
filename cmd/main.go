package main

import (
	"essemfly/go_base_app/config"
	"essemfly/go_base_app/internal/handler"
	"essemfly/go_base_app/internal/service"
	"fmt"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server on port:", cfg.Port)

	services, err := service.InitializeServices(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	handler.SetupRoutes(r, services)

	r.Run(fmt.Sprintf(":%d", cfg.Port))
}
