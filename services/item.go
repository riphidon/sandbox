package services

import (
	"sandbox-api/database"
	"strconv"
)

type IItemService interface {
	Get() ([]*database.Item, error)
	ById(id int) (*database.Item, error)
	Create(new database.Item) (int, error)
	NewItem() *database.Item
	Update(update *database.Item, id int) error
	Delete(id int) error
}
type itemService struct {
	access database.IItemRepo
}

func NewItemService(repo database.IItemRepo) IItemService {
	return &itemService{
		access: repo,
	}
}

func (s *itemService) Get() ([]*database.Item, error) {
	items, err := s.access.GetItems()
	if err != nil {
		logger.Debugf(" [DB] %v", err.Error())
		return nil, err
	}
	return items, nil
}

func (s *itemService) ById(id int) (*database.Item, error) {
	item, err := s.access.ItemById(id)
	if err != nil {
		logger.Debugf(" [DB] %v", err.Error())
		return nil, err
	}
	return item, nil
}

func StrTonInt(str string) (int, error) {
	n, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (s *itemService) Create(new database.Item) (int, error) {

	id, err := s.access.CreateItem(new)
	if err != nil {
		logger.Debugf(" [DB] %v", err.Error())
		return 0, err
	}

	return id, nil
}

func (s *itemService) NewItem() *database.Item {
	return &database.Item{}
}

func (s *itemService) Update(update *database.Item, id int) error {
	err := s.access.UpdateItem(update, id)
	if err != nil {
		logger.Debugf(" [DB] %v", err.Error())
		return err
	}
	return nil
}

func (s *itemService) Delete(id int) error {
	err := s.access.DeleteItem(id)
	if err != nil {
		logger.Debugf(" [DB] %v", err.Error())
		return err
	}
	return nil
}
