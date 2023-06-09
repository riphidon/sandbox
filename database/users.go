package database

import (
	"database/sql"
	"fmt"
)

const (
	userTable    string = "users"
	userKeys     string = "first_name, last_name, email, nickname, password_hash, remember, remember_hash"
	userIdKey    string = "user_id"
	userRemember string = "remember"
)

type User struct {
	Id           int
	FirstName    string
	LastName     string
	Email        string
	Nickname     string
	Password     string
	PasswordHash string
	Remember     string
	RememberHash string
	CreatedAt    string
}

type userAccess struct {
	db *sql.DB
}

type userRepo struct {
	access userAccess
	qb     QueryBuilder
}

// ByEmail implements IUserRepo.
func (u *userRepo) UserByEmail(email string) (*User, error) {
	panic("unimplemented")
}

// ByID implements IUserRepo.
func (u *userRepo) UserByID(id uint) (*User, error) {
	panic("unimplemented")
}

// ByRemember implements IUserRepo.
func (u *userRepo) UserByRemember(token string) (*User, error) {
	panic("unimplemented")
}

// Create implements IUserRepo.
func (u *userRepo) CreateUser(user User) (int, error) {
	id := 0
	values := fmt.Sprintf("'%v', '%v', '%v', '%v', '%v', '%v', '%v'", user.FirstName, user.LastName, user.Email, user.Nickname, user.PasswordHash, user.Remember, user.RememberHash)
	query := u.qb.Create(userTable, userKeys, values, userIdKey)
	err := u.access.db.QueryRow(query).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

// Delete implements IUserRepo.
func (u *userRepo) DeleteUser(id uint) error {
	panic("unimplemented")
}

// Update implements IUserRepo.
func (u *userRepo) UpdateUser(user *User) error {
	panic("unimplemented")
}

type IUserRepo interface {
	// Single user querying
	UserByID(id uint) (*User, error)
	UserByEmail(email string) (*User, error)
	UserByRemember(token string) (*User, error)

	// Users altering methods
	CreateUser(user User) (int, error)
	UpdateUser(user *User) error
	DeleteUser(id uint) error
}

func newUserStore(db *sql.DB) IUserRepo {
	q := NewQueryBuilder()
	return &userRepo{
		access: userAccess{
			db: db,
		},
		qb: q,
	}
}
