package tasks

import "gorm.io/gorm"

type TaskComment struct {
	gorm.Model
	Slug   string `json:"slug"`
	Body   string `json:"body"`
	Author string `json:"author"`
}
