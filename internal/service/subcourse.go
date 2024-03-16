package service

import (
	"errors"
	"fmt"
	"includemy/entity"
	"includemy/internal/repository"
	"includemy/model"
	"includemy/pkg/supabase"
	"time"

	"github.com/google/uuid"
)

type ISubcourseService interface {
	CreateSubcourse(subcourseReq *model.CreateSubcourse) (*entity.Subcourse, error)
	GetCourse(param model.SubcourseParam) (entity.Subcourse, error)
	UploadSubcourseFile(param model.UploadFile) (entity.Subcourse, error)
	GetSubcourseByID(id string) (*entity.Subcourse, error)
	UpdateSubcourse(id string, modifySub *model.SubcourseParam) (*entity.Subcourse, error)
	DeleteSubcourse(id string) error
}

type SubcourseService struct {
	CourseRepository    repository.ICourseRepository
	SubcourseRepository repository.ISubcourseRepository
	supabase            supabase.Interface
}

func NewSubcourseService(subcourse repository.ISubcourseRepository, course repository.ICourseRepository, supabase supabase.Interface) ISubcourseService {
	return &SubcourseService{
		CourseRepository:    course,
		SubcourseRepository: subcourse,
		supabase:            supabase,
	}
}

func (scs *SubcourseService) CreateSubcourse(subcourseReq *model.CreateSubcourse) (*entity.Subcourse, error) {

	_, err := scs.CourseRepository.GetCourseByID(subcourseReq.CourseID.String())
	if err != nil {
		return nil, errors.New("Service: Course not found")
	}

	subcourse := &entity.Subcourse{
		ID:          uuid.New(),
		CourseID:    subcourseReq.CourseID,
		Title:       subcourseReq.Title,
		Description: subcourseReq.Description,
		HowMuchTime: subcourseReq.HowMuchTime,
	}

	subcourse, err = scs.SubcourseRepository.CreateSubcourse(subcourse)
	if err != nil {
		return nil, errors.New("Service: Failed to create subcourse")
	}
	return subcourse, nil
}

func (scs *SubcourseService) GetCourse(param model.SubcourseParam) (entity.Subcourse, error) {
	return scs.SubcourseRepository.GetSubcourse(param)
}

func (scs *SubcourseService) UploadSubcourseFile(param model.UploadFile) (entity.Subcourse, error) {
	sub, err := scs.SubcourseRepository.GetSubcourse(model.SubcourseParam{ID: param.SubcourseID})
	if err != nil {
		return sub, errors.New("Service: Subcourse not found")
	}

	if sub.VideoLink != "" {
		err := scs.supabase.Delete(sub.VideoLink)
		if err != nil {
			return sub, errors.New("Service: Failed to delete previous file")
		}
	}

	param.File.Filename = fmt.Sprintf("%s %s ", time.Now().String(), param.File.Filename)

	link, err := scs.supabase.UploadFile(param.File)
	if err != nil {
		return sub, errors.New("Service: Failed to upload file")
	}

	_, err = scs.SubcourseRepository.UpdateSubcourse(&model.SubcourseParam{
		VideoLink: link}, sub.ID.String())
	if err != nil {
		return sub, errors.New("Service: Failed to update subcourse")
	}
	return sub, nil
}

func (scs *SubcourseService) GetSubcourseByID(id string) (*entity.Subcourse, error) {
	subcourse, err := scs.SubcourseRepository.GetSubcourseByID(id)
	if err != nil {
		return nil, errors.New("Service: Subcourse not found")
	}
	return subcourse, nil
}

func (scs *SubcourseService) DeleteSubcourse(id string) error {
	err := scs.SubcourseRepository.DeleteSubcourse(id)
	if err != nil {
		return errors.New("Service: Failed to delete subcourse")
	}
	return nil
}

func (scs *SubcourseService) UpdateSubcourse(id string, modifySub *model.SubcourseParam) (*entity.Subcourse, error) {
	subcourse, err := scs.SubcourseRepository.UpdateSubcourse(modifySub, id)
	if err != nil {
		return nil, err
	}
	return subcourse, nil
}
