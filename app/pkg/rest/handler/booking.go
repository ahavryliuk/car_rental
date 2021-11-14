package handler

import (
	"carRentalVivino/pkg/rest/request"
	"carRentalVivino/pkg/service"
	"carRentalVivino/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func NewBookingHandler(
	bookingManager *storage.BookingManager,
	carManager *storage.CarManager,
	carService *service.CarService,
	redisService *service.RedisService,
) *BookingHandler {
	return &BookingHandler{
		bookingManager: bookingManager,
		carManager:     carManager,
		carService:     carService,
		redisService:   redisService,
	}
}

type BookingHandler struct {
	bookingManager *storage.BookingManager
	carManager     *storage.CarManager
	carService     *service.CarService
	redisService   *service.RedisService
}

func (h *BookingHandler) Create(ctx *gin.Context) {
	var data request.BookingCreate

	if err := ctx.BindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})
		return
	}

	if err := h.carService.AssertDateRangeAvailability(data.CarUUID, data.DateFrom, data.DateTo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
		return
	}

	car, err := h.carManager.GetByUUID(data.CarUUID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{})
		return
	}

	booking := &storage.Booking{}
	booking.DateFrom = data.DateFrom
	booking.DateTo = data.DateTo
	booking.CarId = car.ID

	if err = h.bookingManager.Persist(booking); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{})
		return
	}

	if err = h.redisService.PublishBooking(booking); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"uuid":         booking.UUID.String(),
		"access_token": booking.AccessToken,
	})
}

func (h *BookingHandler) Get(ctx *gin.Context) {
	bookingUUID, err := uuid.Parse(ctx.Param("uuid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{})
		return
	}

	booking, err := h.bookingManager.GetByUUID(bookingUUID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, booking)
}

func (h *BookingHandler) Update(ctx *gin.Context) {
	var data request.BookingUpdate

	if err := ctx.BindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})
		return
	}

	accessToken := ctx.Param("accessToken")
	bookingUUID, err := uuid.Parse(ctx.Param("uuid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{})
		return
	}

	booking, err := h.bookingManager.GetByUUIDAndAccessToken(bookingUUID, accessToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{})
		return
	}

	if err = h.carService.AssertDateRangeAvailability(data.CarUUID, data.DateFrom, data.DateTo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
		return
	}

	car, err := h.carManager.GetByUUID(data.CarUUID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{})
		return
	}

	booking.DateFrom = data.DateFrom
	booking.DateTo = data.DateTo
	booking.CarId = car.ID

	if err = h.bookingManager.Persist(booking); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, booking)
}

func (h *BookingHandler) Delete(ctx *gin.Context) {
	accessToken := ctx.Param("accessToken")
	bookingUUID, err := uuid.Parse(ctx.Param("uuid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{})
		return
	}

	if err = h.bookingManager.RemoveByUUIDAndAccessToken(bookingUUID, accessToken); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
