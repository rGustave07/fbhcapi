package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rGustave07/fbhcapi/api/model/apperrors"
)

// used to help extract validation errors
type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

// bindData is helper function, returns false if data is not bound
func bindData(c *gin.Context, req interface{}) bool {
	// Bind incoming json to struct and check for validation errors
	if err := c.ShouldBind(req); err != nil {
		log.Printf("Error binding data %+v\n", err)

		// If binding fails
		if errs, ok := err.(validator.ValidationErrors); ok {
			// Could probably extract this, will be in middleware
			var invalidArgs []invalidArgument

			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					err.Field(),
					err.Value().(string),
					err.Tag(),
					err.Param(),
				})
			}

			err := apperrors.NewBadRequestErr("Invalid request params. See invalidArgs")

			c.JSON(err.Status(), gin.H{
				"error":       err,
				"invalidArgs": invalidArgs,
			})
			return false
		}

		// Later: Add code for validating max body size
		// If we can't properly extract validation errors,
		// return an internal server error
		fallBack := apperrors.NewInternalErr()

		c.JSON(fallBack.Status(), gin.H{"error": fallBack})
		return false
	}

	return true
}
