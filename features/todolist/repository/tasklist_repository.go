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
		_, err := t.Mongo.InsertOne(context.TODO(), taskList)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		inserts := make([]interface{}, len(taskList))
		for i, v := range taskList {
			inserts[i] = v
		}
		_, err := t.Mongo.InsertMany(context.TODO(), inserts)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func (t *taskListRepository) DeleteTaskList(taskList ...domain.TaskList) {
	if len(taskList) == 1 {
		_, err := t.Mongo.DeleteOne(context.TODO(), taskList)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		deletes := make([]interface{}, len(taskList))
		for i, v := range taskList {
			deletes[i] = v
		}
		_, err := t.Mongo.DeleteMany(context.TODO(), deletes)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func (t *taskListRepository) UpdateTaskList(taskList ...domain.TaskList) {
	if len(taskList) == 1 {
		_, err := t.Mongo.UpdateOne(context.TODO(), bson.E{Key: "id", Value: taskList[0].Id}, taskList)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		var filter bson.D
		updates := make([]interface{}, len(taskList))
		for i, v := range taskList {
			updates[i] = v
			filter = append(filter, bson.E{Key: "id", Value: v.Id})
		}
		_, err := t.Mongo.UpdateMany(context.TODO(), filter, updates)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func (t *taskListRepository) GetTaskListById(userId int64, taskListId int64) domain.TaskList {
	result := t.Mongo.FindOne(context.TODO(), bson.M{"userId": userId, "taskListId": taskListId})
	var taskList domain.TaskList
	err := result.Decode(taskList)
	if err != nil {
		log.Fatal(err)
	}
	return taskList
}
func (t *taskListRepository) GetTaskLists(userId int64) []domain.TaskList {
	result, err := t.Mongo.Find(context.TODO(), bson.E{Key: "userId", Value: userId})
	if err != nil {
		log.Fatal(err)
	}
	var taskLists []domain.TaskList
	err = result.Decode(taskLists)
	if err != nil {
		log.Fatal(err)
	}
	return taskLists
}
