package db

import (
	"context"
	"detaskify/internal/users"
	"encoding/json"
	"errors"
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"log"
)

type Users struct {
	gorm.Model
	Username     string `gorm:"size:255;not null;unique"`
	ProfilePhoto string
	Email        string          `gorm:"size:255;not null;unique"`
	Provider     string          `gorm:"default:TRADITIONAL"`
	Integrations pq.StringArray  `gorm:"type:text[]"`
	Technologies pq.StringArray  `gorm:"type:text[]"`
	Availability bool            `gorm:"type:boolean"`
	Company      string          `gorm:"size:255"`
	Password     string          `gorm:"size:255"`
	SocialLinks  *datatypes.JSON `gorm:"type:json"`
	IsVerified   bool            `gorm:"type:boolean"`
}

func (d *Database) CreateUser(ctx context.Context, user *users.Users) error {
	var SocialLinks *datatypes.JSON
	marshaled, _ := json.Marshal(user.SocialLinks)
	SocialLinks = (*datatypes.JSON)(&marshaled)
	newUser := &users.Users{
		Username:     user.Username,
		ProfilePhoto: user.ProfilePhoto,
		Email:        user.Email,
		Integrations: user.Integrations,
		Technologies: user.Technologies,
		Availability: user.Availability,
		SocialLinks:  SocialLinks,
		IsVerified:   user.IsVerified,
		Company:      user.Company,
		Password:     user.Password,
	}

	err := d.Client.WithContext(ctx).Create(&newUser).Error
	if err != nil {
		log.Printf("Error creating user: %s", err.Error())
	}
	return nil
}

func (d *Database) GetUserByUsername(ctx context.Context, username string) (*users.Users, error) {
	var user users.Users
	err := d.Client.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		log.Printf("Error getting user: %s", err.Error())
	}
	return &user, nil
}

func (d *Database) GetUserByEmail(ctx context.Context, email string) (*users.Users, error) {
	var user users.Users
	err := d.Client.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Printf("Error getting user: %s", err.Error())
	}
	return &user, nil
}

func (d *Database) UpdateUser(ctx context.Context, username string, updateData *users.Users) error {
	// Remove the password field from updateData
	updateData.Password = ""

	// Proceed with the update
	result := d.Client.WithContext(ctx).Model(&users.Users{}).Where("username = ?", username).Updates(updateData)
	if result.Error != nil {
		log.Printf("Error updating user: %s", result.Error.Error())
		return result.Error
	}
	return nil
}

func (d *Database) DeleteUser(ctx context.Context, username string) error {
	result := d.Client.WithContext(ctx).Where("username = ?", username).Delete(&users.Users{})
	if result.Error != nil {
		log.Printf("Error deleting user: %s", result.Error.Error())
		return result.Error
	}
	return nil
}

func (d *Database) ValidateSignInData(ctx context.Context, identifier, password string) (bool, error) {
	var user users.Users

	err := d.Client.WithContext(ctx).Where("username = ? OR email = ?", identifier, identifier).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		log.Printf("Error getting user: %s", err.Error())
		return false, err
	}

	if user.Password != password {
		return false, nil // Incorrect password
	}

	return true, nil // Successful sign-in
}

func (d *Database) UpdateUserPassword(ctx context.Context, username, newPassword string) error {
	// Update only the password field
	result := d.Client.WithContext(ctx).Model(&users.Users{}).Where("username = ?", username).Update("password", newPassword)
	if result.Error != nil {
		log.Printf("Error updating user password: %s", result.Error.Error())
		return result.Error
	}
	return nil
}
