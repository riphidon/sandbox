package services

import (
	"sandbox-api/database"
)

const ()

type IMakerService interface {
	Get() ([]*database.Maker, error)
	Create(database.Maker) (int, error)
	Update(update database.Maker, id int) error
	NewMaker() *database.Maker
	Delete(id int) error
}

type makerService struct {
	access database.IMakerRepo
}

func NewMakerService(repo database.IMakerRepo) IMakerService {
	return &makerService{
		access: repo,
	}
}

func (s *makerService) NewMaker() *database.Maker {
	return &database.Maker{}
}

func (s *makerService) Get() ([]*database.Maker, error) {
	makers, err := s.access.GetMakers()
	if err != nil {
		logger.Debugf(" [DB] %v", err.Error())
		return nil, err
	}

	return makers, nil

}

// Create implements MakerService
func (s *makerService) Create(new database.Maker) (int, error) {
	id, err := s.access.CreateMaker(new)
	if err != nil {
		logger.Debugf(" [DB] %v", err.Error())
		return 0, err
	}
	return id, nil
}

// Delete implements MakerService
func (s *makerService) Delete(id int) error {
	if err := s.access.DeleteMaker(id); err != nil {
		logger.Debugf(" [DB] %v", err.Error())
		return err
	}
	return nil

}

// Update implements MakerService
func (s *makerService) Update(update database.Maker, id int) error {
	if err := s.access.UpdateMaker(update, id); err != nil {
		logger.Debugf(" [DB] %v", err.Error())
		return err
	}
	return nil

}
