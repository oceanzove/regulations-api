package handler

import (
	"github.com/gin-gonic/gin"
	"regulations-api/models"
	"regulations-api/pkg/usecase"
)

func (h *Handler) getAccount(c *gin.Context) {
	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	output, processStatus := h.usecase.GetAccount(accountId)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, output, processStatus)
}

func (h *Handler) getAccountById(c *gin.Context) {
	accountId := c.Param("accountID")

	output, processStatus := h.usecase.GetAccount(accountId)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, output, processStatus)
}

func (h *Handler) getDepartments(c *gin.Context) {
	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	output, processStatus := h.usecase.GetDepartments(accountId)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, output, processStatus)
}
func (h *Handler) getDepartmentById(c *gin.Context) {
	departmentId := c.Param("departmentID")

	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	output, processStatus := h.usecase.GetDepartmentById(accountId, departmentId)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, output, processStatus)
}

func (h *Handler) getPositions(c *gin.Context) {
	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	output, processStatus := h.usecase.GetPositions(accountId)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, output, processStatus)
}

func (h *Handler) getPositionsByDepartment(c *gin.Context) {
	departmentId := c.Param("departmentID")

	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	positions, processStatus := h.usecase.GetPositionsByDepartment(accountId, departmentId)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}
	h.sendResponseSuccess(c, positions, usecase.Success)
}

func (h *Handler) createEmployee(c *gin.Context) {
	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	var input *models.CreateEmployeeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	employeeStatus := h.usecase.CreateEmployee(accountId, input)
	if employeeStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, employeeStatus)
		return
	}

	h.sendResponseSuccess(c, nil, employeeStatus)
}

func (h *Handler) updateEmployee(c *gin.Context) {
	var input *models.Employee
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	employeeID := c.Param("processID")
	input.ID = employeeID

	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	processStatus := h.usecase.UpdateEmployee(input)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}
}

func (h *Handler) updateAccount(c *gin.Context) {
	var input *models.Account
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	accountID := c.Param("accountID")
	input.ID = accountID

	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	processStatus := h.usecase.UpdateAccount(input)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}
}

func (h *Handler) updateEmployeePosition(c *gin.Context) {
	var input *models.UpdateEmployeePosition
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	EmployeeID := c.Param("EmployeeID")
	input.EmployeeID = EmployeeID

	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	processStatus := h.usecase.UpdateEmployeePosition(input)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}
}

func (h *Handler) updateEmployeeDepartment(c *gin.Context) {
	var input *models.UpdateEmployeeDepartment
	if err := c.ShouldBindJSON(&input); err != nil {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	EmployeeID := c.Param("EmployeeID")
	input.EmployeeID = EmployeeID

	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	processStatus := h.usecase.UpdateEmployeeDepartment(input)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}
}

func (h *Handler) getEmployees(c *gin.Context) {
	accountId := c.GetString(gin.AuthUserKey)
	if accountId == "" {
		h.sendResponseSuccess(c, nil, usecase.InternalServerError)
		return
	}

	output, processStatus := h.usecase.GetEmployees(accountId)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, output, processStatus)
}

func (h *Handler) getEmployeeById(c *gin.Context) {
	employeeId := c.Param("employeeID")

	output, processStatus := h.usecase.GetEmployeeById(employeeId)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, output, processStatus)
}

func (h *Handler) getDepartmentByEmployeeId(c *gin.Context) {
	employeeId := c.Param("employeeID")

	output, processStatus := h.usecase.GetDepartmentByEmployeeId(employeeId)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, output, processStatus)
}

func (h *Handler) getPositionByEmployeeId(c *gin.Context) {
	employeeId := c.Param("employeeID")

	output, processStatus := h.usecase.GetPositionByEmployeeId(employeeId)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, output, processStatus)
}
