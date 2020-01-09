package Controllers

import (
	"../Models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/* POST /v1/trainer */
func CreateTrainer(c *gin.Context) {
	var trainer Models.Trainer
	c.BindJSON(&trainer)
	if err := Models.CreateTrainer(trainer); err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, trainer)
	}
}

/* GET /v1/trainers/ */
func GetTrainers(c *gin.Context) {
	results, err := Models.GetAllTrainers()
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, results)
	}
}

func Test(c *gin.Context) {
	trainer, err := Models.GetTrainerById(c.Param("id"))
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, trainer)
}