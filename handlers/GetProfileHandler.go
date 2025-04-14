package handlers

import (
	"errors"
	"net/http"

	"github.com/FitRang/profile-service/apperror"
	"github.com/FitRang/profile-service/domain"
	"github.com/gin-gonic/gin"
)

func (ph *ProfileHandler) GetProfileHandler(c *gin.Context) {
	type RequestParams struct {
		ID string `json:"id" form:"id" validate:"required,uuid"`
	}
	reqParams := RequestParams{
		ID: c.Param("id"),
	}
	if err := validate.Struct(reqParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": apperror.CustomValidationError(&reqParams, err),
		})
		return
	}
	profile, err := ph.domain.GetProfile(reqParams.ID)

	if err != nil {
		if errors.Is(err, domain.ErrProfileNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"profile": profile,
	})
}
