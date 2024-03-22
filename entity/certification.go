package entity

import (
	"time"

	"github.com/google/uuid"
)

type Certification struct {
	ID                uuid.UUID           `json:"id" gorm:"type:varchar(36);primary_key;"`
	Title             string              `json:"title" gorm:"type:varchar(255);not null;"`
	Creator           string              `json:"creator" gorm:"type:varchar(255);not null;"`
	About             string              `json:"about" gorm:"type:varchar(255);not null;"`
	Field             string              `json:"field" gorm:"type:varchar(255);not null;"`
	Tags              string              `json:"tags" gorm:"type:varchar(255);not null;"`
	Syllabus          string              `json:"syllabus" gorm:"type:varchar(255);not null;"`
	Dissability       string              `json:"dissability" gorm:"type:varchar(255);not null;"`
	Location          string              `json:"location" gorm:"type:varchar(255);not null;"`
	Link              string              `json:"link" gorm:"type:varchar(255);not null;"`
	PhotoLink         string              `json:"photolink" gorm:"type:varchar(200)"`
	Price             int                 `json:"price" gorm:"type:int;not null;"`
	CreatedAt         time.Time           `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt         time.Time           `json:"updatedAt" gorm:"autoUpdateTime"`
	CertificationUser []CertificationUser `json:"certification_user" gorm:"foreignKey:CertificationID;references:ID"`
}
