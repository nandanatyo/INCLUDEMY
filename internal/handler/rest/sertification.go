package rest

import (
	"errors"
	"includemy/model"
	"includemy/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) CreateSertification(ctx *gin.Context) {
	var sertificationCre model.SertificationReq
	if err := ctx.ShouldBindJSON(&sertificationCre); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "invalid request", err)
		return
	}

	sertification, err := r.service.SertificationService.CreateSertification(&sertificationCre)

	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create sertification", err)
		return
	}
	response.Success(ctx, http.StatusCreated, "Sertification created", sertification)
}

func (r *Rest) GetSertificationByTitleOrID(ctx *gin.Context) {
	var searchParam model.SertifSearch

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
	

	sertif, err := r.service.SertificationService.GetSertificationByAny(searchParam)
	if err != nil {
		// Handle errors, such as course not found
		response.Error(ctx, http.StatusNotFound, "Failed to find sertification", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Sertification found", sertif)
}

func (r *Rest) DeleteSertification(ctx *gin.Context) {
	sertifID := ctx.Param("id")
	err := r.service.SertificationService.DeleteSertification(sertifID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete sertification", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete sertification", nil)
}

func (r *Rest) UpdateSertification(ctx *gin.Context) {
	id := ctx.Param("id")
	var sertifReq model.SertificationReq
	if err := ctx.ShouldBindJSON(&sertifReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Invalid request", err)
		return
	}

	sertif, err := r.service.SertificationService.UpdateSertification(id, &sertifReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update sertification", err)
		return
	}
	response.Success(ctx, http.StatusOK, "Sertification updated", sertif)
}

func (r *Rest) UploadSertifPhoto(ctx *gin.Context) {

	sertifID := ctx.PostForm("sertif_id")
	if sertifID == "" {
		response.Error(ctx, http.StatusBadRequest, "sertif_id is required", errors.New("sertif_id is required"))
		return
	}

	parsedSertifID, err := uuid.Parse(sertifID)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid sertifID format", err)
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	sertif, err := r.service.SertificationService.UploadSertificationFile(model.SertifPost{
		ID:   parsedSertifID,
		File: file,
	})
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to upload photo", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to upload photo", sertif)
}
