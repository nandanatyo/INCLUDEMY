package entity

import (
	"time"

	"github.com/google/uuid"
)

type SertificationUser struct {
	ID              uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	UserID          uuid.UUID `json:"userid" gorm:"type:varchar(36);not null;"`
	SertificationID uuid.UUID `json:"sertificationid" gorm:"type:varchar(36);not null;"`
	Pass            bool      `json:"pass" gorm:"type:boolean;not null;"`
	CreatedAt       time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
