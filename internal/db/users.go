package db

import (
	"context"
	"detaskify/internal/users" // Adjust this import path to where your User struct is defined
	"errors"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username     string `gorm:"size:255;not null;unique"`
	ProfilePhoto string
	Email        string         `gorm:"size:255;not null;unique"`
	Provider     string         `gorm:"default:traditional"`
	Integrations datatypes.JSON `gorm:"type:json"`
	Technologies datatypes.JSON `gorm:"type:json"`
	Availability bool           `gorm:"type:boolean"`
	Teams        []*Team        `gorm:"many2many:user_teams;"`
	Password     string         `gorm:"size:255"`
	SocialLinks  datatypes.JSON `gorm:"type:json"`
	IsVerified   bool           `gorm:"type:boolean"`
}

func (d *Database) CreateUser(ctx context.Context, user *users.User) error {
	newUser := &users.User{
		Username:     user.Username,
		ProfilePhoto: user.ProfilePhoto,
		Email:        user.Email,
		Provider:     user.Provider,
		Integrations: user.Integrations,
		Technologies: user.Technologies,
		Availability: user.Availability,
		Teams:        user.Teams,
		Password:     user.Password, // Consider hashing the password before storing
		SocialLinks:  user.SocialLinks,
		IsVerified:   user.IsVerified,
	}

	err := d.Client.WithContext(ctx).Create(newUser).Error
	if err != nil {
		log.Printf("Error creating user: %s", err.Error())
		return err
	}

	return nil
}

func (d *Database) GetUserByUsername(ctx context.Context, username string) (*users.User, error) {
	var user users.User
	err := d.Client.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		log.Printf("Error getting user: %s", err.Error())
		return nil, err
	}
	return &user, nil
}

func (d *Database) GetUserByEmail(ctx context.Context, email string) (*users.User, error) {
	var user users.User
	err := d.Client.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Printf("Error getting user: %s", err.Error())
		return nil, err
	}
	return &user, nil
}

func (d *Database) UpdateUser(ctx context.Context, username string, updateData *users.User) error {
	updateData.Password = ""

	result := d.Client.WithContext(ctx).Model(&users.User{}).Where("username = ?", username).Updates(updateData)
	if result.Error != nil {
		log.Printf("Error updating user: %s", result.Error.Error())
		return result.Error
	}
	return nil
}

func (d *Database) DeleteUser(ctx context.Context, username string) error {
	result := d.Client.WithContext(ctx).Where("username = ?", username).Delete(&users.User{})
	if result.Error != nil {
		log.Printf("Error deleting user: %s", result.Error.Error())
		return result.Error
	}
	return nil
}

func (d *Database) ValidateSignInData(ctx context.Context, identifier, password string) (bool, error) {
	var user users.User

	err := d.Client.WithContext(ctx).Where("username = ? OR email = ?", identifier, identifier).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		log.Printf("Error getting user: %s", err.Error())
		return false, err
	}

	if user.Password != password { // Consider using a password comparison function for hashed passwords
		return false, nil
	}

	return true, nil
}

func (d *Database) UpdateUserPassword(ctx context.Context, username, newPassword string) error {
	result := d.Client.WithContext(ctx).Model(&users.User{}).Where("username = ?", username).Update("password", newPassword) // Consider hashing the new password
	if result.Error != nil {
		log.Printf("Error updating user password: %s", result.Error.Error())
		return result.Error
	}
	return nil
}
