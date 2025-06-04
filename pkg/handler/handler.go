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
		regulation := api.Group("/regulation", h.UserIdentityMiddleware)
		{
			regulation.GET("", h.getRegulations)
			regulation.GET("/:regulationID", h.getRegulationByID)

			regulation.PUT("/:regulationID", h.updateRegulation)
			regulation.POST("", h.createRegulation)
		}
		process := api.Group("/process", h.UserIdentityMiddleware)
		{
			process.GET("", h.getProcesses)
			process.GET("/:processID", h.getProcessByID)
			process.PUT("/:processID", h.updateProcess)
			process.POST("", h.createProcess)

			process.POST("/step", h.createStep)
			process.GET("/:processID/step", h.getStepsByProcess)

			process.POST("/:processID/regulation", h.linkRegulationToProcess)
			process.GET("/:processID/regulation", h.getRegulationsByProcess)
		}
	}

	return router
}
