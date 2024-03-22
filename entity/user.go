package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                uuid.UUID           `json:"id" gorm:"type:varchar(36);primary_key;"`
	Name              string              `json:"name" gorm:"type:varchar(255);not null;"`
	Email             string              `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password          string              `json:"password" gorm:"type:varchar(255);not null;"`
	Role              int                 `json:"role" gorm:"foreignKey:ID; references:roles; not null;"`
	Born              time.Time           `json:"born" gorm:"type:date;not null;"`
	Gender            string              `json:"gender" gorm:"type:varchar(10);not null;"`
	LastJob           string              `json:"lastjob" gorm:"type:varchar(255)"`
	LastEdu           string              `json:"lastedu" gorm:"type:varchar(255)"`
	Contact           string              `json:"contact" gorm:"type:varchar(255)"`
	PhotoLink         string              `json:"photolink" gorm:"type:varchar(200)"`
	CreatedAt         time.Time           `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt         time.Time           `json:"updatedAt" gorm:"autoUpdateTime"`
	UserJoinCourse    []UserJoinCourse    `json:"user_join_course" gorm:"foreignKey:UserID;references:ID"`
	Dissability       string              `json:"dissability" gorm:"type:varchar(255)"`
	Preference        string              `json:"preference" gorm:"type:varchar(255)"`
	Applicant         []Applicant         `json:"applicant" gorm:"foreignKey:UserID;references:ID"`
	CertificationUser []CertificationUser `json:"certification_user" gorm:"foreignKey:UserID;references:ID"`
}
