package main

import (
	"fmt"
	"log"
	"os"
	"test-kp-golang/src/database"
	authHandler "test-kp-golang/src/domain/auth/handler"
	productHandler "test-kp-golang/src/domain/product/handler"
	ProductRepository "test-kp-golang/src/domain/product/repository"
	productUseCase "test-kp-golang/src/domain/product/use-case"
	transactionHandler "test-kp-golang/src/domain/transaction/handler"
	transactionRepository "test-kp-golang/src/domain/transaction/repository"
	transactionUseCase "test-kp-golang/src/domain/transaction/use-case"
	userHandler "test-kp-golang/src/domain/user/handler"
	"test-kp-golang/src/domain/user/repository"
	usecase "test-kp-golang/src/domain/user/use-case"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		log.Fatal("DB_HOST is not set in the environment variables")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		log.Fatal("DB_USER is not set in the environment variables")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME is not set in the environment variables")
	}

	password := os.Getenv("DB_PASSWORD")
	// if password == "" {
	// 	log.Fatal("DB_PASSWORD is not set in the environment variables")
	// }

	port := os.Getenv("DB_PORT")
	if port == "" {
		log.Fatal("DB_PORT is not set in the environment variables")
	}

	fmt.Println(host, user, dbName, password, port)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)

	db, err := database.ConnectToDatabase(dsn)
	if err != nil {
		log.Fatal("Error while connecting to the database:", err)
	}

	userRepo := repository.NewUserRepository(db)
	productRepo := ProductRepository.NewProductRepository(db)
	transactionRepo := transactionRepository.NewTransactionRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepo)
	productUsecase := productUseCase.NewProductUsecase(*productRepo)
	transactionUseCase := transactionUseCase.NewTransactionUseCase(*transactionRepo)

	r := gin.Default()
	userHandler.NewUserHandler(r, userUsecase)
	authHandler.NewAuthHandler(r, userUsecase)
	productHandler.NewProductHandler(r, productUsecase)
	transactionHandler.NewTransactionHandler(r, transactionUseCase)

	r.Run(":8080")
}
