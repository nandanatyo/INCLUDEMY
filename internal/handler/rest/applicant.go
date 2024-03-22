package rest

import (
	"errors"
	"includemy/entity"
	"includemy/model"
	"includemy/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateApplicant(ctx *gin.Context) {
	param := model.ApplicantReq{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to bind input", err)

		return
	}

	appli, err := r.service.ApplicantService.CreateApplicantService(ctx, &param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to make application", err)
		return
	}
	response.Success(ctx, http.StatusCreated, "Success to make application", appli)
}

func (r *Rest) DeleteApplication(ctx *gin.Context) {
	appID := ctx.Param("id")
	err := r.service.ApplicantService.DeleteApplication(appID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete application", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete application", nil)
}

func (r *Rest) UploadApplicantFile(ctx *gin.Context) {
	appID := ctx.PostForm("applicant_id")
	if appID == "" {
		response.Error(ctx, http.StatusBadRequest, "applicant_id is required", errors.New("applicant_id is required"))
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	app, err := r.service.ApplicantService.UploadApplicantFile(&entity.ParamAppFile{
		AppID: appID,
		File:  file,
	})

	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to upload photo", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to upload photo", app)
}
