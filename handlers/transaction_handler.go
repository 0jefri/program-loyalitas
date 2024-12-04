package handlers

import (
	"loyalty-program-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service services.TransactionService
}

func NewTransactionHandler(service services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) Record(c *gin.Context) {
	var input struct {
		UserID        uint    `json:"user_id" binding:"required"`
		TransactionID string  `json:"transaction_id" binding:"required"`
		Nominal       float64 `json:"nominal_transaksi" binding:"required"`
		PaymentMethod string  `json:"payment_method" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := h.service.RecordTransaction(input.UserID, input.TransactionID, input.Nominal, input.PaymentMethod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}
