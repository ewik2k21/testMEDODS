package handler

import (
	"net/http"
	testmedods "testMEDODS"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getPairTokens(ctx *gin.Context) {
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

func (h *Handler) refreshPairTokens(ctx *gin.Context) {

}
