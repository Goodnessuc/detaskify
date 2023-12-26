package db

import (
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username     string `gorm:"size:255;not null;unique"`
	ProfilePhoto string
	Email        string         `gorm:"size:255;not null;unique"`
	Integrations pq.StringArray `gorm:"type:text[]"`
	Technologies pq.StringArray `gorm:"type:text[]"`
	Availability bool           `gorm:"type:boolean"`

	// GitHub, Wakatime, GitLab, Linkedin
	SocialLinks *datatypes.JSON `gorm:"type:json"`
}
