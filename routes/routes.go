package routers

import (
	"loyalty-program-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handlers.UserHandler, transactionHandler *handlers.TransactionHandler, redemptionHandler *handlers.RedemptionHandler, historyHandler *handlers.HistoryHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/users/register", userHandler.Register)
	r.POST("/transactions", transactionHandler.Record)
	r.POST("/redemptions", redemptionHandler.Redeem)
	r.GET("/users/:user_id/points", userHandler.GetPoints)
	r.GET("/users/:user_id/history", historyHandler.GetHistory)

	return r
}
