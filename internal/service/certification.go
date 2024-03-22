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

type ICertificationService interface {
	CreateCertification(CertificationReq *model.CertificationReq) (*entity.Certification, error)
	UpdateCertification(id string, modifyCertif *model.CertificationReq) (*entity.Certification, error)
	GetCertificationByAny(param model.CertifSearch) ([]*entity.Certification, error)
	DeleteCertification(id string) error
	UploadCertificationFile(param model.CertifPost) (entity.Certification, error)
}

type CertificationService struct {
	CertificationRepository repository.ICertificationRepository
	supabase                supabase.Interface
}

func NewCertificationService(certif repository.ICertificationRepository, supabase supabase.Interface) *CertificationService {
	return &CertificationService{
		CertificationRepository: certif,
		supabase:                supabase,
	}
}

func (ss *CertificationService) CreateCertification(CertificationReq *model.CertificationReq) (*entity.Certification, error) {
	certification := &entity.Certification{
		ID:          uuid.New(),
		Title:       CertificationReq.Title,
		About:       CertificationReq.About,
		Tags:        CertificationReq.Tags,
		Link:        CertificationReq.Link,
		Creator:     CertificationReq.Creator,
		Field:       CertificationReq.Field,
		Syllabus:    CertificationReq.Syllabus,
		Location:    CertificationReq.Location,
		Dissability: CertificationReq.Dissability,
		Price:       CertificationReq.Price,
	}

	certification, err := ss.CertificationRepository.CreateCertification(certification)
	if err != nil {
		return nil, errors.New("Service: Failed to create certification")
	}
	return certification, nil
}

func (ss *CertificationService) GetCertificationByAny(param model.CertifSearch) ([]*entity.Certification, error) {

	if param.ID != uuid.Nil {
		certif, err := ss.CertificationRepository.GetCertificationByID(param.ID.String())
		if err != nil {
			return nil, errors.New("Service: Certification not found by ID")
		}
		return []*entity.Certification{certif}, nil
	} else if param.Title != "" {
		certif, err := ss.CertificationRepository.GetCertificationByName(param.Title)
		if err != nil {
			return nil, errors.New("Service: Certification not found by title")
		}
		return certif, nil
	} else if param.Tags != "" {
		certif, err := ss.CertificationRepository.GetCertificationByTags(param.Tags)
		if err != nil {
			return nil, errors.New("Service: Certification not found by tags")
		}
		return certif, nil
	} else if param.Dissability != "" {
		certif, err := ss.CertificationRepository.GetCertificationByDissability(param.Dissability)
		if err != nil {
			return nil, errors.New("Service: Certification not found by dissability")
		}
		return certif, nil
	} else if param.Field != "" {
		certif, err := ss.CertificationRepository.GetCertificationByField(param.Field)
		if err != nil {
			return nil, errors.New("Service: Certification not found by field")
		}
		return certif, nil
	} else {
		return nil, errors.New("Service: No search criteria provided")
	}

}

func (ss *CertificationService) DeleteCertification(id string) error {
	err := ss.CertificationRepository.DeleteCertification(id)
	if err != nil {
		return errors.New("Service: Failed to delete certification")
	}
	return nil
}

func (ss *CertificationService) UpdateCertification(id string, modifyCertif *model.CertificationReq) (*entity.Certification, error) {
	certif, err := ss.CertificationRepository.UpdateCertification(id, modifyCertif)
	if err != nil {
		return nil, err
	}
	return certif, nil
}

func (ss *CertificationService) UploadCertificationFile(param model.CertifPost) (entity.Certification, error) {
	certif, err := ss.CertificationRepository.GetCertificationByID(param.ID.String())
	if err != nil {
		return *certif, errors.New("Service: Certification not found")
	}

	if certif.PhotoLink != "" {
		err := ss.supabase.Delete(certif.PhotoLink)
		if err != nil {
			return *certif, errors.New("Service: Failed to delete previous file")
		}
	}

	param.File.Filename = fmt.Sprintf("%s %s ", time.Now().String(), param.File.Filename)

	link, err := ss.supabase.UploadFile(param.File)
	if err != nil {
		return *certif, errors.New("Service: Failed to upload file")
	}

	_, err = ss.CertificationRepository.UpdateCertification(certif.ID.String(), &model.CertificationReq{
		PhotoLink: link})
	if err != nil {
		return *certif, errors.New("Service: Failed to update certification")
	}
	return *certif, nil
}
