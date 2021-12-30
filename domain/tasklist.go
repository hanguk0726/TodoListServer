package domain

type TaskList struct {
	Name string `json:"name"`
	CreatedTimestamp int64 `json:"created_timestamp"`
	Id int64 `json:"id"`
	UserId int64 `json:"user_id"`
}

type TaskListUsecase interface {
	AddTaskList(userId int64, taskList ...TaskList)
	DeleteTaskList(userId int64, taskList ...TaskList)
	GetTaskListById(userId int64, taskListId int64) TaskList
	GetTaskLists(userId int64) []TaskList
	UpdateTaskList(userId int64, taskList ...TaskList)
}

type TaskListRepository interface {
	AddTaskList(userId int64, taskList ...TaskList)
	DeleteTaskList(userId int64, taskList ...TaskList)
	GetTaskListById(userId int64, taskListId int64) TaskList
	GetTaskLists(userId int64) []TaskList
	UpdateTaskList(userId int64, taskList ...TaskList)
}