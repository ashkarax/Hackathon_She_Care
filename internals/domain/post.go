package domain

import "gorm.io/gorm"

type Posts struct {
	gorm.Model

	CreatorId uint  `gorm:"not null"`
	Creator   Users `gorm:"foreignKey:CreatorId"`

	Title   string
	Content string

	VoteCount int `gorm:"default:0"`
	//commentcount needs to be calculated by the foreignkey relation to posts.

	Status string `gorm:"default:active"`
}
