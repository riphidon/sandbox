package database

import "fmt"

const (
	SELECT_FROM string = "SELECT %v FROM %v "
	DELETE_FROM string = "DELETE FROM %v "
	INSERT_INTO string = "INSERT INTO %v (%v) "
	UPDATE_SET  string = "UPDATE %v SET  (%v) = (%v)"
	WHERE_EQUAL string = "WHERE %v = %v "
	ORDER_BY    string = "ORDER BY %v %v"
	VALUES      string = "VALUES (%v) "
	RETURNING   string = "RETURNING %v"
	ASC         string = "ASC"
	DESC        string = "DESC"
	all         string = "*"
)

type QueryBuilder interface {
	SelectAll(table, orderBy string) string
	SelectByID(table, idColumn string, id int) string
	Delete(table, idColumn string, id int) string
	Update(table, keys, values, idColumn string, id int) string
	Create(table, keys, values, returned string) string
}

type queryHandler struct {
	QueryBuilder
}

func NewQueryBuilder() QueryBuilder {
	return &queryHandler{
		QueryBuilder: nil,
	}
}

func (q queryHandler) SelectAll(table, orderBy string) string {
	statement := SELECT_FROM + ORDER_BY + ";"
	return fmt.Sprintf(statement, all, table, orderBy, ASC)

}

func (q queryHandler) SelectByID(table, idColumn string, id int) string {
	statement := SELECT_FROM + WHERE_EQUAL + ";"
	return fmt.Sprintf(statement, all, table, idColumn, id)
}

func (q queryHandler) Delete(table, idColumn string, id int) string {
	statement := DELETE_FROM + WHERE_EQUAL + ";"
	return fmt.Sprintf(statement, table, idColumn, id)
}

func (q queryHandler) Update(table, keys, values, idColumn string, id int) string {
	statement := UPDATE_SET + WHERE_EQUAL + ";"
	return fmt.Sprintf(statement, table, keys, values, idColumn, id)
}

func (q queryHandler) Create(table, keys, values, returned string) string {
	statement := INSERT_INTO + VALUES + RETURNING + ";"
	return fmt.Sprintf(statement, table, keys, values, returned)
}

