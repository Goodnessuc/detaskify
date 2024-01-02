package users

import (
	"context"
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

type TeamService interface {
	CreateTeam(ctx context.Context, newTeam Team) error
	GetTeamByID(ctx context.Context, name string) (Team, error)
	UpdateTeam(ctx context.Context, name string, updateData Team) error
	DeleteTeam(ctx context.Context, name string) error
	AddUserToTeam(ctx context.Context, teamName, username string) error
	RemoveUserFromTeam(ctx context.Context, teamName, username string) error
}

// TeamRepository is the blueprint for user-related logic
type TeamRepository struct {
	service TeamService
}

// NewTeamService creates a new team service
func NewTeamService(service TeamService) TeamRepository {
	return TeamRepository{
		service: service,
	}
}

func (t *TeamRepository) CreateTeam(ctx context.Context, newTeam Team) error {
	return t.service.CreateTeam(ctx, newTeam)
}

func (t *TeamRepository) GetTeamByID(ctx context.Context, name string) (Team, error) {
	team, err := t.service.GetTeamByID(ctx, name)
	return team, err

}
func (t *TeamRepository) UpdateTeam(ctx context.Context, name string, updateData Team) error {
	return t.service.UpdateTeam(ctx, name, updateData)

}
func (t *TeamRepository) DeleteTeam(ctx context.Context, name string) error {
	return t.service.DeleteTeam(ctx, name)

}
func (t *TeamRepository) AddUserToTeam(ctx context.Context, teamName, username string) error {
	return t.service.AddUserToTeam(ctx, teamName, username)

}
func (t *TeamRepository) RemoveUserFromTeam(ctx context.Context, teamName, username string) error {
	return t.service.RemoveUserFromTeam(ctx, teamName, username)

}
