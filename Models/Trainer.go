package Models

import (
	"../Config"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `json:"name,omitempty"`
	Age  int                `json:"age,omitempty"`
	City string             `json:"city,omitempty"`
}

func CreateTrainer(trainer *Trainer) (err error) {
	trainer.ID = primitive.NewObjectID()
	_, err = Config.Trainers.InsertOne(context.TODO(), trainer)
	if err != nil {
		return err
	}
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

func GetTrainerById(id string) (trainer *Trainer, err error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return trainer, err
	}
	Config.Trainers.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&trainer)
	return trainer, nil
}

func UpdateTrainerById(id string, trainer *Trainer) (updatedTrainer *Trainer, err error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return trainer, err
	}
	updateValue := bson.M{}
	if trainer.Name != "" {
		updateValue["name"] = trainer.Name
	}
	if trainer.Age != 0 {
		updateValue["age"] = trainer.Age
	}
	if trainer.City != "" {
		updateValue["city"] = trainer.City
	}
	update := bson.M{
		"$set": updateValue,
	}
	after := options.After
	result := Config.Trainers.FindOneAndUpdate(
		context.TODO(),
		bson.M{"_id": _id},
		update,
		&options.FindOneAndUpdateOptions{ReturnDocument: &after})
	err = result.Decode(&updatedTrainer)
	if err != nil {
		return trainer, err
	}
	return updatedTrainer, nil
}
