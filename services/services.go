package services

import (
	"sandbox-api/config"
	"sandbox-api/database"
	"sandbox-api/logs"
)

var logger *logs.AppLogger

type Services struct {
	IItemService
	IMakerService
	IUserService
}

func NewAppService(access *database.RepoStore, cfg *config.Config) *Services {
	logger = logs.NewAppLogger()
	return &Services{
		IItemService:  NewItemService(access.Items),
		IMakerService: NewMakerService(access.Makers),
		IUserService:  NewUserService(access.Users, cfg.Pepper, cfg.HMACKey),
	}
}
