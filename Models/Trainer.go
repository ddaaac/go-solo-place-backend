package Models

import (
	"../Config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func CreateTrainer(trainer Trainer) (err error) {
	insertResult, err := Config.Trainers.InsertOne(context.TODO(), trainer)
	if err != nil {
		return err
	}
	fmt.Println(insertResult)
	return nil
}

func GetAllTrainers() (results []*Trainer, err error) {
	cur, err := Config.Trainers.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem Trainer
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context.TODO())
	return results, nil
}

func GetTrainerById(id string) (trainer Trainer, err error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return trainer, err
	}
	Config.Trainers.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&trainer)
	return trainer, nil
}
