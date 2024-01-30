package cache

import "github.com/redis/go-redis/v9"

func Connect() *redis.Client {
	opt, err := redis.ParseURL("redis://localhost:6379/task-management")
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opt)
}
