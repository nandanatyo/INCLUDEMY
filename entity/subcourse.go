package entity

import (
	"time"

	"github.com/google/uuid"
)

type Subcourse struct {
	ID            uuid.UUID       `json:"id" gorm:"type:varchar(36);primary_key;"`
	CourseID      uuid.UUID       `json:"course_id" gorm:"type:varchar(36);"`
	Title         string          `json:"title" gorm:"type:varchar(255);not null;"`
	Description   string          `json:"description" gorm:"type:varchar(255);not null;"`
	HowMuchTime   int             `json:"how_much_time" gorm:"type:int;not null;"`
	VideoLink     string          `json:"video_link" gorm:"type:varchar(255);not null;"`
	UserSubcourse []UserSubcourse `gorm:"foreignKey:SubcourseID;references:ID"`
	CreatedAt     time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
}
