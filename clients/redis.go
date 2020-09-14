package clients

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/voluntariado-ucc-ing/volunteer_api/config"
	"github.com/voluntariado-ucc-ing/volunteer_api/domain/apierrors"
	"time"
)

var ctx = context.Background()

var (
	Client *redis.Client
)

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.GetRedisHost(), config.GetRedisPort()),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	res, err := Client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Response from redis: ", res)
}

func getKey(volunteerId int64) string {
	return fmt.Sprintf("vol_%d", volunteerId)
}

func SetVolunteerAlreadyLoggedIn(volunteerId int64) apierrors.ApiError{
	err := Client.Set(ctx, getKey(volunteerId), time.Now().String(), 0)
	if err != nil {
		return apierrors.NewInternalServerApiError("error while setting key value in redis db", err.Err())
	}
	return nil
}

func GetVolunteerAlreadyLoggedIn(volunteerId int64) apierrors.ApiError {
	res, err := Client.Get(ctx, getKey(volunteerId)).Result()
	if err != nil {
		fmt.Println("User never logged in", err)
		return apierrors.NewInternalServerApiError("error getting key in redis db", err)
	}
	fmt.Println("User already logged in", res)
	return nil
}