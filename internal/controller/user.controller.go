package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/trannhutphat23/Go-Backend-EcommerceApp-API/internal/service"
	"github.com/trannhutphat23/Go-Backend-EcommerceApp-API/pkg/response"
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
	response.SuccessResponse(c, 20001, "test success")
}
