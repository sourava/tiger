package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/sourava/tiger/app/config"
	"github.com/sourava/tiger/app/constants"
	"github.com/sourava/tiger/app/handlers"
	"github.com/sourava/tiger/app/middlewares"
	service2 "github.com/sourava/tiger/business/auth/service"
	models2 "github.com/sourava/tiger/business/tiger/models"
	"github.com/sourava/tiger/business/tiger/request"
	service3 "github.com/sourava/tiger/business/tiger/service"
	"github.com/sourava/tiger/business/user/models"
	"github.com/sourava/tiger/business/user/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func initRouter(db *gorm.DB, tigerSightingNotificationChannel chan<- *request.SendTigerSightingNotificationRequest) *gin.Engine {
	jwtSecret := os.Getenv("JWT_PRIVATE_KEY")

	userService := service.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	authService := service2.NewAuthService(db, jwtSecret)
	authHandler := handlers.NewAuthHandler(authService)

	tigerService := service3.NewTigerService(db, tigerSightingNotificationChannel)
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

func initDB(dbConfig config.DBConfig) *gorm.DB {
	dsn := fmt.Sprintf(constants.DbConnectionString, dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
		panic("failed to connect database")
	}
	log.Info("database connected successfully", db)

	return db
}

func initDBMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models2.Tiger{})
	db.AutoMigrate(&models2.TigerSighting{})
}

func initDBSeeds(appConfig config.AppConfig, db *gorm.DB) {
	record := db.Create(&models.User{
		Username: appConfig.Superuser.Username,
		Email:    appConfig.Superuser.Email,
		Password: appConfig.Superuser.HashedPassword,
	})
	if record.Error != nil {
		log.Error(record.Error)
	}
}

func main() {
	appConfig := config.AppConfig{
		ServicePort: os.Getenv(constants.ServicePort),
		DB: config.DBConfig{
			Host:     os.Getenv(constants.MysqlHost),
			Port:     os.Getenv(constants.MysqlPort),
			Username: os.Getenv(constants.MysqlUser),
			Password: os.Getenv(constants.MysqlPassword),
			Database: os.Getenv(constants.MysqlDb),
		},
		Superuser: config.Superuser{
			Username:       os.Getenv(constants.SuperuserUsername),
			Email:          os.Getenv(constants.SuperuserEmail),
			HashedPassword: os.Getenv(constants.SuperuserHashedPassword),
		},
	}

	tigerSightingNotificationChannel := make(chan *request.SendTigerSightingNotificationRequest, 1000)

	go func() {
		for {
			select {
			case message := <-tigerSightingNotificationChannel:
				fmt.Println("+++++++++++++++++++++message", message.Reporters[0].Email)
			}
		}
	}()

	db := initDB(appConfig.DB)
	initDBMigrations(db)
	initDBSeeds(appConfig, db)

	r := initRouter(db, tigerSightingNotificationChannel)
	r.Run(fmt.Sprintf(":%v", appConfig.ServicePort))
}
