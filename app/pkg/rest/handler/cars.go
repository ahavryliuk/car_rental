package handler

import (
	"carRentalVivino/pkg/rest/request"
	"carRentalVivino/pkg/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewCarHandler(
	carManager *storage.CarManager,
) *CarHandler {
	return &CarHandler{
		carManager: carManager,
	}
}

type CarHandler struct {
	carManager *storage.CarManager
}

func (h *CarHandler) List(ctx *gin.Context) {
	var data request.CarList

	if err := ctx.BindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})

		return
	}

	cars, err := h.carManager.GetAvailableByDateRange(data.DateFrom, data.DateTo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, cars)
}
