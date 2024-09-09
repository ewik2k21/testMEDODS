package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) tokenValidation(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty header")
		return
	}

	headerSplit := strings.Split(header, " ")
	if len(headerSplit) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, "invalid header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerSplit[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	ctx.Set(userCtx, userId)

}
