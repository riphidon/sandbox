package services

import (
	"context"
	"sandbox-api/database"

	"golang.org/x/crypto/bcrypt"
)

const userKey string = "user"

type userService struct {
	access  database.IUserRepo
	pepper  string
	HMACKey string
}

// newUser implements IUserService.
func (*userService) NewUser() *database.User {
	return &database.User{}
}

// AccessCheck implements IUserService.
func (us *userService) UserCheck(ctx context.Context) bool {
	user := CheckValidUser(ctx)
	return user != nil
}

// ByEmail implements IUserService.
func (us *userService) ByEmail(email string) (*database.User, error) {
	panic("unimplemented")
}

// ByID implements IUserService.
func (us *userService) ByID(id uint) (*database.User, error) {
	panic("unimplemented")
}

// ByRemember implements IUserService.
func (us *userService) ByRemember(token string, ctx context.Context) (context.Context, error) {
	user, err := us.access.UserByRemember(token)
	if err != nil {
		return nil, err
	}
	return UserContext(ctx, user), nil
}

// Create implements IUserService.
func (us *userService) Create(user database.User) (int, error) {
	var err error
	if user.Nickname == "" {
		user.Nickname = user.FirstName
	}
	if user.Password == "" {
		return 0, err
	}
	pwBytes := []byte(user.Password + us.pepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""

	if user.Remember == "" {
		token, err := RememberToken()
		if err != nil {
			return 0, err
		}
		user.Remember = token
	}
	hmac := NewHMAC(us.HMACKey)
	user.RememberHash = hmac.Hash(user.Remember)

	id, err := us.access.CreateUser(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Delete implements IUserService.
func (us *userService) Delete(id uint) error {
	panic("unimplemented")
}

// Update implements IUserService.
func (us *userService) Update(user *database.User) error {
	panic("unimplemented")
}

type IUserService interface {
	UserCheck(ctx context.Context) bool
	ByID(id uint) (*database.User, error)
	ByEmail(email string) (*database.User, error)
	ByRemember(token string, ctx context.Context) (context.Context, error)
	Create(user database.User) (int, error)
	NewUser() *database.User
	Update(user *database.User) error
	Delete(id uint) error
}

func NewUserService(repo database.IUserRepo, p string, h string) IUserService {
	return &userService{
		access:  repo,
		pepper:  p,
		HMACKey: h,
	}
}

func UserContext(ctx context.Context, user *database.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func CheckValidUser(ctx context.Context) *database.User {
	if temp := ctx.Value(userKey); temp != nil {
		if user, ok := temp.(*database.User); ok {
			return user
		}
	}
	return nil
}
