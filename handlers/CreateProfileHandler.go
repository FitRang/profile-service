package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/FitRang/profile-service/domain"
	"github.com/FitRang/profile-service/model"
	"github.com/gin-gonic/gin"
)

func (ph *ProfileHandler) CreateProfileHandler(c *gin.Context) {
	profile := model.ProfileCreateRequest{}
	if err := c.ShouldBind(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Custom Validation Error",
		})
		return
	} 
	err := ph.domain.CreateProfile(&profile)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrIDAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{
				"message":"A profile with this ID already exits",
			})
        case errors.Is(err, domain.ErrEmailAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{
				"message":"A profile with this email already exits",
			})
		case errors.Is(err, domain.ErrPhoneNumberAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{
				"message":"A profile with this phone number already exits",
			})
		default:
			log.Printf("Unexpected error while creating profile: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message":"Internal server error",
				"code":"INTERNAL_ERROR",
			})
		}
	} 
	c.JSON(http.StatusCreated, gin.H{
		"message":"Profile created successfully",
		"data":profile,
	})
}
