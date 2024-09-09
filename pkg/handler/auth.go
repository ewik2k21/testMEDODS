package handler

import (
	"net/http"
	testmedods "testMEDODS"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addUser(ctx *gin.Context) {
	var input testmedods.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:email binding:"required"`
	Password string `json:password binding:"required"`
}

func (h *Handler) getPairTokens(ctx *gin.Context) {
	var input signInInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, err := h.services.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	refreshToken, err := h.services.NewRefreshToken()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func (h *Handler) refreshPairTokens(ctx *gin.Context) {

}
