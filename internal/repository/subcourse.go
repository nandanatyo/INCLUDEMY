package repository

import (
	"errors"
	"includemy/entity"
	"includemy/model"

	"gorm.io/gorm"
)

type ISubcourseRepository interface {
	CreateSubcourse(subcourse *entity.Subcourse) (*entity.Subcourse, error)
	UpdateSubcourse(modifySub *model.SubcourseParam, id string) (*entity.Subcourse, error)
	GetSubcourse(param model.SubcourseParam) (entity.Subcourse, error)
	GetSubcourseByID(id string) (*entity.Subcourse, error)
	DeleteSubcourse(id string) error
}

type SubcourseRepository struct {
	db *gorm.DB
}

func NewSubcourseRepository(db *gorm.DB) ISubcourseRepository {
	return &SubcourseRepository{
		db: db,
	}
}

func (scr *SubcourseRepository) CreateSubcourse(subcourse *entity.Subcourse) (*entity.Subcourse, error) {
	err := scr.db.Debug().Create(&subcourse).Error
	if err != nil {
		return subcourse, errors.New("Repository: Failed to create subcourse")
	}
	return subcourse, nil
}

func (scr *SubcourseRepository) GetSubcourse(param model.SubcourseParam) (entity.Subcourse, error) {
	subcourse := entity.Subcourse{}
	err := scr.db.Debug().Where(&param).First(&subcourse).Error
	if err != nil {
		return subcourse, errors.New("Repository: Subcourse not found")
	}
	return subcourse, nil
}

func (scr *SubcourseRepository) UpdateSubcourse(modifySub *model.SubcourseParam, id string) (*entity.Subcourse, error) {
	var subcourse entity.Subcourse

	if err := scr.db.Debug().Where("id = ?", id).First(&subcourse).Error; err != nil {
		return nil, err
	}

	subParse, err := parseUpdateSub(modifySub, &subcourse)
	if err != nil {
		return nil, err
	}

	if err := scr.db.Model(&subcourse).Save(&subParse).Error; err != nil {
		return nil, err
	}
	return subParse, nil
}

func parseUpdateSub(modifySub *model.SubcourseParam, subcourse *entity.Subcourse) (*entity.Subcourse, error) {
	if modifySub.Title != "" {
		subcourse.Title = modifySub.Title
	}
	if modifySub.Description != "" {
		subcourse.Description = modifySub.Description
	}
	if modifySub.HowMuchTime != 0 {
		subcourse.HowMuchTime = modifySub.HowMuchTime
	}
	if modifySub.VideoLink != "" {
		subcourse.VideoLink = modifySub.VideoLink
	}
	return subcourse, nil
}

func (scr *SubcourseRepository) GetSubcourseByID(id string) (*entity.Subcourse, error) {
	subcourse := entity.Subcourse{}
	err := scr.db.Debug().Where("id = ?", id).First(&subcourse).Error
	if err != nil {
		return nil, errors.New("Repository: Subcourse not found")
	}
	return &subcourse, nil
}

func (scr *SubcourseRepository) DeleteSubcourse(id string) error {
	if err := scr.db.Debug().Where("id = ?", id).Delete(&entity.Subcourse{}).Error; err != nil {
		return errors.New("Repository: Failed to delete subcourse")
	}
	return nil
}
