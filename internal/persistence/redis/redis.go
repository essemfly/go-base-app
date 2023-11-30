package redis

import (
	"context"
	"encoding/json"
	"errors"
	"essemfly/go_base_app/config"
	"essemfly/go_base_app/internal/domain"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

const LOG_ANALYTICS_KEY = "logResults"
const LOG_ANALYTICS_AGGREGATED_KEY = "logAggregatedResults"

type Redis struct {
	rdb *redis.Client
}

func NewRedisClient(cfg config.Config) (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURL,
		Password: "",
		DB:       0,
	})

	if rdb == nil {
		return nil, errors.New("redis client 생성 실패")
	}

	return &Redis{rdb: rdb}, nil
}

func (r *Redis) SaveLogAnalytics(score float64, jsonData []byte) error {
	ctx := context.Background()

	member := &redis.Z{
		Score:  score,
		Member: jsonData,
	}
	return r.rdb.ZAdd(ctx, LOG_ANALYTICS_KEY, member).Err()
}

func (r *Redis) GetLatestLogResult() ([]byte, error) {
	ctx := context.Background()

	results, err := r.rdb.ZRevRangeWithScores(ctx, LOG_ANALYTICS_AGGREGATED_KEY, 0, 0).Result()
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, errors.New("no data")
	}

	return []byte(results[0].Member.(string)), nil
}

func (r *Redis) SaveAggregateData() error {
	ctx := context.Background()

	oneHourAgo := time.Now().Add(-1 * time.Hour)

	results, err := r.rdb.ZRangeByScoreWithScores(ctx, LOG_ANALYTICS_KEY, &redis.ZRangeBy{
		Min: fmt.Sprintf("%d", oneHourAgo.Unix()),
		Max: fmt.Sprintf("%d", time.Now().Unix()),
	}).Result()
	if err != nil {
		panic(err)
	}

	if len(results) == 0 {
		return errors.New("no data to aggregate")
	}

	var totalNumLogs int
	var totalScore float64

	for _, result := range results {
		var logResult domain.LogAnalysisResult
		err := json.Unmarshal([]byte(result.Member.(string)), &logResult)
		if err != nil {
			panic(err)
		}

		if logResult.Topic == "logs" {
			totalNumLogs += logResult.NumLogs
			totalScore += logResult.AvgScore * float64(logResult.NumLogs)
		}
	}

	var avgScore float64
	if totalNumLogs > 0 {
		avgScore = totalScore / float64(totalNumLogs)
	}
	summary := domain.LogAnalysisResult{
		Topic:     "summary",
		AvgScore:  avgScore,
		NumLogs:   totalNumLogs,
		CreatedAt: time.Now(),
	}

	summaryJson, err := json.Marshal(summary)
	if err != nil {
		panic(err)
	}

	_, err = r.rdb.ZAdd(ctx, "LOG_ANALYTICS_AGGREGATED_KEY", &redis.Z{
		Score:  float64(summary.CreatedAt.Unix()),
		Member: summaryJson,
	}).Result()
	if err != nil {
		panic(err)
	}

	return nil
}
