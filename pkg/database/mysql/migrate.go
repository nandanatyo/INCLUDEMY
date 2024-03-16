package mysql

import (
	"gorm.io/gorm"
	"includemy/entity"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Course{},
		&entity.Role{},
		&entity.Subcourse{},
		&entity.UserSubcourse{},
		&entity.UserJoinCourse{},
		&entity.Job{},
		&entity.Applicant{},
		&entity.Sertification{},
		&entity.SertificationUser{},
	); err != nil {
		return err
	}
	return nil
}
