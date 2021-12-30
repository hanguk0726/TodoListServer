package usecase

import (	

	"todolist-server/domain"
)



type taskListUseCase struct {
	taskListRepository domain.TaskListRepository
}



func NewTaskListUseCase(t domain.TaskListRepository) domain.TaskListUsecase {
	return &taskListUseCase{
		taskListRepository: t,
	}
} 


func (t *taskListUseCase)AddTaskList(userId int64, taskList ...domain.TaskList)
func (t *taskListUseCase)DeleteTaskList(userId int64, taskList ...domain.TaskList)
func (t *taskListUseCase)GetTaskListById(userId int64, taskListId int64) domain.TaskList
func (t *taskListUseCase)GetTaskLists(userId int64) []domain.TaskList
func (t *taskListUseCase)UpdateTaskList(userId int64, taskList ...domain.TaskList)