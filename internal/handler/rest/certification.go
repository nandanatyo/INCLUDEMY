package rest

import (
	"errors"
	"includemy/model"
	"includemy/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) CreateCertification(ctx *gin.Context) {
	var certificationCre model.CertificationReq
	if err := ctx.ShouldBindJSON(&certificationCre); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "invalid request", err)
		return
	}

	certification, err := r.service.CertificationService.CreateCertification(&certificationCre)

	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create certification", err)
		return
	}
	response.Success(ctx, http.StatusCreated, "Certification created", certification)
}

func (r *Rest) GetCertificationByTitleOrID(ctx *gin.Context) {
	var searchParam model.CertifSearch

	title := ctx.Query("title")
	idStr := ctx.Query("id")
	tags := ctx.Query("tags")
	dissability := ctx.Query("dissability")
	field := ctx.Query("field")

	if title != "" {
		searchParam.Title = title
	}

	if idStr != "" {
		id, err := uuid.Parse(idStr)
		if err != nil {
			response.Error(ctx, http.StatusBadRequest, "Invalid ID format", err)
			return
		}
		searchParam.ID = id
	}

	if tags != "" {
		searchParam.Tags = tags
	}

	if dissability != "" {
		searchParam.Dissability = dissability
	}

	if field != "" {
		searchParam.Field = field
	}

	certif, err := r.service.CertificationService.GetCertificationByAny(searchParam)
	if err != nil {
		// Handle errors, such as course not found
		response.Error(ctx, http.StatusNotFound, "Failed to find certification", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Certification found", certif)
}

func (r *Rest) DeleteCertification(ctx *gin.Context) {
	certifID := ctx.Param("id")
	err := r.service.CertificationService.DeleteCertification(certifID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete certification", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete certification", nil)
}

func (r *Rest) UpdateCertification(ctx *gin.Context) {
	id := ctx.Param("id")
	var certifReq model.CertificationReq
	if err := ctx.ShouldBindJSON(&certifReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Invalid request", err)
		return
	}

	certif, err := r.service.CertificationService.UpdateCertification(id, &certifReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update certification", err)
		return
	}
	response.Success(ctx, http.StatusOK, "Certification updated", certif)
}

func (r *Rest) UploadCertifPhoto(ctx *gin.Context) {

	certifID := ctx.PostForm("certif_id")
	if certifID == "" {
		response.Error(ctx, http.StatusBadRequest, "certif_id is required", errors.New("certif_id is required"))
		return
	}

	parsedCertifID, err := uuid.Parse(certifID)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid certifID format", err)
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	certif, err := r.service.CertificationService.UploadCertificationFile(model.CertifPost{
		ID:   parsedCertifID,
		File: file,
	})
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to upload photo", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to upload photo", certif)
}
