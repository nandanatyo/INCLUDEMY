package entity

type OrderCourse struct {
	ID       string `json:"id" gorm:"type:varchar(36);primary_key;"`
	UserID   string `json:"user_id" gorm:"type:varchar(36);not null;"`
	CourseID string `json:"course_id" gorm:"type:varchar(36);not null;"`
	User     User   `json:"user" gorm:"foreignKey:UserID"`
	Course   Course `json:"course" gorm:"foreignKey:CourseID"`
	Paid     bool   `json:"paid" gorm:"type:boolean;not null;"`
	PaidAt   string `json:"paid_at" gorm:"type:datetime;not null;"`
}

type OrderSertification struct {
	ID              string        `json:"id" gorm:"type:varchar(36);primary_key;"`
	UserID          string        `json:"user_id" gorm:"type:varchar(36);not null;"`
	SertificationID string        `json:"sertification_id" gorm:"type:varchar(36);not null;"`
	User            User          `json:"user" gorm:"foreignKey:UserID"`
	Sertification   Sertification `json:"sertification" gorm:"foreignKey:SertificationID"`
	Paid            bool          `json:"paid" gorm:"type:boolean;not null;"`
	PaidAt          string        `json:"paid_at" gorm:"type:datetime;not null;"`
}
