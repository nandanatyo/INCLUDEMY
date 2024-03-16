package rest

import (
	"includemy/model"
	"includemy/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
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
