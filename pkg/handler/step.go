package handler

import (
	"github.com/gin-gonic/gin"
	"regulations-api/models"
	"regulations-api/pkg/usecase"
)

func (h *Handler) createSteps(c *gin.Context) {
	email := c.GetString(gin.AuthUserKey)
	if email == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	var input *models.CreateStepsInput
	if err := c.ShouldBindJSON(&input); err != nil || len(input.Steps) == 0 {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	status := h.usecase.CreateSteps(input)
	if status != usecase.Success {
		h.sendResponseSuccess(c, nil, status)
		return
	}

	h.sendResponseSuccess(c, nil, status)
}
