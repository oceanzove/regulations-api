package handler

import (
	"github.com/gin-gonic/gin"
	"regulations-api/models"
	"regulations-api/pkg/usecase"
)

func (h *Handler) createRegulation(c *gin.Context) {
	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	var input *models.CreateRegulationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	processStatus := h.usecase.CreateRegulation(accountId, input)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, nil, processStatus)
}

func (h *Handler) getRegulations(c *gin.Context) {
	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	output, processStatus := h.usecase.GetRegulation(accountId)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, output, processStatus)
}

func (h *Handler) updateRegulation(c *gin.Context) {
	var input models.UpdateRegulationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	regulationID := c.Param("regulationID")
	input.ID = regulationID

	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	processStatus := h.usecase.UpdateRegulation(input, accountId)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}
}
