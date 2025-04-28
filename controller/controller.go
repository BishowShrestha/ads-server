package controller

import (
	"ad-server/model"
	"ad-server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Controller struct {
	DB     *gorm.DB
	Gin    *gin.Engine
	Logger *zap.Logger
}

func NewController(cfg *utils.Config) *Controller {
	ctl := &Controller{}
	ctl.DB = NewDB(cfg)
	ctl.Gin = gin.Default()
	ctl.Logger = utils.NewLogger()
	ctl.Routes()
	return ctl
}

func (ctl *Controller) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logrus.Infof("server start on port %s ", port)
	return ctl.Gin.Run(":" + port)
}

func NewDB(cfg *utils.Config) *gorm.DB {
	fmt.Println("ddd", cfg.DatabaseURL)
	//dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	db.AutoMigrate(&model.ClickEvent{}, &model.Ad{})
	return db

}
