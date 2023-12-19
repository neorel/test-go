package services

import (
	"context"
	"fmt"
	"encoding/json"
	"github.com/neorel/test-go/models"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var ctx context.Context

func InitRedis() bool {
	redisOpt, err := redis.ParseURL("redis://redis")
	if err != nil {
		fmt.Println("redis error", err)
		return false
	}

	client = redis.NewClient(redisOpt)

	ctx = context.Background()

	return true
}

func SendRaceLap(raceLap *models.RaceLap) bool{
	raceLapJson, err := json.Marshal(raceLap)
	if err != nil {
		fmt.Println(err)
		return false
	}

	err = client.Set(ctx, "raceLap", string(raceLapJson), 0).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
