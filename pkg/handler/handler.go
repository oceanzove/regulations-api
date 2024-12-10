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
		}
	}

	return router
}
