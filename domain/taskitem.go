package domain

type TaskItem struct {
	Title       string `json:"title" bson:"title"`
	Detail      string `json:"detail" bson:"detail"`
	IsCompleted bool   `json:"is_completed" bson:"isCompleted"`
	Timestamp   int64  `json:"timestamp" bson:"timestamp"`
	TaskListId  int64  `json:"task_list_id" bson:"taskListId"`
	Id          int64  `json:"id" bson:"id"`
	UserId      int64  `json:"user_id" bson:"userId"`
}

type TaskItemUsecase interface {
	AddTaskItem(taskItem ...TaskItem)
	DeleteTaskItem(taskItem ...TaskItem)
	GetTaskItemById(userId int64, taskItemId int64) TaskItem
	GetTaskItemsByTaskListId(userId int64, taskListId int64) []TaskItem
	UpdateTaskItem(taskItem ...TaskItem)
}

type TaskItemRepository interface {
	AddTaskItem(taskItem ...TaskItem)
	DeleteTaskItem(taskItem ...TaskItem)
	GetTaskItemById(userId int64, taskItemId int64) TaskItem
	GetTaskItemsByTaskListId(userId int64, taskListId int64) []TaskItem
	UpdateTaskItem(taskItem ...TaskItem)
}
