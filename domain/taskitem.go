package domain

type TaskItem struct {
	Title string `json:"title"`
	Detail string `json:"detail"`
	IsCompleted bool `json:"is_completed"`
	Timestamp string `json:"timestamp"`	
	TaskListId int64 `json:"task_list_id"`
	Id int64 `json:"id"`
}

type TaskItemUsecase interface {

	AddTaskItem(userId int64, taskItem ...TaskItem)
	DeleteTaskItem(userId int64, taskItem ...TaskItem)
	GetTaskItemById(userId int64, taskItemId int64) TaskItem
	GetTaskItemsByTaskListId(userId int64, taskListId int64) []TaskItem
	UpdateTaskItem(userId int64, taskItem ...TaskItem)

}

type TaskItemRepository interface {

	AddTaskItem(userId int64, taskItem ...TaskItem)
	DeleteTaskItem(userId int64, taskItem ...TaskItem)
	GetTaskItemById(userId int64, taskItemId int64) TaskItem
	GetTaskItemsByTaskListId(userId int64, taskListId int64) []TaskItem
	UpdateTaskItem(userId int64, taskItem ...TaskItem)

}