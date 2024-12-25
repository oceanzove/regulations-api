package handler

import (
	"github.com/gin-gonic/gin"
	"regulations-api/models"
	"regulations-api/pkg/usecase"
)

func (h *Handler) createProcess(c *gin.Context) {
	email := c.GetString(gin.AuthUserKey)
	if email == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	output, processStatus := h.usecase.CreateProcess(email)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, output, processStatus)
}

func (h *Handler) getProcesses(c *gin.Context) {
	email := c.GetString(gin.AuthUserKey)
	if email == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	output, processStatus := h.usecase.GetProcesses(email)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, output, processStatus)
}

func (h *Handler) updateProcess(c *gin.Context) {
	var input models.UpdateProcessInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	processID := c.Param("processID")
	input.ID = processID

	email := c.GetString(gin.AuthUserKey)
	if email == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	processStatus := h.usecase.UpdateProcess(input, email)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}
}
