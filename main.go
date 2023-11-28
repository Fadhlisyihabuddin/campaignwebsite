package main

import (
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
	userHandler := handler.NewUserHandler(userSevice)

	r := gin.Default()
	api := r.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)

	r.Run()
}
