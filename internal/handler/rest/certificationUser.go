package rest

import (
	"includemy/model"
	"includemy/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateCertificationUser(ctx *gin.Context) {
	param := model.CertificationGet{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	regis, err := r.service.CertificationUserService.CreateCertificationUser(ctx, &param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to register certification", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Success to register certification", regis)
}

func (r *Rest) DeleteCertificationUser(ctx *gin.Context) {
	certifUserID := ctx.Param("id")
	err := r.service.CertificationUserService.DeleteCertificationUser(certifUserID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete certification", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete certification", nil)
}
