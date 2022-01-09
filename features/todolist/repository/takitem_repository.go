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
		_bson, err := bson.Marshal(taskItem[0])
		if err != nil {
			log.Println(err)
		}
		_, err = t.Mongo.InsertOne(context.TODO(), _bson)
		if err != nil {
			log.Println(err)
		}
	} else {
		inserts := make([]interface{}, len(taskItem))
		for i, v := range taskItem {
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
func (t *taskItemRepository) DeleteTaskItem(taskItem ...domain.TaskItem) {
	if len(taskItem) == 1 {
		_bson, err := bson.Marshal(taskItem[0])
		if err != nil {
			log.Println(err)
		}
		_, err = t.Mongo.DeleteOne(context.TODO(), _bson)
		if err != nil {
			log.Println(err)
		}
	} else {
		deletes := make([]interface{}, len(taskItem))
		for i, v := range taskItem {
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
func (t *taskItemRepository) UpdateTaskItem(taskItem ...domain.TaskItem) {
	if len(taskItem) == 1 {
		update := bson.M{
			"$set": bson.M{
				"title":            taskItem[0].Title,
				"detail":           taskItem[0].Detail,
				"isCompleted":      taskItem[0].IsCompleted,
				"createdTimestamp": taskItem[0].CreatedTimestamp,
				"taskListId":       taskItem[0].TaskListId,
				"id":               taskItem[0].Id}}
		_, err := t.Mongo.UpdateOne(context.TODO(), bson.M{"id": taskItem[0].Id}, update)
		if err != nil {
			log.Println(err)
		}
	} else {
		var filter []bson.M
		updates := make([]interface{}, len(taskItem))
		for i, v := range taskItem {
			updates[i] = bson.M{
				"$set": bson.M{
					"title":            taskItem[0].Title,
					"detail":           taskItem[0].Detail,
					"isCompleted":      taskItem[0].IsCompleted,
					"createdTimestamp": taskItem[0].CreatedTimestamp,
					"taskListId":       taskItem[0].TaskListId,
					"id":               taskItem[0].Id}}
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
func (t *taskItemRepository) GetTaskItemById(userId string, taskItemId int64) domain.TaskItem {
	result := t.Mongo.FindOne(context.TODO(), bson.M{"userId": userId, "id": taskItemId})
	var taskItem domain.TaskItem
	err := result.Decode(&taskItem)
	if err != nil {
		log.Println(err)
	}
	return taskItem
}
func (t *taskItemRepository) GetTaskItemsByTaskListId(userId string, taskListId int64) []domain.TaskItem {
	result, err := t.Mongo.Find(context.TODO(), bson.M{"userId": userId, "taskListId": taskListId})
	if err != nil {
		log.Println(err)
	}
	var taskItems []domain.TaskItem
	err = result.All(context.TODO(), &taskItems)
	if err != nil {
		log.Println(err)
	}
	return taskItems
}

func (t *taskItemRepository) DoesExists(taskItem domain.TaskItem) bool {
	result := t.Mongo.FindOne(context.TODO(), bson.M{"userId": taskItem.UserId, "taskListId": taskItem.TaskListId})
	err := result.Err()
	if err != nil {
		return false
	}
	return true
}
