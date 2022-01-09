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
func (t *taskListUseCase) SynchronizeTaskList(taskList ...domain.TaskList) {
	for _, v := range taskList {
		if t.taskListRepository.DoesExists(v) {
			t.taskListRepository.UpdateTaskList(v)
		} else {
			t.taskListRepository.AddTaskList(v)
		}
	}
}

func (t *taskListUseCase) GetTaskListById(userId string, taskListId int64) domain.TaskList {
	return t.taskListRepository.GetTaskListById(userId, taskListId)
}
func (t *taskListUseCase) GetTaskLists(userId string) []domain.TaskList {
	return t.taskListRepository.GetTaskLists(userId)
}
