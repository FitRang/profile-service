package handlers

import (
	"errors"
	"net/http"

	"github.com/FitRang/profile-service/apperror"
	"github.com/FitRang/profile-service/domain"
	"github.com/FitRang/profile-service/model"

	"github.com/gin-gonic/gin"
)

func (ph *ProfileHandler) UpdateProfileHandler(c *gin.Context) {
	profile := model.ProfileUpdateRequest{}
	if err := c.ShouldBindBodyWithJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": apperror.CustomValidationError(&profile, err),
		})
		return
	}
	resProfile, err := ph.domain.UpdateProfile(&profile)
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
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Profile Updated Successfully",
		"data":    resProfile,
	})
}
