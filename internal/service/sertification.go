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

type ISertificationService interface {
	CreateSertification(SertificationReq *model.SertificationReq) (*entity.Sertification, error)
	UpdateSertification(id string, modifySertif *model.SertificationReq) (*entity.Sertification, error)
	GetSertificationByTitleOrID(param model.SertifSearch) ([]*entity.Sertification, error)
	DeleteSertification(id string) error
	UploadSertificationFile(param model.SertifPost) (entity.Sertification, error)
}

type SertificationService struct {
	SertificationRepository repository.ISertificationRepository
	supabase                supabase.Interface
}

func NewSertificationService(sertif repository.ISertificationRepository, supabase supabase.Interface) *SertificationService {
	return &SertificationService{
		SertificationRepository: sertif,
		supabase:                supabase,
	}
}

func (ss *SertificationService) CreateSertification(SertificationReq *model.SertificationReq) (*entity.Sertification, error) {
	sertification := &entity.Sertification{
		ID:          uuid.New(),
		Title:       SertificationReq.Title,
		About:       SertificationReq.About,
		Tags:        SertificationReq.Tags,
		Link:        SertificationReq.Link,
		Creator:     SertificationReq.Creator,
		Field:       SertificationReq.Field,
		Syllabus:    SertificationReq.Syllabus,
		Location:    SertificationReq.Location,
		Dissability: SertificationReq.Dissability,
		Price:       SertificationReq.Price,
	}

	sertification, err := ss.SertificationRepository.CreateSertification(sertification)
	if err != nil {
		return nil, errors.New("Service: Failed to create sertification")
	}
	return sertification, nil
}

func (ss *SertificationService) GetSertificationByTitleOrID(param model.SertifSearch) ([]*entity.Sertification, error) {

	if param.ID != uuid.Nil {
		sertif, err := ss.SertificationRepository.GetSertificationByID(param.ID.String())
		if err != nil {
			return nil, errors.New("Service: Sertification not found by ID")
		}
		return []*entity.Sertification{sertif}, nil
	} else if param.Title != "" {
		sertif, err := ss.SertificationRepository.GetSertiicationByName(param.Title)
		if err != nil {
			return nil, errors.New("Service: Sertification not found by title")
		}
		return sertif, nil
	} else {
		return nil, errors.New("Service: No search criteria provided")
	}
}

func (ss *SertificationService) DeleteSertification(id string) error {
	err := ss.SertificationRepository.DeleteSertification(id)
	if err != nil {
		return errors.New("Service: Failed to delete sertification")
	}
	return nil
}

func (ss *SertificationService) UpdateSertification(id string, modifySertif *model.SertificationReq) (*entity.Sertification, error) {
	sertif, err := ss.SertificationRepository.UpdateSertification(id, modifySertif)
	if err != nil {
		return nil, err
	}
	return sertif, nil
}

func (ss *SertificationService) UploadSertificationFile(param model.SertifPost) (entity.Sertification, error) {
	sertif, err := ss.SertificationRepository.GetSertificationByID(param.ID.String())
	if err != nil {
		return *sertif, errors.New("Service: Sertification not found")
	}

	if sertif.PhotoLink != "" {
		err := ss.supabase.Delete(sertif.PhotoLink)
		if err != nil {
			return *sertif, errors.New("Service: Failed to delete previous file")
		}
	}

	param.File.Filename = fmt.Sprintf("%s %s ", time.Now().String(), param.File.Filename)

	link, err := ss.supabase.UploadFile(param.File)
	if err != nil {
		return *sertif, errors.New("Service: Failed to upload file")
	}

	_, err = ss.SertificationRepository.UpdateSertification(sertif.ID.String(), &model.SertificationReq{
		PhotoLink: link})
	if err != nil {
		return *sertif, errors.New("Service: Failed to update sertification")
	}
	return *sertif, nil
}
