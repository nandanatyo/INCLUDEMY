package rest

import (
	"errors"
	"includemy/entity"
	"includemy/model"
	"includemy/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) CreateJob(ctx *gin.Context) {
	var jobCre model.JobReq
	if err := ctx.ShouldBindJSON(&jobCre); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "invalid request", err)
		return
	}

	job, err := r.service.JobService.CreateJob(&jobCre)

	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create job", err)
		return
	}
	response.Success(ctx, http.StatusCreated, "Job created", job)
}

func (r *Rest) DeleteJob(ctx *gin.Context) {
	jobID := ctx.Param("id")
	err := r.service.SertificationService.DeleteSertification(jobID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete job", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete job", nil)
}

func (r *Rest) UpdateJob(ctx *gin.Context) {
	id := ctx.Param("id")
	var jobReq model.JobReq
	if err := ctx.ShouldBindJSON(&jobReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Invalid request", err)
		return
	}

	job, err := r.service.JobService.UpdateJob(id, &jobReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update job", err)
		return
	}
	response.Success(ctx, http.StatusOK, "Job updated", job)
}

func (r *Rest) GetJobByTitleOrID(ctx *gin.Context) {
	var jobParam model.JobSearch

	title := ctx.Query("title")
	idStr := ctx.Query("id")

	if title != "" {
		jobParam.Title = title
	}

	if idStr != "" {
		id, err := uuid.Parse(idStr)
		if err != nil {
			response.Error(ctx, http.StatusBadRequest, "Invalid ID format", err)
			return
		}
		jobParam.ID = id
	}

	job, err := r.service.JobService.GetJobByTitleOrID(jobParam)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to find job", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Job found", job)
}

func (r *Rest) UploadJobFile(ctx *gin.Context) {
	jobID := ctx.PostForm("job_id")
	if jobID == "" {
		response.Error(ctx, http.StatusBadRequest, "job_id is required", errors.New("job_id is required"))
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	job, err := r.service.JobService.UploadJobFile(&entity.ParamJobFile{
		JobID: jobID,
		File:  file,
	})

	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to upload photo", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to upload photo", job)
}
