package middleware

import (
	"errors"
	"includemy/model"
	"includemy/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *middleware) AuthenticateUser(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")
	if bearer == "" {
		response.Error(ctx, http.StatusUnauthorized, "empty token", errors.New(""))
		ctx.Abort()
		return
	}

	//cth: Bearer token
	token := strings.Split(bearer, " ")[1]

	userId, err := m.jwtAuth.ValidateToken(token)

	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "invalid token", err)
		ctx.Abort()
		return
	}

	user, err := m.service.UserService.GetUserParam(model.UserParam{ID: userId})
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "user not found", err)
		ctx.Abort()
		return
	}

	ctx.Set("user", user)

	ctx.Next()
}
