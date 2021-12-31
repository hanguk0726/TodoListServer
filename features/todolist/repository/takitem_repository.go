package repository

import (

	"todolist-server/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskItemRepository struct {
	Conn *mongo.Database
}

func NewTaskItemRepository(Conn *mongo.Database) domain.TaskItemRepository {
	return &taskItemRepository{Conn}
}

func (t *taskItemRepository) AddTaskItem(userId int64, taskItem ...domain.TaskItem)
func (t *taskItemRepository) DeleteTaskItem(userId int64, taskItem ...domain.TaskItem)
func (t *taskItemRepository) GetTaskItemById(userId int64, taskItemId int64) domain.TaskItem
func (t *taskItemRepository) GetTaskItemsByTaskListId(userId int64, taskListId int64) []domain.TaskItem
func (t *taskItemRepository) UpdateTaskItem(userId int64, taskItem ...domain.TaskItem)