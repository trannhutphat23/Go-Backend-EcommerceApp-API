package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trannhutphat23/Go-Backend-EcommerceApp-API/internal/service"
)

type USerController struct {
	userService *service.UserService
}

func NewUserController() *USerController {
	return &USerController{
		userService: service.NewUserService(),
	}
}

// controller -> service -> repo -> models -> dbs
func (uc *USerController) GetUserByID(c *gin.Context) {
	firstName := c.DefaultQuery("firstName", "defaultName")
	lastName := c.Query("lastName")

	c.JSON(http.StatusOK, gin.H{
		"message": "ping...pong " + firstName + " " + lastName,
	})
}