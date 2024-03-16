package entity

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID             uuid.UUID        `json:"id" gorm:"type:varchar(36);primary_key;"`
	Title          string           `json:"title" gorm:"type:varchar(255);not null;"`
	Teacher        string           `json:"teacher" gorm:"type:varchar(255);not null;"`
	Company        string           `json:"company" gorm:"type:varchar(255);not null;"`
	Price          int              `json:"price" gorm:"type:int;not null;"`
	Description    string           `json:"description" gorm:"type:varchar(255);not null;"`
	HowMuchTime    int              `json:"how_much_time" gorm:"type:int;not null;"` //cari cara otomatis
	HowManyStudent uint             `json:"how_many_student" gorm:"type:int;not null;"`
	HowManyCourse  int              `json:"how_many_course" gorm:"type:int;not null;"`
	Tags           string           `json:"tags" gorm:"type:varchar(255);not null;"`
	About          string           `json:"about" gorm:"type:varchar(255);not null;"`
	PhotoLink      string           `json:"photolink" gorm:"type:varchar(200)"`
	Subcourse      []Subcourse      `json:"sub_course" gorm:"foreignKey:CourseID"`
	CreatedAt      time.Time        `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt      time.Time        `json:"updatedAt" gorm:"autoUpdateTime"`
	UserJoinCourse []UserJoinCourse `gorm:"foreignKey:CourseID;references:ID"`
}
