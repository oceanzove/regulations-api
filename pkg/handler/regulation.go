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

func (h *Handler) createSection(c *gin.Context) {
	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	var input *models.CreateSectionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	processStatus := h.usecase.CreateSection(accountId, input)
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

func (h *Handler) getSections(c *gin.Context) {
	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	output, sectionStatus := h.usecase.GetSections(accountId)
	if sectionStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, sectionStatus)
		return
	}

	h.sendResponseSuccess(c, output, sectionStatus)
}

func (h *Handler) getRegulationByID(c *gin.Context) {
	regulationID := c.Param("regulationID")

	accountID := c.GetString(gin.AuthUserKey)
	if accountID == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	regulation, regulationStatus := h.usecase.GetRegulationByID(accountID, regulationID)
	if regulationStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}
	h.sendResponseSuccess(c, regulation, usecase.Success)
}

func (h *Handler) getSectionById(c *gin.Context) {
	regulationID := c.Param("regulationID")

	regulation, regulationStatus := h.usecase.GetSectionById(regulationID)
	if regulationStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}
	h.sendResponseSuccess(c, regulation, usecase.Success)
}

func (h *Handler) deleteRegulationById(c *gin.Context) {
	regulationID := c.Param("regulationID")

	processStatus := h.usecase.DeleteRegulationById(regulationID)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}
}

func (h *Handler) linkSectionToRegulation(c *gin.Context) {
	accountID := c.GetString(gin.AuthUserKey)
	if accountID == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	regulationID := c.Param("regulationID")
	var input *models.LinkSectionToRegulation
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}
	input.RegulationID = regulationID

	linkStatus := h.usecase.LinkSectionToRegulation(input)
	if linkStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, linkStatus)
		return
	}

	h.sendResponseSuccess(c, nil, usecase.Success)
}

func (h *Handler) unlinkSectionToRegulation(c *gin.Context) {
	accountID := c.GetString(gin.AuthUserKey)
	if accountID == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	regulationID := c.Param("regulationID")
	var input struct {
		SectionID string `json:"sectionId" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil || input.SectionID == "" {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	linkStatus := h.usecase.UnlinkSectionToRegulation(regulationID, input.SectionID)
	if linkStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, linkStatus)
		return
	}

	h.sendResponseSuccess(c, nil, usecase.Success)
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
