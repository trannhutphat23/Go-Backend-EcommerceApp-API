package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/trannhutphat23/Go-Backend-EcommerceApp-API/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1") 
	{
		v1.GET("/ping", controller.NewUserController().GetUserByID)
		// v1.PUT("/ping", Pong)
		// v1.PATCH("/ping", Pong)
		// v1.DELETE("/ping", Pong)
		// v1.HEAD("/ping", Pong)
		// v1.OPTIONS("/ping", Pong)
	}

	return r
}

// func Pong(c *gin.Context) {
// 	firstName := c.DefaultQuery("firstName", "defaultName")
// 	lastName := c.Query("lastName")

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "ping...pong " + firstName + " " + lastName,
// 	})	
// }