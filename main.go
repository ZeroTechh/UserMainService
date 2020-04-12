package main

import (
	"go.uber.org/zap"

	"github.com/ZeroTechh/VelocityCore/logger"
	proto "github.com/ZeroTechh/VelocityCore/proto/UserMainService"
	"github.com/ZeroTechh/VelocityCore/services"
	"github.com/ZeroTechh/VelocityCore/utils"
	"github.com/ZeroTechh/hades"

	"github.com/ZeroTechh/UserMainService/serviceHandler"
)

var (
	config = hades.GetConfig("main.yaml", []string{"config"})
	log    = logger.GetLogger(
		config.Map("service").Str("logFile"),
		config.Map("service").Bool("debug"),
	)
)

func main() {
	defer utils.HandlePanic(log)
	defer log.Info("Service Stopped")

	grpcServer, listner := utils.CreateGRPCServer(
		services.UserMainService,
		log,
	)

	serviceHandler := serviceHandler.New()

	proto.RegisterUserMainServer(grpcServer, serviceHandler)

	log.Info("Service Started")
	if err := grpcServer.Serve(*listner); err != nil {
		log.Fatal("Service Failed With Error", zap.Error(err))
	}
}
