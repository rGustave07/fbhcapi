package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rGustave07/fbhcapi/model"
)

// Handler struct holds required services for the Handler to function
type Handler struct {
	UserService model.UserService
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R           *gin.Engine
	UserService model.UserService
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly (Via pointer) with a reference to the gin Engine
func NewHandler(c *Config) {
	h := &Handler{
		UserService: c.UserService,
	} /* Handler gets injected with configurations */

	// Create route group
	userGroupRoutes := c.R.Group(os.Getenv("USER_API_BASE_URL"))

	// TODO: Implement the following route/handlers
	// playerGroupRoutes := c.R.Group("/api/players")

	userGroupRoutes.GET("/me", h.Me)
	userGroupRoutes.POST("/signup", h.Signup)
	userGroupRoutes.POST("/signin", h.Signin)
	userGroupRoutes.POST("/signout", h.Signout)
	userGroupRoutes.POST("/tokens", h.Tokens)
}

func (h *Handler) Signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Handler": "Signup",
	})
}

func (h *Handler) Signin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Handler": "Signin",
	})
}

func (h *Handler) Signout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Handler": "Signout",
	})
}

func (h *Handler) Tokens(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Handler": "Tokens",
	})
}
