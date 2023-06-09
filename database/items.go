package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	itemTable   string = "items"
	itemKeys    string = "item_name, description, maker_name, category"
	itemIdKey   string = "item_id"
	itemNameKey string = "item_name"
)

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Maker       string `json:"maker"`
	Category    string `json:"category"`
}

type itemAccess dbAccess

type IItemRepo interface {
	GetItems() ([]*Item, error)
	ItemById(id int) (*Item, error)
	CreateItem(new Item) (int, error)
	UpdateItem(update *Item, id int) error
	DeleteItem(id int) error
}

type itemRepo struct {
	acces itemAccess
	qb    QueryBuilder
}

func newItemStore(db *sql.DB) IItemRepo {
	q := NewQueryBuilder()
	return &itemRepo{
		acces: itemAccess{db: db},
		qb:    q,
	}
}

func (i *itemRepo) GetItems() ([]*Item, error) {
	query := i.qb.SelectAll(itemTable, itemNameKey)
	rows, err := i.acces.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := make([]*Item, 0)
	for rows.Next() {
		item := new(Item)
		if err = rows.Scan(&item.ID, &item.Name, &item.Description, &item.Maker, &item.Category); err != nil {

			return nil, err
		}
		items = append(items, item)

		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return items, nil
}

func (i *itemRepo) ItemById(id int) (*Item, error) {
	item := &Item{}
	query := i.qb.SelectByID(itemTable, itemIdKey, id)
	row := i.acces.db.QueryRow(query)

	switch err := row.Scan(&item.ID, &item.Name, &item.Description, &item.Maker, &item.Category); err {
	case sql.ErrNoRows:
		return nil, err
	case nil:
		return item, nil
	default:
		return nil, err
	}
}

func (i *itemRepo) CreateItem(new Item) (int, error) {
	id := 0
	values := fmt.Sprintf("(%v, %v, %v, %v)", new.Name, new.Description, new.Maker, new.Category)
	query := i.qb.Create(itemTable, itemKeys, values, itemIdKey)
	err := i.acces.db.QueryRow(query).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil

}

func (i *itemRepo) UpdateItem(update *Item, id int) error {

	values := fmt.Sprintf("(%v, %v, %v, %v) ", update.Name, update.Description, update.Maker, update.Category)
	query := i.qb.Update(itemTable, itemKeys, values, itemIdKey, id)

	_, err := i.acces.db.Exec(query)
	if err != nil {
		return err
	}

	return nil

}

func (i *itemRepo) DeleteItem(id int) error {
	query := i.qb.Delete(itemTable, itemIdKey, id)

	_, err := i.acces.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil

}
