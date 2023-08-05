package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/sourava/tiger/app/handlers"
	service2 "github.com/sourava/tiger/business/auth/service"
	"github.com/sourava/tiger/business/user/models"
	"github.com/sourava/tiger/business/user/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
		panic("failed to connect database")
	}
	log.Info("database connected successfully", db)

	db.AutoMigrate(&models.User{})

	userService := service.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	authService := service2.NewAuthService(db)
	authHandler := handlers.NewAuthHandler(authService)

	r := gin.Default()
	r.POST("/user", userHandler.CreateUser)
	r.POST("/login", authHandler.Login)
	r.Run()
}
