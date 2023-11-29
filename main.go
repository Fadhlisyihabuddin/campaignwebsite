package main

import (
	"campaignwebsite/auth"
	"campaignwebsite/handler"
	"campaignwebsite/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=Zenab123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	userRepository := user.NewRepository(db)
	userSevice := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userSevice, authService)

	r := gin.Default()
	api := r.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/email_checker", userHandler.CheckEmailAvailable)
	api.POST("/avatars", userHandler.UploadAvatar)

	r.Run()
}
