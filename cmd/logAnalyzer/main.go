package main

import (
	"essemfly/go_base_app/config"
	"essemfly/go_base_app/internal/persistence"
	"log"
	"time"
)

const ANALYZE_REPORT_TIME_CYCLE = 5

func main() {
	cfg, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal(err)
	}

	persistences, err := persistence.InitializePersistence(cfg)
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("LogAnalyzer is running...")
		persistences.RedisClient.SaveAggregateData()
		time.Sleep(ANALYZE_REPORT_TIME_CYCLE * time.Second)
	}
}
