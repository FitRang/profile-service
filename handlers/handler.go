package handlers

import (
	"github.com/FitRang/profile-service/domain"
	"github.com/gin-gonic/gin"
) 

type ProfileHandlerInterface interface {
	CreateProfileHandler(c *gin.Context) 
	GetProfileHandler(c *gin.Context) 
}

type ProfileHandler struct {
	domain *domain.ProfileService
}

func NewProfileHandler(domain *domain.ProfileService) *ProfileHandler {
	return &ProfileHandler{
		domain: domain,
	}
}
