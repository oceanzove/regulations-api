package handler

import (
	mobile "github.com/floresj/go-contrib-mobile"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"regulations-api/models"
	"regulations-api/pkg/service"
	"regulations-api/pkg/usecase"
	"regulations-api/pkg/utils"
	"strings"
	"time"
)

type Handler struct {
	usecase *usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase, services *service.Service) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) InitHTTPRoutes(config *models.ServerConfig) *gin.Engine {
	router := gin.Default()

	allowOrigins := strings.Split("http://localhost:5173", ",")

	router.Use(cors.New(cors.Config{
		AllowOrigins: allowOrigins,
		AllowMethods: []string{http.MethodPut, http.MethodPost, http.MethodGet, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{"Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin",
			utils.HeaderAuthorization, utils.HeaderClientRequestId},
		ExposeHeaders: []string{"Content-Length", utils.HeaderTimestamp,
			utils.HeaderClientRequestId, utils.HeaderRequestId},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(mobile.Resolver())

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-in", h.signIn)
			auth.POST("/refresh", h.refresh)
		}
		organization := api.Group("/organization", h.UserIdentityMiddleware)
		{
			employee := organization.Group("/employee")
			{
				employee.POST("", h.createEmployee)
				employee.GET("", h.getEmployees)
				employee.GET("/:employeeID", h.getEmployeeById)
				employee.PUT("/:employeeID", h.updateEmployee)

				employee.DELETE("/:employeeID", h.deleteEmployeeById)

				department := employee.Group("/department")
				{
					department.GET("/:employeeID", h.getDepartmentByEmployeeId)

					department.GET("", h.getEmployeeDepartment)

					department.PUT("/:employeeID", h.updateEmployeeDepartment)
				}

				position := employee.Group("/position")
				{
					position.GET("/:employeeID", h.getPositionByEmployeeId)

					position.GET("", h.getEmployeePosition)

					position.PUT("/:employeeID", h.updateEmployeePosition)
				}

				account := employee.Group("/account")
				{
					account.GET("", h.getAccount)
					account.GET("/:accountID", h.getAccountById)
					account.PUT("/:accountID", h.updateAccount)
				}
			}
			department := organization.Group("/department")
			{
				department.GET("/:departmentID", h.getDepartmentById)

				department.GET("/position", h.getDepartmentPosition)

				department.GET("", h.getDepartments)

				department.POST("", h.createDepartment)

				department.PUT("/:departmentID", h.updateDepartmentById)
			}
			position := organization.Group("/position")
			{
				position.GET("", h.getPositions)
				position.GET("/:departmentID", h.getPositionsByDepartment)

				position.POST("", h.createPosition)
				position.PUT("/:positionID", h.updatePositionById)
			}
		}
		regulation := api.Group("/regulation", h.UserIdentityMiddleware)
		{
			regulation.GET("", h.getRegulations)
			regulation.GET("/:regulationID", h.getRegulationByID)

			regulation.PUT("/:regulationID", h.updateRegulation)
			regulation.POST("", h.createRegulation)

			regulation.POST("/section", h.createSection)
			regulation.GET("/section", h.getSections)

			regulation.POST("/:regulationID/section/link", h.linkSectionToRegulation)
			regulation.POST("/:regulationID/section/unlink", h.unlinkSectionToRegulation)

			regulation.GET("/:regulationID/section", h.getSectionById)

			regulation.DELETE("/:regulationID", h.deleteRegulationById)
		}
		process := api.Group("/process", h.UserIdentityMiddleware)
		{
			process.GET("", h.getProcesses)
			process.GET("/:processID", h.getProcessByID)
			process.PUT("/:processID", h.updateProcessById)
			process.POST("", h.createProcess)

			process.POST("/step", h.createStep)
			process.PUT("/step/:stepID", h.updateStepById)
			process.GET("/:processID/step", h.getStepsByProcess)

			process.POST("/:processID/regulation/link", h.linkRegulationToProcess)
			process.POST("/:processID/regulation/unlink", h.unlinkRegulationToProcess)
			process.GET("/:processID/regulation", h.getRegulationsByProcess)

			process.DELETE("/:processID", h.deleteProcessById)
			process.DELETE("/step/:stepID", h.deleteStepById)
		}
	}

	return router
}
