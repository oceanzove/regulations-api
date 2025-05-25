package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"regulations-api/models"
	"regulations-api/pkg/usecase"
)

func (h *Handler) signIn(c *gin.Context) {
	var input models.SignInInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.Error(err.Error())
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	output, processStatus := h.usecase.SignIn(&input)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, output, processStatus)
}

func (h *Handler) refresh(c *gin.Context) {
	var input models.RefreshInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.Error(err.Error())
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	logrus.Infof("Received refresh token: %v", input.RefreshToken)

	if input.RefreshToken == "" {
		logrus.Error("Refresh token is empty")
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	output, processStatus := h.usecase.Refresh(input.RefreshToken)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, output, processStatus)
}
