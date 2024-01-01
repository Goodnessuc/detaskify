package db

import (
	"context"
	"detaskify/internal/users"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"log"
	"time"
)

// Team represents a team in the system.
type Team struct {
	gorm.Model
	Name         string         `gorm:"size:255;not null;unique"`
	Members      []*User        `gorm:"many2many:user_teams;"`
	Projects     datatypes.JSON `gorm:"type:json"`
	Description  string         `gorm:"size:255"`
	ProfilePhoto string
	Banner       string
	IsVerified   bool `gorm:"type:boolean"`type UserTeam struct {
	UserID    uint `gorm:"primaryKey"`
	TeamID    uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

}


// CreateTeam creates a new team in the database.
func (d *Database) CreateTeam(ctx context.Context, newTeam *users.Team) error {
	err := d.Client.WithContext(ctx).Create(newTeam).Error
	if err != nil {
		log.Printf("Error creating team: %s", err.Error())
		return err
	}

	return nil
}

func (d *Database) GetTeamByID(ctx context.Context, name string) (*users.Team, error) {
	var tm users.Team
	err := d.Client.WithContext(ctx).Where("name = ?", name).First(&tm).Error
	if err != nil {
		log.Printf("Error getting team: %s", err.Error())
		return nil, err
	}
	return &tm, nil
}

// UpdateTeam updates a team's information in the database.
func (d *Database) UpdateTeam(ctx context.Context, name string, updateData *users.Team) error {
	result := d.Client.WithContext(ctx).Model(&users.Team{}).Where("name = ?", name).Updates(updateData)
	if result.Error != nil {
		log.Printf("Error updating team: %s", result.Error.Error())
		return result.Error
	}
	return nil
}

// DeleteTeam deletes a team from the database.
func (d *Database) DeleteTeam(ctx context.Context, name string) error {
	result := d.Client.WithContext(ctx).Where("name = ?", name).Delete(&users.Team{})
	if result.Error != nil {
		log.Printf("Error deleting team: %s", result.Error.Error())
		return result.Error
	}
	return nil
}

func (d *Database) AddUserToTeam(ctx context.Context, teamName, username string) error {
	var teamID uint
	var userID uint

	// Use a transaction for multiple related operations
	tx := d.Client.WithContext(ctx).Begin()

	// Find the team ID by name
	if err := tx.Table("teams").Select("id").Where("name = ?", teamName).Scan(&teamID).Error; err != nil {
		tx.Rollback()
		log.Printf("Error finding team: %s", err.Error())
		return err
	}

	// Find the user ID by username
	if err := tx.Table("users").Select("id").Where("username = ?", username).Scan(&userID).Error; err != nil {
		tx.Rollback()
		log.Printf("Error finding user: %s", err.Error())
		return err
	}

	// Insert into UserTeam
	userTeam := UserTeam{UserID: userID, TeamID: teamID, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if err := tx.Create(&userTeam).Error; err != nil {
		tx.Rollback()
		log.Printf("Error adding user to team: %s", err.Error())
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}

func (d *Database) RemoveUserFromTeam(ctx context.Context, teamName, username string) error {
	var teamID uint
	var userID uint

	// Use a transaction for consistency
	tx := d.Client.WithContext(ctx).Begin()

	// Find the team ID by name
	if err := tx.Table("teams").Select("id").Where("name = ?", teamName).Scan(&teamID).Error; err != nil {
		tx.Rollback()
		log.Printf("Error finding team: %s", err.Error())
		return err
	}

	// Find the user ID by username
	if err := tx.Table("users").Select("id").Where("username = ?", username).Scan(&userID).Error; err != nil {
		tx.Rollback()
		log.Printf("Error finding user: %s", err.Error())
		return err
	}

	// Delete the association from UserTeam
	if err := tx.Where("user_id = ? AND team_id = ?", userID, teamID).Delete(&UserTeam{}).Error; err != nil {
		tx.Rollback()
		log.Printf("Error removing user from team: %s", err.Error())
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}
