package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserJoinCourse struct {
	ID       uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	UserID   uuid.UUID `json:"user_id" gorm:"type:varchar(36);not null;"`
	Course Course `gorm:"foreignKey:CourseID"`
	CourseID uuid.UUID `json:"course_id" gorm:"type:varchar(36);not null;"`
	UserSubcourse []UserSubcourse `gorm:"foreignKey:UserID;references:UserID"`
	CreatedAt     time.Time       `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time       `json:"updatedAt" gorm:"autoUpdateTime"`
}

type Invoice struct {
	ID               uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	OrderID          string    `json:"order_id" gorm:"type:varchar(36);not null;"`
	UserID           string    `json:"user_id" gorm:"type:varchar(36);not null;"`
	CourseorSertifID string    `json:"course_id" gorm:"type:varchar(36);not null;"`
	Status           string    `json:"status" gorm:"type:varchar(36);not null;"`
}
