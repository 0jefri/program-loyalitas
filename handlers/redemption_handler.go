package handlers

import (
	"loyalty-program-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RedemptionHandler struct {
	service services.RedemptionService
}

func NewRedemptionHandler(service services.RedemptionService) *RedemptionHandler {
	return &RedemptionHandler{service: service}
}

func (h *RedemptionHandler) Redeem(c *gin.Context) {
	var input struct {
		UserID uint   `json:"user_id" binding:"required"`
		Reward string `json:"reward" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	redemption, err := h.service.RedemptionPoints(input.UserID, input.Reward)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, redemption)
}
