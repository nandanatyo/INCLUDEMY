package repository

import (
	"errors"
	"includemy/entity"
	"includemy/model"

	"gorm.io/gorm"
)

type ISertificationRepository interface {
	CreateSertification(sertification *entity.Sertification) (*entity.Sertification, error)
	DeleteSertification(id string) error
	GetSertiicationByName(title string) ([]*entity.Sertification, error)
	GetSertificationByID(id string) (*entity.Sertification, error)
	UpdateSertification(id string, modifySertif *model.SertificationReq) (*entity.Sertification, error)
	//update
	//delete
}

type SertificationRepository struct {
	db *gorm.DB
}

func NewSertificationRepository(db *gorm.DB) ISertificationRepository {
	return &SertificationRepository{db}
}

func (sr *SertificationRepository) CreateSertification(sertification *entity.Sertification) (*entity.Sertification, error) {
	if err := sr.db.Debug().Create(sertification).Error; err != nil {
		return nil, errors.New("Repository: Failed to create Sertification")
	}
	return sertification, nil
}

func (sr *SertificationRepository) GetSertificationByID(id string) (*entity.Sertification, error) {
	var sertif entity.Sertification
	if err := sr.db.Debug().Where("id = ?", id).First(&sertif).Error; err != nil {
		return nil, errors.New("Repository: Course not found")
	}
	return &sertif, nil
}

func (sr *SertificationRepository) GetSertiicationByName(title string) ([]*entity.Sertification, error) {
	var sertif []*entity.Sertification
	if err := sr.db.Debug().Where("title LIKE ?", "%"+title+"%").Find(&sertif).Error; err != nil {
		return nil, errors.New("Repository: Course not found")
	}
	return sertif, nil
}

func (sr *SertificationRepository) DeleteSertification(id string) error {
	if err := sr.db.Debug().Where("id = ?", id).Delete(&entity.Sertification{}).Error; err != nil {
		return errors.New("Repository: Failed to delete sertification")
	}
	return nil
}

func (sr *SertificationRepository) UpdateSertification(id string, modifySertif *model.SertificationReq) (*entity.Sertification, error) {
	var sertif entity.Sertification
	if err := sr.db.Debug().Where("id = ?", id).First(&sertif).Error; err != nil {
		return nil, err
	}

	sertifParse := parseUpdateSertif(modifySertif, &sertif)
	if err := sr.db.Model(&sertif).Save(&sertifParse).Error; err != nil {
		return nil, err
	}
	return sertifParse, nil
}

func parseUpdateSertif(modifySertif *model.SertificationReq, sertif *entity.Sertification) *entity.Sertification {
	if modifySertif.Title != "" {
		sertif.Title = modifySertif.Title
	}
	if modifySertif.About != "" {
		sertif.About = modifySertif.About
	}
	if modifySertif.Tags != "" {
		sertif.Tags = modifySertif.Tags
	}
	if modifySertif.Link != "" {
		sertif.Link = modifySertif.Link
	}
	if modifySertif.Field != "" {
		sertif.Field = modifySertif.Field
	}
	if modifySertif.Creator != "" {
		sertif.Creator = modifySertif.Creator
	}
	if modifySertif.Location != "" {
		sertif.Location = modifySertif.Location
	}
	if modifySertif.PhotoLink != "" {
		sertif.PhotoLink = modifySertif.PhotoLink
	}
	if modifySertif.Syllabus != "" {
		sertif.Syllabus = modifySertif.Syllabus
	}
	if modifySertif.Dissability != "" {
		sertif.Dissability = modifySertif.Dissability
	}
	if modifySertif.Price != 0 {
		sertif.Price = modifySertif.Price
	}

	return sertif
}
