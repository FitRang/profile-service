package handlers

import (
	"reflect"
	"strings"
	"github.com/FitRang/profile-service/domain"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
) 


var (
	validate = validator.New()
)

func init() {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

type ProfileHandlerInterface interface {
	CreateProfileHandler(c *gin.Context) 
	GetProfileHandler(c *gin.Context) 
	UpdateProfileHandler(c *gin.Context) 
}

type ProfileHandler struct {
	domain *domain.ProfileService
}

func NewProfileHandler(domain *domain.ProfileService) *ProfileHandler {
	return &ProfileHandler{
		domain: domain,
	}
}
