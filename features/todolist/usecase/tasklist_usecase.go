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

func (t *taskListUseCase) AddTaskList(taskList ...domain.TaskList) {
	t.taskListRepository.AddTaskList(taskList...)
}
func (t *taskListUseCase) DeleteTaskList(taskList ...domain.TaskList) {
	t.taskListRepository.DeleteTaskList(taskList...)
}
func (t *taskListUseCase) UpdateTaskList(taskList ...domain.TaskList) {
	t.taskListRepository.UpdateTaskList(taskList...)
}
func (t *taskListUseCase) GetTaskListById(userId int64, taskListId int64) domain.TaskList {
	return t.taskListRepository.GetTaskListById(userId, taskListId)
}
func (t *taskListUseCase) GetTaskLists(userId int64) []domain.TaskList {
	return t.taskListRepository.GetTaskLists(userId)
}
