package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rGustave07/fbhcapi/api/model"
	"github.com/rGustave07/fbhcapi/api/model/apperrors"
)

// Used for validation
type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

// Signup handles user signup
func (h *Handler) Signup(c *gin.Context) {
	// Define a variable and we'll bind it to the incoming json body
	// expected body: {email, password}
	var req signupReq

	if ok := bindData(c, &req); !ok {
		// Return because the bind data function will handle errors
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.UserService.Signup(c, u)

	if err != nil {
		log.Printf("Failed to signup user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	// Create token pair as strings
	tokens, err := h.TokenService.NewPairFromUser(c, u, "")

	if err != nil {
		log.Printf("Failed to create tokens for user: %v\n", err.Error())

		// TODO: If tokens are failed to create, reverse creation of user

		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"tokens": tokens,
	})
}
