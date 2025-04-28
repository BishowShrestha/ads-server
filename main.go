package main

import (
	"ad-server/controller"
	"ad-server/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := utils.LoadConfig()
	logrus.Info("successfully load configuration")
	ctl := controller.NewController(cfg)
	_ = ctl.Run(cfg)
}
