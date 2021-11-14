package main

import (
	"carRentalVivino/pkg/config"
	"carRentalVivino/pkg/rest"
	"carRentalVivino/pkg/service"
	"carRentalVivino/pkg/storage"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"runtime"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	gin.SetMode(gin.ReleaseMode)

	cfg := config.NewConfigFromFile("./config/default.ini")
	api := gin.Default()
	db := storage.BuildDb(cfg.MySQL)
	ctx := context.Background()

	bookingManager := storage.NewBookingManager(db)
	carManager := storage.NewCarManager(db)
	bookingService := service.NewCarService(bookingManager, carManager)
	redisService := service.NewRedisService(ctx, &cfg)

	server := rest.NewServer(
		db,
		api,
		bookingManager,
		carManager,
		bookingService,
		redisService,
	)

	server.BuildRoutes()

	log.Printf("=================================================================")
	log.Printf("Rental API (runtime %s, numCPU=%d) starting...", runtime.Version(), runtime.NumCPU())
	log.Printf("=================================================================")

	err := server.Run()
	if err != nil {
		log.Fatal("Run error: ", err)
	}
}
