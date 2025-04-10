package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ph *ProfileHandler) GetProfileHandler(c *gin.Engine) {
	type RequestParams struct {
		ID string `json:"id" form:"id" validate:"required,uuid"`
	}
	reqParams := RequestParams{
		ID: c.Param("id"),
	}
	if err := validate.Struct(reqParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"validation error",
		})
		return
	}
	profile, err := ph.domain.GetProfile(reqParams.ID)
}
