package repository

import (

	"todolist-server/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskListRepository struct {
	Conn *mongo.Database
}

func NewTaskListRepository(Conn *mongo.Database) domain.TaskListRepository {
	return &taskListRepository{Conn}
}


func (t *taskListRepository) AddTaskList(userId int64, taskList ...domain.TaskList)
func (t *taskListRepository) DeleteTaskList(userId int64, taskList ...domain.TaskList)
func (t *taskListRepository) GetTaskListById(userId int64, taskListId int64) domain.TaskList
func (t *taskListRepository) GetTaskLists(userId int64) []domain.TaskList
func (t *taskListRepository) UpdateTaskList(userId int64, taskList ...domain.TaskList)