// handlers/history_handler.go
package handlers

import (
	"loyalty-program-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HistoryHandler struct {
	TransactionService services.TransactionService
	RedemptionService  services.RedemptionService
}

func NewHistoryHandler(transactionService services.TransactionService, redemptionService services.RedemptionService) *HistoryHandler {
	return &HistoryHandler{
		TransactionService: transactionService,
		RedemptionService:  redemptionService,
	}
}

func (h *HistoryHandler) GetHistory(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Mendapatkan riwayat transaksi
	transactions, err := h.TransactionService.GetTransactionHistory(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Mendapatkan riwayat penukaran
	redemptions, err := h.RedemptionService.GetRedemptionHistory(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":      userID,
		"transactions": transactions,
		"redemptions":  redemptions,
	})
}
