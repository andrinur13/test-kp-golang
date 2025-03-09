package main

import (
	"test-kp-golang/src/user/handler"
	"test-kp-golang/src/user/repository"
	usecase "test-kp-golang/src/user/use-case"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userRepo := repository.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)

	handler.NewUserHandler(r, userUsecase)

	r.Run(":8080")
}
