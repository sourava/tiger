package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/sourava/tiger/app/handlers"
	"github.com/sourava/tiger/app/middlewares"
	service2 "github.com/sourava/tiger/business/auth/service"
	models2 "github.com/sourava/tiger/business/tiger/models"
	service3 "github.com/sourava/tiger/business/tiger/service"
	"github.com/sourava/tiger/business/user/models"
	"github.com/sourava/tiger/business/user/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func initRouter(db *gorm.DB) *gin.Engine {
	jwtSecret := os.Getenv("JWT_PRIVATE_KEY")

	userService := service.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	authService := service2.NewAuthService(db, jwtSecret)
	authHandler := handlers.NewAuthHandler(authService)

	tigerService := service3.NewTigerService(db)
	tigerHandler := handlers.NewTigerHandler(tigerService)

	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/login", authHandler.Login)
		api.GET("/tiger", tigerHandler.ListAllTigers)
		api.GET("/tiger/:tigerID/sightings", tigerHandler.ListAllTigerSightings)
		secured := api.Use(middlewares.Auth(jwtSecret))
		{
			secured.POST("/user", userHandler.CreateUser)
			secured.POST("/tiger", tigerHandler.CreateTiger)
			secured.POST("/tiger/sighting", tigerHandler.CreateTigerSighting)
		}
	}
	return router
}

func main() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
		panic("failed to connect database")
	}
	log.Info("database connected successfully", db)

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models2.Tiger{})
	db.AutoMigrate(&models2.TigerSighting{})
	db.Create(&models.User{Username: "user1", Email: "user1@email.com", Password: "$2a$04$npZR8DN1y2I0VNRrrPG6XOk.C2lfQLzCOhK5T9lR40oQuecSEHkhm"})

	r := initRouter(db)
	r.Run()
}
