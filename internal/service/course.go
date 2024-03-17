package service

import (
	"fmt"
	"includemy/entity"
	"includemy/internal/repository"
	"includemy/pkg/supabase"
	"strings"
	"time"

	"includemy/model"

	"errors"

	"github.com/google/uuid"
)

type ICourseService interface {
	CreateCourse(courseReq *model.CourseReq) (*entity.Course, error)
	GetCourseByAny(param model.CourseSearch) ([]*entity.Course, error)
	DeleteCourse(id string) error
	GetSubcourseWithinCourse(temp model.CourseGet) (entity.Course, error)
	UpdateCourse(id string, modifyCourse *model.CourseReq) (*entity.Course, error)
	UploadCoursePhoto(param model.CoursePhoto) (entity.Course, error)
}

type CourseService struct {
	course   repository.ICourseRepository
	supabase supabase.Interface
}

func NewCourseService(course repository.ICourseRepository, supabase supabase.Interface) *CourseService {
	return &CourseService{
		course:   course,
		supabase: supabase,
	}
}

func (cs *CourseService) CreateCourse(courseReq *model.CourseReq) (*entity.Course, error) {
	course := &entity.Course{
		ID:             uuid.New(),
		Title:          courseReq.Title,
		Teacher:        courseReq.Teacher,
		Company:        courseReq.Company,
		Price:          courseReq.Price,
		Description:    courseReq.Description,
		HowMuchTime:    courseReq.HowMuchTime,
		HowManyStudent: courseReq.HowManyStudent,
		HowManyCourse:  courseReq.HowManyCourse,
		Tags:           courseReq.Tags,
		About:          courseReq.About,
		Dissability: courseReq.Dissability,
	}

	course, err := cs.course.CreateCourse(course)
	if err != nil {
		return nil, errors.New("Service: Failed to create course")
	}
	return course, nil
}

func (cs *CourseService) GetCourseByAny(param model.CourseSearch) ([]*entity.Course, error) {

	if param.ID != uuid.Nil {
		course, err := cs.course.GetCourseByID(param.ID.String())
		if err != nil {
			return nil, errors.New("Service: Course not found by ID")
		}
		fmt.Print(course)
		return []*entity.Course{course}, nil
	} else if param.Title != "" {
		course, err := cs.course.GetCourseByName(param.Title)
		if err != nil {
			return nil, errors.New("Service: Course not found by title")
		}
		return course, nil
	} else if param.Tags != ""{
		course, err := cs.course.GetCourseByTags(param.Tags)
		if err != nil{
			return nil, errors.New("Service: Course not found by tags")
		}
		return course, nil
	} else if param.Dissability != ""{
		course, err := cs.course.GetCourseByDissability(param.Dissability)
		if err != nil{
			return nil, errors.New("Service: Course not found by dissability")
		}
		return course, nil
	} else {
		return nil, errors.New("Service: No search criteria provided")
	}
}

func (cs *CourseService) UpdateCourse(id string, modifyCourse *model.CourseReq) (*entity.Course, error) {
	course, err := cs.course.UpdateCourse(id, modifyCourse)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (cs *CourseService) UploadCoursePhoto(param model.CoursePhoto) (entity.Course, error) {
	course, err := cs.course.GetCourseByID(param.CourseID.String())
	if err != nil {
		return *course, errors.New("Service: Course not found")
	}

	if course.PhotoLink != "" {
		err := cs.supabase.Delete(course.PhotoLink)
		if err != nil {
			return *course, errors.New("Service: Failed to delete previous file")
		}
	}

	param.PhotoLink.Filename = fmt.Sprintf("%s.%s", strings.ReplaceAll(time.Now().String(), " ", ""), strings.Split(param.PhotoLink.Filename, ".")[1])

	link, err := cs.supabase.UploadFile(param.PhotoLink)
	if err != nil {
		return *course, errors.New("Service: Failed to upload file")
	}

	course, err = cs.course.UpdateCourse(course.ID.String(), &model.CourseReq{
		PhotoLink: link})
	if err != nil {
		return *course, errors.New("Service: Failed to update subcourse")
	}
	return *course, nil
}

func (cs *CourseService) DeleteCourse(id string) error {
	err := cs.course.DeleteCourse(id)
	if err != nil {
		return errors.New("Service: Failed to delete course")
	}
	return nil
}

func (cs *CourseService) GetSubcourseWithinCourse(temp model.CourseGet) (entity.Course, error) {
	subcourses, err := cs.course.GetSubcourseWithinCourse(temp)
	if err != nil {
		return subcourses, errors.New("Service: Failed to load subcourse within course")
	}
	return subcourses, nil
}
