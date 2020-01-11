package Routes

import (
	"../Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/trainers", Controllers.GetTrainers)
		v1.POST("/trainer", Controllers.CreateTrainer)
		v1.GET("/trainer/:id", Controllers.GetTrainer)
		v1.GET("/test/:id", Controllers.Test)
		//v1.GET("todo/:id", Controllers.GetATodo)
		v1.PUT("/trainer/:id", Controllers.UpdateTrainer)
		//v1.DELETE("todo/:id", Controllers.DeleteATodo)
	}
}
