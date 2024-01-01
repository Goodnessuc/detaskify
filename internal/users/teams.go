package users

import (
	"gorm.io/datatypes"
)

type Team struct {
	Name         string         `json:"name" validate:"required,max=255"`
	Members      []*User        `json:"-"`
	Projects     datatypes.JSON `json:"projects"`
	Description  string         `json:"description"`
	ProfilePhoto string         `json:"profile_photo"`
	Banner       string         `json:"banner"`
	IsVerified   bool           `json:"is_verified"`
}

type UserTeam struct {
	UserID uint `json:"user_id"`
	TeamID uint `json:"team_id"`
}
