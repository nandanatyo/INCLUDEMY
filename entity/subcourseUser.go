package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserSubcourse struct {
	ID          uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:varchar(36);not null;"`
	SubcourseID uuid.UUID `json:"sub_course_id" gorm:"type:varchar(36);not null;"`
	Checked     bool      `json:"checked" gorm:"type:boolean;not null;"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
