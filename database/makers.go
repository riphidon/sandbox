package database

import (
	"database/sql"
	"fmt"
)

const (
	MAKER_DB     string = "database/makers.json"
	makerTable   string = "makers"
	makerKeys    string = "(maker_name)"
	makerIdKey   string = "id_maker"
	makerNameKey string = "maker_name"
)

type Maker struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type makerAccess dbAccess

type IMakerRepo interface {
	GetMakers() ([]*Maker, error)
	MakerById(id int) (*Maker, error)
	CreateMaker(new Maker) (int, error)
	UpdateMaker(update Maker, id int) error
	DeleteMaker(id int) error
	NewMaker() Maker
}

type makerRepo struct {
	access makerAccess
	qb     QueryBuilder
}

func newMakerStore(db *sql.DB) IMakerRepo {
	q := NewQueryBuilder()
	return &makerRepo{
		access: makerAccess{db: db},
		qb:     q,
	}
}

func (m makerRepo) NewMaker() Maker {
	return Maker{}
}

// Get implements MakerRepoAccess
func (m *makerRepo) GetMakers() ([]*Maker, error) {

	query := m.qb.SelectAll(makerTable, makerNameKey)
	rows, err := m.access.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	makers := make([]*Maker, 0)
	for rows.Next() {
		maker := new(Maker)
		if err = rows.Scan(&maker.ID, &maker.Name); err != nil {
			return nil, err
		}
		makers = append(makers, maker)

		if err = rows.Err(); err != nil {

			return nil, err
		}
	}
	return makers, nil

}

// CreateMaker implements IMakerRepo
func (m *makerRepo) CreateMaker(new Maker) (int, error) {
	values := fmt.Sprintf("(%v)", new.Name)
	query := m.qb.Create(makerTable, makerKeys, values, makerIdKey)
	if err := m.access.db.QueryRow(query, new.Name).Scan(&new.ID); err != nil {
		return 0, err
	}

	return new.ID, nil
}

// DeleteMaker implements IMakerRepo
func (m *makerRepo) DeleteMaker(id int) error {
	query := m.qb.Delete(makerTable, makerIdKey, id)
	_, err := m.access.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// MakerById implements IMakerRepo
func (m *makerRepo) MakerById(id int) (*Maker, error) {
	maker := &Maker{}
	query := m.qb.SelectByID(makerTable, makerIdKey, id)
	row := m.access.db.QueryRow(query)

	switch err := row.Scan(maker); err {
	case sql.ErrNoRows:
		return nil, err
	case nil:
		return maker, nil
	default:
		return nil, err

	}
}

// UpdateMaker implements IMakerRepo
func (m *makerRepo) UpdateMaker(update Maker, id int) error {
	values := fmt.Sprintf("(%v)", update.Name)
	query := m.qb.Update(makerTable, makerKeys, values, makerIdKey, id)

	_, err := m.access.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil

}
