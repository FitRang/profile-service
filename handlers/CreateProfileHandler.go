package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/FitRang/profile-service/domain"
	"github.com/FitRang/profile-service/model"
	"github.com/FitRang/profile-service/apperror"
	"github.com/gin-gonic/gin"
)

func (ph *ProfileHandler) CreateProfileHandler(c *gin.Context) {
	profile := model.ProfileCreateRequest{}
	if err := c.ShouldBindBodyWithJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": apperror.CustomValidationError(&profile, err),
		})
		return
	} 
	err := ph.domain.CreateProfile(&profile)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrIDAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{
				"message":"A profile with this ID already exists",
			})
        case errors.Is(err, domain.ErrEmailAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{
				"message":"A profile with this email already exists",
			})
		case errors.Is(err, domain.ErrPhoneNumberAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{
				"message":"A profile with this phone number already exists",
			})
		default:
			log.Printf("Unexpected error while creating profile: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message":"Internal server error",
				"code":"INTERNAL_ERROR",
			})
		}
		return
	} 
	c.JSON(http.StatusCreated, gin.H{
		"message":"Profile created successfully",
		"data":profile,
	})
}
