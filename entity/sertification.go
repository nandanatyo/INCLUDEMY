package entity

import (
	"time"

	"github.com/google/uuid"
)

type Sertification struct {
	ID                uuid.UUID           `json:"id" gorm:"type:varchar(36);primary_key;"`
	Title             string              `json:"title" gorm:"type:varchar(255);not null;"`
	Creator           string              `json:"creator" gorm:"type:varchar(255);not null;"`
	About             string              `json:"about" gorm:"type:varchar(255);not null;"`
	Field             string              `json:"field" gorm:"type:varchar(255);not null;"`
	Tags              string              `json:"tags" gorm:"type:varchar(255);not null;"`
	Location          string              `json:"location" gorm:"type:varchar(255);not null;"`
	Link              string              `json:"link" gorm:"type:varchar(255);not null;"`
	PhotoLink         string              `json:"photolink" gorm:"type:varchar(200)"`
	CreatedAt         time.Time           `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt         time.Time           `json:"updatedAt" gorm:"autoUpdateTime"`
	SertificationUser []SertificationUser `json:"sertification_user" gorm:"foreignKey:SertificationID;references:ID"`
}
