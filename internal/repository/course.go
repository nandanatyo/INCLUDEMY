package repository

import (
	"errors"
	"includemy/entity"
	"includemy/model"

	"gorm.io/gorm"
)

type ICourseRepository interface {
	CreateCourse(course *entity.Course) (*entity.Course, error)
	GetCourseByID(id string) (*entity.Course, error)
	GetCourseByName(title string) ([]*entity.Course, error)
	GetCourseByDissability(dissability string) ([]*entity.Course, error)
	GetCourseByTags(tags string) ([]*entity.Course, error)
	DeleteCourse(id string) error
	GetSubcourseWithinCourse(param model.CourseGet) (entity.Course, error)
	UpdateCourse(id string, modifyCourse *model.CourseReq) (*entity.Course, error)
}

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) ICourseRepository {
	return &CourseRepository{db}
}

func (cr *CourseRepository) CreateCourse(course *entity.Course) (*entity.Course, error) { //OnlyAdmin
	if err := cr.db.Debug().Create(course).Error; err != nil {
		return nil, errors.New("Repository: Failed to create course")
	}
	return course, nil
}

func (cr *CourseRepository) GetCourseByID(id string) (*entity.Course, error) {
	var course entity.Course
	if err := cr.db.Debug().Where("id = ?", id).First(&course).Error; err != nil {
		return nil, errors.New("Repository: Course not found")
	}
	return &course, nil
}

func (cr *CourseRepository) GetCourseByName(title string) ([]*entity.Course, error) {
	var course []*entity.Course
	if err := cr.db.Debug().Where("title LIKE ?", "%"+title+"%").Find(&course).Error; err != nil {
		return nil, errors.New("Repository: Course not found")
	}
	return course, nil
}

func (cr *CourseRepository) GetCourseByTags(tags string) ([]*entity.Course, error) {
	var course []*entity.Course
	if err := cr.db.Debug().Where("tags LIKE ?", "%"+tags+"%").Find(&course).Error; err != nil {
		return nil, errors.New("Repository: Course not found")
	}
	return course, nil
}

func (cr *CourseRepository) GetCourseByDissability(dissability string) ([]*entity.Course, error) {
	var course []*entity.Course
	if err := cr.db.Debug().Where("dissability LIKE ?", "%"+dissability+"%").Find(&course).Error; err != nil {
		return nil, errors.New("Repository: Course not found")
	}
	return course, nil
}

func (cr *CourseRepository) GetSubcourseWithinCourse(param model.CourseGet) (entity.Course, error) {
	course := entity.Course{}
	err := cr.db.Debug().Where("id = ?", param.ID).Preload("Subcourse").First(&course).Error
	if err != nil {
		return course, errors.New("Repository: Failed to load subcourse within course")
	}
	return course, nil
}

func (cr *CourseRepository) DeleteCourse(id string) error {
	if err := cr.db.Debug().Where("id = ?", id).Delete(&entity.Course{}).Error; err != nil {
		return errors.New("Repository: Failed to delete course")
	}
	return nil
}

func (cr *CourseRepository) UpdateCourse(id string, modifyCourse *model.CourseReq) (*entity.Course, error) {
	var course entity.Course
	if err := cr.db.Debug().Where("id = ?", id).First(&course).Error; err != nil {
		return nil, err
	}

	courseParse := parseUpdate(modifyCourse, &course)
	if err := cr.db.Model(&course).Save(&courseParse).Error; err != nil {
		return nil, err
	}
	return courseParse, nil
}

func parseUpdate(modifyCourse *model.CourseReq, course *entity.Course) *entity.Course {
	if modifyCourse.Title != "" {
		course.Title = modifyCourse.Title
	}
	if modifyCourse.Teacher != "" {
		course.Teacher = modifyCourse.Teacher
	}
	if modifyCourse.Company != "" {
		course.Company = modifyCourse.Company
	}
	if modifyCourse.Description != "" {
		course.Description = modifyCourse.Description
	}
	if modifyCourse.Tags != "" {
		course.Tags = modifyCourse.Tags
	}
	if modifyCourse.About != "" {
		course.About = modifyCourse.About
	}
	if modifyCourse.HowManyCourse != 0 {
		course.HowManyCourse = modifyCourse.HowManyCourse
	}
	if modifyCourse.HowManyStudent != 0 {
		course.HowManyStudent = modifyCourse.HowManyStudent
	}
	if modifyCourse.HowMuchTime != 0 {
		course.HowMuchTime = modifyCourse.HowMuchTime
	}
	if modifyCourse.Price != 0 {
		course.Price = modifyCourse.Price
	}
	if modifyCourse.PhotoLink != "" {
		course.PhotoLink = modifyCourse.PhotoLink
	}
	if modifyCourse.Dissability != "" {
		course.Dissability = modifyCourse.Dissability
	}

	return course
}
