package domain

type TaskList struct {
	Name             string `json:"name" bson:"name"`
	CreatedTimestamp int64  `json:"created_timestamp" bson:"createdTimestamp"`
	Id               int64  `json:"id" bson:"id"`
	UserId           string `json:"user_id" bson:"userId"`
}

type TaskListUsecase interface {
	AddTaskList(taskList ...TaskList)
	DeleteTaskList(taskList ...TaskList)
	GetTaskListById(userId string, taskListId int64) TaskList
	GetTaskLists(userId string) []TaskList
	UpdateTaskList(taskList ...TaskList)
	SynchronizeTaskList(taskList ...TaskList)
}

type TaskListRepository interface {
	AddTaskList(taskList ...TaskList)
	DeleteTaskList(taskList ...TaskList)
	GetTaskListById(userId string, taskListId int64) TaskList
	GetTaskLists(userId string) []TaskList
	UpdateTaskList(taskList ...TaskList)
	DoesExists(taskList TaskList) bool
}
