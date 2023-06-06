package services

import (
	"sandbox-api/database"
	"sandbox-api/logs"
)

var logger *logs.AppLogger

type Services struct {
	IItemService
	IMakerService
}

func NewAppService(access *database.RepoStore) *Services {
	logger = logs.NewAppLogger()
	return &Services{
		IItemService:  NewItemService(access.Items),
		IMakerService: NewMakerService(access.Makers),
	}
}
