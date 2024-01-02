package users

import (
	"context"
	"detaskify/internal/utils"
	"gorm.io/datatypes"
)

type User struct {
	Username     string         `json:"username" validate:"required,max=255"`
	ProfilePhoto string         `json:"profile_photo"`
	Banner       string         `json:"banner"`
	Email        string         `json:"email" validate:"required,email"`
	Provider     string         `json:"provider"`
	Integrations datatypes.JSON `json:"integrations"`
	Technologies datatypes.JSON `json:"technologies"`
	Availability bool           `json:"availability"`
	Teams        []*Team        `json:"-"`
	Password     string         `json:"-"`
	SocialLinks  datatypes.JSON `json:"social_links"`
	IsVerified   bool           `json:"is_verified"`
}

type socials struct {
	Twitter  string `json:"twitter"`
	LinkedIn string `json:"linkedin"`
	GitHub   string `json:"github"`
	Wakatime string `json:"wakatime"`
	Website  string `json:"website" validate:"url"`
}

type UserService interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	UpdateUser(ctx context.Context, username string, user *User) error
	DeleteUser(ctx context.Context, username string) error
	UpdateUserPassword(ctx context.Context, username, newPassword string) error
	ValidateSignIn(ctx context.Context, identifier, password string) (bool, error)
}

// UserRepository is the blueprint for user-related logic
type UserRepository struct {
	service UserService
}

// NewUserService creates a new user service
func NewUserService(service UserService) UserRepository {
	return UserRepository{
		service: service,
	}
}

func (u *UserRepository) CreateUser(ctx context.Context, user *User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return u.service.CreateUser(ctx, user)

}

func (u *UserRepository) ValidateSignIn(ctx context.Context, identifier, password string) (bool, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return false, err
	}

	return u.service.ValidateSignIn(ctx, identifier, hashedPassword)

}

func (u *UserRepository) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	return u.service.GetUserByUsername(ctx, username)
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	return u.service.GetUserByEmail(ctx, email)
}
func (u *UserRepository) UpdateUser(ctx context.Context, username string, user *User) error {
	return u.service.UpdateUser(ctx, username, user)
}
func (u *UserRepository) DeleteUser(ctx context.Context, user string) error {
	return u.service.DeleteUser(ctx, user)

}

// TODO:
// For the reset functionality,
// the user should be tried to sign in and then if correct then they can update the password

func (u *UserRepository) UpdateUserPassword(ctx context.Context, username, newPassword string) error {
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}
	return u.service.UpdateUserPassword(ctx, username, hashedPassword)
}
