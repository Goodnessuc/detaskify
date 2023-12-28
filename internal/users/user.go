package users

import (
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username     string         `json:"username"`
	ProfilePhoto string         `json:"profile_photo"`
	Email        string         `json:"email"`
	Integrations pq.StringArray `json:"integrations""`
	Technologies pq.StringArray `json:"technologies"`
	Availability bool           `json:"availability"`

	// GitHub, Wakatime, GitLab, Linkedin Website
	SocialLinks *datatypes.JSON `json:"social_links"`
	IsVerified  string          `json:"is_verified"`
	Company     string          `json:"company"`
}
