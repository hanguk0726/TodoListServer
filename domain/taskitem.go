package domain

type TaskItem struct {
	Title            string `json:"title" bson:"title"`
	Detail           string `json:"detail" bson:"detail"`
	IsCompleted      bool   `json:"is_completed" bson:"isCompleted"`
	CreatedTimestamp int64  `json:"created_timestamp" bson:"createdTimestamp"`
	TaskListId       int64  `json:"task_list_id" bson:"taskListId"`
	Id               int64  `json:"id" bson:"id"`
	UserId           string `json:"user_id" bson:"userId"`
}

type TaskItemUsecase interface {
	AddTaskItem(taskItem ...TaskItem)
	DeleteTaskItem(taskItem ...TaskItem)
	GetTaskItemById(userId string, taskItemId int64) TaskItem
	GetTaskItemsByTaskListId(userId string, taskListId int64) []TaskItem
	UpdateTaskItem(taskItem ...TaskItem)
	SynchronizeTaskList(taskItem ...TaskItem)
}

type TaskItemRepository interface {
	AddTaskItem(taskItem ...TaskItem)
	DeleteTaskItem(taskItem ...TaskItem)
	GetTaskItemById(userId string, taskItemId int64) TaskItem
	GetTaskItemsByTaskListId(userId string, taskListId int64) []TaskItem
	UpdateTaskItem(taskItem ...TaskItem)
	DoesExists(taskItem TaskItem) bool
}
