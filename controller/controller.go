package controller

import (
	"ad-server/model"
	"ad-server/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func (ctl *Controller) Run(cfg *utils.Config) error {

	logrus.Infof("server start on port %s ", cfg.Port)
	return ctl.Gin.Run(":" + cfg.Port)
}

func NewDB(cfg *utils.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	db.AutoMigrate(&model.ClickEvent{}, &model.Ad{})
	return db

}
