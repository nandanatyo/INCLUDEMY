package repository

import (
	"errors"
	"includemy/entity"

	"gorm.io/gorm"
)

type IInvoiceRepository interface {
	CreateInvoice(invoice *entity.Invoice) (*entity.Invoice, error)
	UpdateInvoice(status string, orderID string) (*entity.Invoice, error)
	GetInvoiceByID(id string) (*entity.Invoice, error)
}

type InvoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) IInvoiceRepository {
	return &InvoiceRepository{db}
}

func (u *InvoiceRepository) CreateInvoice(invoice *entity.Invoice) (*entity.Invoice, error) {
	if err := u.db.Debug().Create(invoice).Error; err != nil {
		return nil, errors.New("Repository: Failed to create invoice")
	}
	return invoice, nil
}

func (u *InvoiceRepository) GetInvoiceByID(id string) (*entity.Invoice, error) {
	var invoice entity.Invoice
	if err := u.db.Debug().Where("order_id = ?", id).First(&invoice).Error; err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (u *InvoiceRepository) UpdateInvoice(status string, orderID string) (*entity.Invoice, error) {
	var UpdateInvoice entity.Invoice
	if err := u.db.Debug().Where("order_id = ?", orderID).First(&UpdateInvoice).Error; err != nil {
		return nil, err
	}

	updated := entity.Invoice{
		Status: status,
	}
	if err := u.db.Debug().Model(&UpdateInvoice).Where("order_id = ?", orderID).Updates(&updated).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}
