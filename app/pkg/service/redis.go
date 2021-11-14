package service

import (
	"carRentalVivino/pkg/config"
	"carRentalVivino/pkg/storage"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func NewRedisService(ctx context.Context, cfg *config.Config) *RedisService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       0,
	})

	return &RedisService{
		ctx:    ctx,
		client: redisClient,
		cfg:    cfg,
	}
}

type RedisService struct {
	ctx    context.Context
	client *redis.Client
	cfg    *config.Config
}

func (s *RedisService) PublishBooking(booking *storage.Booking) error {
	return s.client.Publish(s.ctx, s.cfg.App.BookingLogChannel, booking).Err()
}
