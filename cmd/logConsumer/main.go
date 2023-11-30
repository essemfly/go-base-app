package main

import (
	"essemfly/go_base_app/config"
	"essemfly/go_base_app/internal/persistence"
	"essemfly/go_base_app/internal/service"
	"log"
	"time"
)

const CONSUME_TIME_CYCLE = 2

func main() {
	cfg, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal(err)
	}

	persistences, err := persistence.InitializePersistence(cfg)
	if err != nil {
		log.Fatal(err)
	}

	services, err := service.InitializeServices(persistences)
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	for {
		log.Println("LogConsumer is running...")

		analyticResult := services.LogService.ConsumeLogs()
		err := services.LogService.SaveLogAnalytics(analyticResult)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(CONSUME_TIME_CYCLE * time.Second)
		i++
	}
}
