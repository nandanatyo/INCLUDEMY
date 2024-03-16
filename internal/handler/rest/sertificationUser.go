package rest

import (
	"includemy/model"
	"includemy/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreatSertificationUser(ctx *gin.Context) {
	param := model.CreateSertificationUser{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	regis, err := r.service.SertificationUserService.CreateSertificationUser(&param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to register sertification", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Success to register sertification", regis)
}

func (r *Rest) DeleteSertificationUser(ctx *gin.Context) {
	sertifUserID := ctx.Param("id")
	err := r.service.SertificationUserService.DeleteSertificationUser(sertifUserID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete sertification", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete sertification", nil)
}
