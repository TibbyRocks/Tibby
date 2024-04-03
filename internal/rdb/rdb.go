package rdb

import (
	"context"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tibbyrocks/tibby/internal/utils"
)

var (
	botDB     redis.Client
	namespace string
	Log       = utils.Log
	ctx       = context.Background()
)

func init() {
	botDB = *redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Username: os.Getenv("REDIS_USER_NAMESPACE"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       utils.AtoiButGood(os.Getenv("REDIS_DATABASE")),
	})
	namespace = os.Getenv("REDIS_USER_NAMESPACE") + ":"
}

func TestRedis() {
	n := strconv.Itoa(int(rand.Int31()))

	Log.Debug("Testing Redis connection")
	Log.Debug("Writing test key to Redis", "Set", n)

	generalRedisErr := `Redis check failed (run with --debug)`

	err := Set("TestKey", n, 0)
	if err != nil {
		Log.Error(generalRedisErr)
		Log.Debug("Couldn't write test key", "Err", err.Error())
		os.Exit(3)
	}
	result, err := Get("TestKey")
	if err != nil {
		Log.Error(generalRedisErr)
		Log.Debug("Couldn't read test key", "Err", err.Error())
		os.Exit(3)
	} else {
		Log.Debug("Got test key back from Redis", "Got", result)
	}

	if result != n {
		Log.Error(generalRedisErr)
		Log.Debug("TestKey didn't match", "Expected", n, "Got", result)
		os.Exit(3)
	} else {

		Log.Debug("Redis test successful :)")
	}

}

func Get(key string) (string, error) {
	result, err := botDB.Get(ctx, namespace+key).Result()
	if err != nil {
		return "", err
	} else {
		return result, nil
	}
}

func Set(key string, value string, expiration time.Duration) error {
	err := botDB.Set(ctx, namespace+key, value, expiration).Err()
	if err != nil {
		return err
	} else {
		return nil
	}
}
