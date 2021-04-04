package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rGustave07/fbhcapi/api/model"
	"github.com/rGustave07/fbhcapi/api/model/apperrors"
)

// Me handler(controller) calls services for getting user details
func (h *Handler) Me(c *gin.Context) {
	// A *model.User will eventually be added to context via middleware
	user, exists := c.Get("user")

	if !exists {
		log.Printf("Unable to extract user from request context for unknown reasons %v\n", c)
		err := apperrors.NewInternalErr()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})

		return
	}

	uid := user.(*model.User).UID

	u, err := h.UserService.Get(c, uid)

	if err != nil {
		log.Printf("Unable to find user: %v\n%v", uid, err)
		e := apperrors.NewNotFoundErr("user", uid.String())

		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
