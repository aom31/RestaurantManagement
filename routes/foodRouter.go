package routes

import (
	"restuarant-management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.GET("/foods/:foodId", controllers.GetFood())
	incomingRoutes.POST("foods", controllers.CreateFood())
	incomingRoutes.PATCH("foods/:foodId", controllers.UpdateFood())
}
