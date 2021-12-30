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


func (t *taskItemUseCase) AddTaskItem(userId int64, taskItem ...domain.TaskItem) {

}


func (t *taskItemUseCase) DeleteTaskItem(userId int64, taskItem ...domain.TaskItem){

}

func (t *taskItemUseCase) GetTaskItemById(userId int64, taskItemId int64) domain.TaskItem

func (t *taskItemUseCase) GetTaskItemsByTaskListId(userId int64, taskListId int64) []domain.TaskItem 

func (t *taskItemUseCase)  UpdateTaskItem(userId int64, taskItem ...domain.TaskItem){

}