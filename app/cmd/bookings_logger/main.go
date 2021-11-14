package main

import (
	"carRentalVivino/pkg/config"
	"carRentalVivino/pkg/storage"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"runtime"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	cfg := config.NewConfigFromFile("./config/default.ini")

	ctx := context.Background()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       0,
	})

	err := redisClient.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}

	log.Printf("=================================================================")
	log.Printf("Bookings Logger (runtime %s, numCPU=%d) starting...", runtime.Version(), runtime.NumCPU())
	log.Printf("=================================================================")

	topic := redisClient.Subscribe(ctx, cfg.App.BookingLogChannel)
	channel := topic.Channel()
	for msg := range channel {
		booking := &storage.Booking{}

		err := booking.UnmarshalBinary([]byte(msg.Payload))
		if err != nil {
			panic(err)
		}

		log.Printf("[BOOKING] %+v\n", booking)
	}
}
