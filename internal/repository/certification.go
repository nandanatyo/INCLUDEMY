package repository

import (
	"errors"
	"includemy/entity"
	"includemy/model"

	"gorm.io/gorm"
)

type ICertificationRepository interface {
	CreateCertification(certification *entity.Certification) (*entity.Certification, error)
	DeleteCertification(id string) error
	GetCertificationByName(title string) ([]*entity.Certification, error)
	GetCertificationByID(id string) (*entity.Certification, error)
	UpdateCertification(id string, modifyCertif *model.CertificationReq) (*entity.Certification, error)
	GetCertificationByDissability(dissability string) ([]*entity.Certification, error)
	GetCertificationByField(field string) ([]*entity.Certification, error)
	GetCertificationByTags(tags string) ([]*entity.Certification, error)
}

type CertificationRepository struct {
	db *gorm.DB
}

func NewCertificationRepository(db *gorm.DB) ICertificationRepository {
	return &CertificationRepository{db}
}

func (sr *CertificationRepository) CreateCertification(certification *entity.Certification) (*entity.Certification, error) {
	if err := sr.db.Debug().Create(certification).Error; err != nil {
		return nil, errors.New("Repository: Failed to create certification")
	}
	return certification, nil
}

func (sr *CertificationRepository) GetCertificationByID(id string) (*entity.Certification, error) {
	var certif entity.Certification
	if err := sr.db.Debug().Where("id = ?", id).First(&certif).Error; err != nil {
		return nil, errors.New("Repository: Certification not found")
	}
	return &certif, nil
}

func (sr *CertificationRepository) GetCertificationByName(title string) ([]*entity.Certification, error) {
	var certif []*entity.Certification
	if err := sr.db.Debug().Where("title LIKE ?", "%"+title+"%").Find(&certif).Error; err != nil {
		return nil, errors.New("Repository: Certification not found")
	}
	return certif, nil
}

func (sr *CertificationRepository) GetCertificationByTags(tags string) ([]*entity.Certification, error) {
	var certif []*entity.Certification
	if err := sr.db.Debug().Where("tags LIKE ?", "%"+tags+"%").Find(&certif).Error; err != nil {
		return nil, errors.New("Repository: Certification not found")
	}
	return certif, nil
}

func (sr *CertificationRepository) GetCertificationByField(field string) ([]*entity.Certification, error) {
	var certif []*entity.Certification
	if err := sr.db.Debug().Where("field LIKE ?", "%"+field+"%").Find(&certif).Error; err != nil {
		return nil, errors.New("Repository: Certification not found")
	}
	return certif, nil
}

func (sr *CertificationRepository) GetCertificationByDissability(dissability string) ([]*entity.Certification, error) {
	var certif []*entity.Certification
	if err := sr.db.Debug().Where("dissability LIKE ?", "%"+dissability+"%").Find(&certif).Error; err != nil {
		return nil, errors.New("Repository: Certification not found")
	}
	return certif, nil
}

func (sr *CertificationRepository) DeleteCertification(id string) error {
	if err := sr.db.Debug().Where("id = ?", id).Delete(&entity.Certification{}).Error; err != nil {
		return errors.New("Repository: Failed to delete certification")
	}
	return nil
}

func (sr *CertificationRepository) UpdateCertification(id string, modifyCertif *model.CertificationReq) (*entity.Certification, error) {
	var certif entity.Certification
	if err := sr.db.Debug().Where("id = ?", id).First(&certif).Error; err != nil {
		return nil, err
	}

	certifParse := parseUpdateCertif(modifyCertif, &certif)
	if err := sr.db.Model(&certif).Save(&certifParse).Error; err != nil {
		return nil, err
	}
	return certifParse, nil
}

func parseUpdateCertif(modifyCertif *model.CertificationReq, certif *entity.Certification) *entity.Certification {
	if modifyCertif.Title != "" {
		certif.Title = modifyCertif.Title
	}
	if modifyCertif.About != "" {
		certif.About = modifyCertif.About
	}
	if modifyCertif.Tags != "" {
		certif.Tags = modifyCertif.Tags
	}
	if modifyCertif.Link != "" {
		certif.Link = modifyCertif.Link
	}
	if modifyCertif.Field != "" {
		certif.Field = modifyCertif.Field
	}
	if modifyCertif.Creator != "" {
		certif.Creator = modifyCertif.Creator
	}
	if modifyCertif.Location != "" {
		certif.Location = modifyCertif.Location
	}
	if modifyCertif.PhotoLink != "" {
		certif.PhotoLink = modifyCertif.PhotoLink
	}
	if modifyCertif.Syllabus != "" {
		certif.Syllabus = modifyCertif.Syllabus
	}
	if modifyCertif.Dissability != "" {
		certif.Dissability = modifyCertif.Dissability
	}
	if modifyCertif.Price != 0 {
		certif.Price = modifyCertif.Price
	}

	return certif
}
