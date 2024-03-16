package entity

type Applicant struct {
	ID        string `json:"id" gorm:"type:varchar(36);primary_key;"`
	JobID     string `json:"job_id" gorm:"type:varchar(36);not null;"`
	UserID    string `json:"user_id" gorm:"type:varchar(36)not null;"`
	FirstName string `json:"first_name" gorm:"type:varchar(255);not null;"`
	LastName  string `json:"last_name" gorm:"type:varchar(255);not null;"`
	Email     string `json:"email" gorm:"type:varchar(255);not null;"`
	Phone     string `json:"phone" gorm:"type:varchar(255);not null;"`
	Resume    string `json:"resume" gorm:"type:varchar(255);not null;"`
	Status    string `json:"status" gorm:"type:varchar(255);not null;"`
	CreatedAt string `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updatedAt" gorm:"autoUpdateTime"`
}
