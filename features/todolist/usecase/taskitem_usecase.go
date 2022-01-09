package usecase

import (
	"todolist-server/domain"
)

type taskItemUseCase struct {
	taskItemRepository domain.TaskItemRepository
}

func NewTaskItemUseCase(t domain.TaskItemRepository) domain.TaskItemUsecase {
	return &taskItemUseCase{
		taskItemRepository: t,
	}
}

func (t *taskItemUseCase) AddTaskItem(taskItem ...domain.TaskItem) {
	t.taskItemRepository.AddTaskItem(taskItem...)
}

func (t *taskItemUseCase) DeleteTaskItem(taskItem ...domain.TaskItem) {
	t.taskItemRepository.DeleteTaskItem(taskItem...)
}

func (t *taskItemUseCase) UpdateTaskItem(taskItem ...domain.TaskItem) {
	t.taskItemRepository.UpdateTaskItem(taskItem...)
}
func (t *taskItemUseCase) GetTaskItemById(userId string, taskItemId int64) domain.TaskItem {
	return t.taskItemRepository.GetTaskItemById(userId, taskItemId)
}

func (t *taskItemUseCase) GetTaskItemsByTaskListId(userId string, taskListId int64) []domain.TaskItem {
	return t.taskItemRepository.GetTaskItemsByTaskListId(userId, taskListId)
}
