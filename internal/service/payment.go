package service

import (
	"includemy/entity"
	"includemy/internal/repository"
	"includemy/model"
	"includemy/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type IPaymentService interface {
	GetPaymentCourse(ctx *gin.Context, param *model.PaymentBind) (model.PaymentResponse, error)
	CallbackCourse(notificationPayload map[string]interface{})
	GetPaymentSertif(ctx *gin.Context, param *model.PaymentBind) (model.PaymentResponse, error)
	CallbackSertif(notificationPayload map[string]interface{})
}

type PaymentService struct {
	Invoice repository.IInvoiceRepository
	User    repository.IUserRepository
	Course  repository.ICourseRepository
	Sertif  repository.ISertificationRepository
	jwt     jwt.Interface
}

func NewPaymentService(invoice repository.IInvoiceRepository, user repository.IUserRepository, course repository.ICourseRepository, sertif repository.ISertificationRepository, jwt jwt.Interface) *PaymentService {
	return &PaymentService{
		Invoice: invoice,
		User:    user,
		Course:  course,
		Sertif:  sertif,
		jwt:     jwt,
	}
}

func (p *PaymentService) GetPaymentCourse(ctx *gin.Context, param *model.PaymentBind) (model.PaymentResponse, error) {
	user, err := p.jwt.GetLogin(ctx)
	if err != nil {
		return model.PaymentResponse{}, err
	}

	course, err := p.Course.GetCourseByID(param.ItemID.String())
	if err != nil {
		return model.PaymentResponse{}, err
	}

	payReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  uuid.New().String(),
			GrossAmt: course.Price,
		},
		Expiry: &snap.ExpiryDetails{
			Duration: 15,
			Unit:     "minute",
		},
	}

	resp, err := snap.CreateTransaction(payReq)
	_, err = p.Invoice.CreateInvoice(&entity.Invoice{
		OrderID:          payReq.TransactionDetails.OrderID,
		UserID:           user.ID.String(),
		CourseorSertifID: param.ItemID.String(),
		Status:           "pending",
	})
	if err != nil {
		return model.PaymentResponse{}, err
	}

	result := model.PaymentResponse{
		Token:   resp.Token,
		SnapUrl: resp.RedirectURL,
	}
	return result, nil

}

func (p *PaymentService) CallbackCourse(notificationPayload map[string]interface{}) {
	orderID := notificationPayload["order_id"]
	transactionStatus := notificationPayload["transaction_status"]
	fraudStatus := notificationPayload["fraud_status"]

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			// TODO set transaction status on your database to 'challenge'
			p.Invoice.UpdateInvoice("challenge", orderID.(string))
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
		} else if fraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			p.Invoice.UpdateInvoice("success", orderID.(string))
		}
	} else if transactionStatus == "settlement" {
		// TODO set transaction status on your databaase to 'success'
		p.Invoice.UpdateInvoice("success", orderID.(string))
	} else if transactionStatus == "deny" {
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		// TODO set transaction status on your databaase to 'failure'
		p.Invoice.UpdateInvoice("failure", orderID.(string))
	} else if transactionStatus == "pending" {
		// TODO set transaction status on your databaase to 'pending' / waiting payment
		p.Invoice.UpdateInvoice("pending", orderID.(string))
	}
}

func (p *PaymentService) GetPaymentSertif(ctx *gin.Context, param *model.PaymentBind) (model.PaymentResponse, error) {
	user, err := p.jwt.GetLogin(ctx)
	if err != nil {
		return model.PaymentResponse{}, err
	}

	sertif, err := p.Sertif.GetSertificationByID(param.ItemID.String())
	if err != nil {
		return model.PaymentResponse{}, err
	}

	payReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  uuid.New().String(),
			GrossAmt: int64(sertif.Price),
		},
		Expiry: &snap.ExpiryDetails{
			Duration: 15,
			Unit:     "minute",
		},
	}

	resp, err := snap.CreateTransaction(payReq)
	_, err = p.Invoice.CreateInvoice(&entity.Invoice{
		OrderID:          payReq.TransactionDetails.OrderID,
		UserID:           user.ID.String(),
		CourseorSertifID: param.ItemID.String(),
		Status:           "pending",
	})
	if err != nil {
		return model.PaymentResponse{}, err
	}

	result := model.PaymentResponse{
		Token:   resp.Token,
		SnapUrl: resp.RedirectURL,
	}
	return result, nil

}

func (p *PaymentService) CallbackSertif(notificationPayload map[string]interface{}) {
	orderID := notificationPayload["order_id"]
	transactionStatus := notificationPayload["transaction_status"]
	fraudStatus := notificationPayload["fraud_status"]

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			// TODO set transaction status on your database to 'challenge'
			p.Invoice.UpdateInvoice("challenge", orderID.(string))
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
		} else if fraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			p.Invoice.UpdateInvoice("success", orderID.(string))
		}
	} else if transactionStatus == "settlement" {
		// TODO set transaction status on your databaase to 'success'
		p.Invoice.UpdateInvoice("success", orderID.(string))
	} else if transactionStatus == "deny" {
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		// TODO set transaction status on your databaase to 'failure'
		p.Invoice.UpdateInvoice("failure", orderID.(string))
	} else if transactionStatus == "pending" {
		// TODO set transaction status on your databaase to 'pending' / waiting payment
		p.Invoice.UpdateInvoice("pending", orderID.(string))
	}
}
