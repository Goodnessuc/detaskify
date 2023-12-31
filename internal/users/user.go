package users

import (
	"context"
	"detaskify/internal/utils"
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username     string          `json:"username"`
	ProfilePhoto string          `json:"profile_photo"`
	Email        string          `json:"email"`
	Integrations pq.StringArray  `json:"integrations""`
	Technologies pq.StringArray  `json:"technologies"`
	Availability bool            `json:"availability"`
	SocialLinks  *datatypes.JSON `json:"social_links"`
	Provider     string          `json:"provider"`
	Password     string          `json:"password"`
	IsVerified   bool            `json:"is_verified"`
	Company      string          `json:"company"`
}

const (
	Wakatime = iota
	GitHub
	Website
	Twitter
	LinkedIn
)

const (
	user = iota
	organization
)

type socials struct {
	Twitter  string `json:"twitter"`
	LinkedIn string `json:"linkedin"`
	GitHub   string `json:"github"`
	Wakatime string `json:"wakatime"`
	Website  string `json:"website" validate:"url"`
}

type UserService interface {
	CreateUser(ctx context.Context, user *Users) error
	GetUserByUsername(ctx context.Context, username string) (*Users, error)
	GetUserByEmail(ctx context.Context, email string) (*Users, error)
	UpdateUser(ctx context.Context, username string, user *Users) error
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

func (u *UserRepository) CreateUser(ctx context.Context, user *Users) error {
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

func (u *UserRepository) GetUserByUsername(ctx context.Context, username string) (*Users, error) {
	return u.service.GetUserByUsername(ctx, username)
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*Users, error) {
	return u.service.GetUserByEmail(ctx, email)
}
func (u *UserRepository) UpdateUser(ctx context.Context, username string, user *Users) error {
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
