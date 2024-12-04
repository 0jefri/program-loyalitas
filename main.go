package main

import (
	"log"
	"loyalty-program-api/config"
	"loyalty-program-api/handlers"
	"loyalty-program-api/repositories"
	routers "loyalty-program-api/routes"
	"loyalty-program-api/services"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Initialize the database
	config.InitDB()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(config.DB)
	transactionRepo := repositories.NewTransactionRepository(config.DB)
	redemptionRepo := repositories.NewRedemptionRepository(config.DB)

	// Initialize services
	userService := services.NewUserService(userRepo)
	transactionService := services.NewTransactionService(transactionRepo, userRepo)
	redemptionService := services.NewRedemptionService(redemptionRepo, userRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)
	redemptionHandler := handlers.NewRedemptionHandler(redemptionService)

	// Setup handlers
	historyHandler := handlers.NewHistoryHandler(transactionService, redemptionService)

	// Setup router
	r := routers.SetupRouter(userHandler, transactionHandler, redemptionHandler, historyHandler)

	// Run the server
	r.Run(":8080")
}
