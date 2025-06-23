package handler

import (
	"github.com/gin-gonic/gin"
	"regulations-api/models"
	"regulations-api/pkg/usecase"
)

func (h *Handler) createProcess(c *gin.Context) {
	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	var input *models.CreateProcessInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	processStatus := h.usecase.CreateProcess(accountId, input)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, nil, processStatus)
}

func (h *Handler) getProcesses(c *gin.Context) {
	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	output, processStatus := h.usecase.GetProcesses(accountId)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, output, processStatus)
}

func (h *Handler) getProcessByID(c *gin.Context) {
	processID := c.Param("processID")

	accountID := c.GetString(gin.AuthUserKey)
	if accountID == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	process, processStatus := h.usecase.GetProcessByID(accountID, processID)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}
	h.sendResponseSuccess(c, process, usecase.Success)
}

func (h *Handler) updateProcessById(c *gin.Context) {
	var input *models.UpdateProcessInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	processID := c.Param("processID")
	input.ID = processID

	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	processStatus := h.usecase.UpdateProcess(input)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}
}

func (h *Handler) updateStepById(c *gin.Context) {
	var input *models.Step
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	stepId := c.Param("stepID")
	input.ID = stepId

	processStatus := h.usecase.UpdateStepById(input)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}
}

func (h *Handler) deleteProcessById(c *gin.Context) {
	processID := c.Param("processID")

	processStatus := h.usecase.DeleteProcessById(processID)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}
}
func (h *Handler) deleteStepById(c *gin.Context) {
	stepID := c.Param("stepID")

	processStatus := h.usecase.DeleteStepById(stepID)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}
}

func (h *Handler) linkRegulationToProcess(c *gin.Context) {
	accountID := c.GetString(gin.AuthUserKey)
	if accountID == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	processID := c.Param("processID")
	var input struct {
		RegulationID string `json:"regulation_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil || input.RegulationID == "" {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	linkStatus := h.usecase.LinkRegulationToProcess(processID, input.RegulationID)
	if linkStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, linkStatus)
		return
	}

	h.sendResponseSuccess(c, nil, usecase.Success)
}

func (h *Handler) unlinkRegulationToProcess(c *gin.Context) {
	accountID := c.GetString(gin.AuthUserKey)
	if accountID == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	processID := c.Param("processID")
	var input struct {
		RegulationID string `json:"regulation_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil || input.RegulationID == "" {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	linkStatus := h.usecase.UnlinkRegulationToProcess(processID, input.RegulationID)
	if linkStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, linkStatus)
		return
	}

	h.sendResponseSuccess(c, nil, usecase.Success)
}

func (h *Handler) getRegulationsByProcess(c *gin.Context) {
	processID := c.Param("processID")
	accountID := c.GetString(gin.AuthUserKey)
	if accountID == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	regulations, err := h.usecase.GetRegulationsByProcess(processID)
	if err != nil {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	h.sendResponseSuccess(c, regulations, usecase.Success)
}

func (h *Handler) getStepsByProcess(c *gin.Context) {
	processID := c.Param("processID")
	accountID := c.GetString(gin.AuthUserKey)
	if accountID == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	steps, err := h.usecase.GetStepsByProcess(processID)
	if err != nil {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	h.sendResponseSuccess(c, steps, usecase.Success)
}

func (h *Handler) createStep(c *gin.Context) {
	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	var input *models.Step
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	stepStatus := h.usecase.CreateStep(input)
	if stepStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, stepStatus)
		return
	}

	h.sendResponseSuccess(c, nil, stepStatus)
}
