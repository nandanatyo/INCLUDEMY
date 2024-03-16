package middleware

import (
	"errors"
	"includemy/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *middleware) OnlyAdmin(ctx *gin.Context) {
	user, err := m.jwtAuth.GetLogin(ctx)
	if err != nil {
		response.Error(ctx, http.StatusForbidden, "invalid token", err)
		ctx.Abort()
		return
	}

	if user.Role != 1 {
		response.Error(ctx, http.StatusForbidden, "The endpoint can't be access", errors.New(""))
		ctx.Abort()
		return
	}

	ctx.Next()

}
