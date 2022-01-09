package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"todolist-server/domain"
)

type taskListRepository struct {
	Mongo *mongo.Collection
}

func NewTaskListRepository(Mongo *mongo.Collection) domain.TaskListRepository {
	return &taskListRepository{Mongo}
}

func (t *taskListRepository) AddTaskList(taskList ...domain.TaskList) {
	if len(taskList) == 1 {
		_bson, err := bson.Marshal(taskList[0])
		if err != nil {
			log.Println(err)
		}
		_, err = t.Mongo.InsertOne(context.TODO(), _bson)
		if err != nil {
			log.Println(err)
		}
	} else {
		inserts := make([]interface{}, len(taskList))
		for i, v := range taskList {
			var err error
			inserts[i], err = bson.Marshal(v)
			if err != nil {
				log.Println(err)
			}
		}
		_, err := t.Mongo.InsertMany(context.TODO(), inserts)
		if err != nil {
			log.Println(err)
		}
	}
}
func (t *taskListRepository) DeleteTaskList(taskList ...domain.TaskList) {
	if len(taskList) == 1 {
		_bson, err := bson.Marshal(taskList[0])
		if err != nil {
			log.Println(err)
		}
		_, err = t.Mongo.DeleteOne(context.TODO(), _bson)
		if err != nil {
			log.Println(err)
		}
	} else {
		deletes := make([]interface{}, len(taskList))
		for i, v := range taskList {
			var err error
			deletes[i], err = bson.Marshal(v)
			if err != nil {
				log.Println(err)
			}
		}
		_, err := t.Mongo.DeleteMany(context.TODO(), deletes)
		if err != nil {
			log.Println(err)
		}
	}
}
func (t *taskListRepository) UpdateTaskList(taskList ...domain.TaskList) {
	if len(taskList) == 1 {
		update := bson.M{
			"$set": bson.M{
				"name":             taskList[0].Name,
				"createdTimestamp": taskList[0].CreatedTimestamp,
				"id":               taskList[0].Id,
				"userId":           taskList[0].UserId}}
		_, err := t.Mongo.UpdateOne(context.TODO(), bson.M{"id": taskList[0].Id}, update)
		if err != nil {
			log.Println(err)
		}
	} else {
		var filter []bson.M
		updates := make([]interface{}, len(taskList))
		for i, v := range taskList {
			updates[i] = bson.M{
				"$set": bson.M{
					"name":             v.Name,
					"createdTimestamp": v.CreatedTimestamp,
					"id":               v.Id,
					"userId":           v.UserId}}
			filter = append(filter, bson.M{"id": v.Id})
		}
		for i, v := range updates {
			_, err := t.Mongo.UpdateOne(context.TODO(), filter[i], v)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
func (t *taskListRepository) GetTaskListById(userId string, taskListId int64) domain.TaskList {

	result := t.Mongo.FindOne(context.TODO(), bson.M{"userId": userId, "id": taskListId})
	var taskList domain.TaskList
	err := result.Decode(&taskList)
	if err != nil {
		log.Println(err)
	}
	return taskList
}
func (t *taskListRepository) GetTaskLists(userId string) []domain.TaskList {
	result, err := t.Mongo.Find(context.TODO(), bson.M{"userId": userId})
	if err != nil {
		log.Println(err)
	}

	var taskLists []domain.TaskList
	err = result.All(context.TODO(), &taskLists)
	if err != nil {
		log.Println(err)
	}
	return taskLists
}
