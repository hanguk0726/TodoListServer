package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"todolist-server/domain"
)

type taskItemRepository struct {
	Mongo *mongo.Collection
}

func NewTaskItemRepository(Mongo *mongo.Collection) domain.TaskItemRepository {
	return &taskItemRepository{Mongo}
}

func (t *taskItemRepository) AddTaskItem(taskItem ...domain.TaskItem) {
	if len(taskItem) == 1 {
		_, err := t.Mongo.InsertOne(context.TODO(), taskItem)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		inserts := make([]interface{}, len(taskItem))
		for i, v := range taskItem {
			inserts[i] = v
		}
		_, err := t.Mongo.InsertMany(context.TODO(), inserts)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func (t *taskItemRepository) DeleteTaskItem(taskItem ...domain.TaskItem) {
	if len(taskItem) == 1 {
		_, err := t.Mongo.DeleteOne(context.TODO(), taskItem)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		deletes := make([]interface{}, len(taskItem))
		for i, v := range taskItem {
			deletes[i] = v
		}
		_, err := t.Mongo.DeleteMany(context.TODO(), deletes)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func (t *taskItemRepository) UpdateTaskItem(taskItem ...domain.TaskItem) {
	if len(taskItem) == 1 {
		_, err := t.Mongo.UpdateOne(context.TODO(), bson.E{Key: "id", Value: taskItem[0].Id}, taskItem)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		var filter bson.D
		updates := make([]interface{}, len(taskItem))
		for i, v := range taskItem {
			updates[i] = v
			filter = append(filter, bson.E{Key: "id", Value: v.Id})
		}
		_, err := t.Mongo.UpdateMany(context.TODO(), filter, updates)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func (t *taskItemRepository) GetTaskItemById(userId int64, taskItemId int64) domain.TaskItem {
	result := t.Mongo.FindOne(context.TODO(), bson.M{"userId": userId, "taskItemId": taskItemId})
	var taskItem domain.TaskItem
	err := result.Decode(taskItem)
	if err != nil {
		log.Fatal(err)
	}
	return taskItem
}
func (t *taskItemRepository) GetTaskItemsByTaskListId(userId int64, taskListId int64) []domain.TaskItem {
	result, err := t.Mongo.Find(context.TODO(), bson.E{Key: "userId", Value: userId})
	if err != nil {
		log.Fatal(err)
	}
	var taskItems []domain.TaskItem
	err = result.Decode(taskItems)
	if err != nil {
		log.Fatal(err)
	}
	return taskItems
}
