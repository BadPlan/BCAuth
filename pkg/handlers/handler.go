package handlers

import (
	"BCAuth/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services services.Service
}

var (
	BindJsonError   = map[string]string{"message": "could not bind model"}
	BindParamsError = map[string]string{"message": "could not bind params"}
	BindIdError     = map[string]string{"message": "could not bind ID"}
)

func CORSMiddleware() gin.HandlerFunc {
	allowList := map[string]bool{
		"http://127.0.0.1:8080": true,
		"http://127.0.0.1":      true,
	}
	return func(c *gin.Context) {
		if origin := c.Request.Header.Get("Origin"); allowList[origin] {
			c.Header("Access-Control-Allow-Origin", origin)
		}
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Cookie")
		c.Header("Access-Control-Allow-Methods", "POST, PATCH, OPTIONS, GET, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	engine := gin.New()
	engine.Use(CORSMiddleware())

	api := engine.Group("/api/v1")
	{
		api.POST("/sign-up", h.Register)
		api.POST("/sign-in", h.SignIn)
		users := api.Group("/users", h.ValidateToken)
		{
			users.GET("/:id", h.UserInfo)
			users.GET("/", h.BrowseUsers)
			users.PATCH("/:id", h.UpdateUser)
			users.DELETE("/:id", h.DeleteUser)
		}
		roles := api.Group("/roles", h.ValidateToken)
		{
			roles.GET("/:id", h.RoleInfo)
			roles.GET("/", h.BrowseRole)
			roles.POST("/", h.CreateRole)
			roles.PATCH("/:id", h.UpdateRole)
			roles.DELETE("/:id", h.DeleteRole)
		}
		usersRoles := api.Group("/users-roles", h.ValidateToken)
		{
			usersRoles.GET("/user/:user_id", h.RolesByUser)
			usersRoles.POST("/", h.AddRole)
		}
	}

	return engine
}

func HandlerInit(service *services.Service) *Handler {
	return &Handler{
		services: *service,
	}
}

func (h *Handler) Greeting(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]string{
		"message": "Hello",
	})
}
