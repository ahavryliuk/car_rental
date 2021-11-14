package rest

import (
	"carRentalVivino/pkg/rest/handler"
	"carRentalVivino/pkg/service"
	"carRentalVivino/pkg/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewServer(
	db *gorm.DB,
	api *gin.Engine,
	bookingManager *storage.BookingManager,
	carManager *storage.CarManager,
	bookingService *service.CarService,
	redisService *service.RedisService,
) *Server {
	return &Server{
		db:             db,
		api:            api,
		bookingManager: bookingManager,
		carManager:     carManager,
		bookingService: bookingService,
		redisService:   redisService,
	}
}

type Server struct {
	db             *gorm.DB
	api            *gin.Engine
	bookingManager *storage.BookingManager
	carManager     *storage.CarManager
	bookingService *service.CarService
	redisService   *service.RedisService
}

func (s *Server) BuildRoutes() {
	bookingHandler := handler.NewBookingHandler(s.bookingManager, s.carManager, s.bookingService, s.redisService)
	carsHandler := handler.NewCarHandler(s.carManager)

	v1 := s.api.Group("/api/v1")
	{
		cars := v1.Group("/cars")
		{
			cars.GET("", carsHandler.List)
		}

		booking := v1.Group("/bookings")
		{
			booking.POST("", bookingHandler.Create)
			booking.GET("/:uuid", bookingHandler.Get)
			booking.PUT("/:uuid/:accessToken", bookingHandler.Update)
			booking.DELETE("/:uuid/:accessToken", bookingHandler.Delete)
		}
	}
}

func (s Server) Run() error {
	return s.api.Run()
}
