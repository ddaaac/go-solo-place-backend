package Controllers

import (
	"../Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

/* POST /v1/trainer */
func CreateTrainer(c *gin.Context) {
	var trainer *Models.Trainer
	c.BindJSON(&trainer)
	if err := Models.CreateTrainer(trainer); err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, trainer)
	}
}

/* GET /v1/trainers */
func GetTrainers(c *gin.Context) {
	trainers, err := Models.GetAllTrainers()
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, trainers)
	}
}

/* GET /v1/trainer/:id */
func GetTrainer(c *gin.Context) {
	trainer, err := Models.GetTrainerById(c.Param("id"))
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, trainer)
	}
}

/* PUT /v1/trainer/:id */
func UpdateTrainer(c *gin.Context) {
	var trainer *Models.Trainer
	c.BindJSON(&trainer)
	trainer, err := Models.UpdateTrainerById(c.Param("id"), trainer)
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, trainer)
	}
}

func Test(c *gin.Context) {
	test := &struct {
		Id  primitive.ObjectID `json:"id"`
		Str string             `json:"str"`
		Num int                `json:"num"`
	}{}
	c.BindJSON(test)
	test.Id = primitive.NewObjectID()
	fmt.Println(test)
	//var trainer Models.Trainer
	//c.BindJSON(&trainer)
	c.JSON(http.StatusOK, test)
}

func makeBodyAsJson(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	c.JSON(http.StatusOK, reqBody)
}
