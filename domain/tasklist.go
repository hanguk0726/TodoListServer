package domain

type TaskList struct {
	Name             string `json:"name" bson:"name"`
	CreatedTimestamp int64  `json:"created_timestamp" bson:"createdTimestamp"`
	Id               int64  `json:"id" bson:"id"`
	UserId           int64  `json:"user_id" bson:"userId"`
}

type TaskListUsecase interface {
	AddTaskList(taskList ...TaskList)
	DeleteTaskList(taskList ...TaskList)
	GetTaskListById(userId int64, taskListId int64) TaskList
	GetTaskLists(userId int64) []TaskList
	UpdateTaskList(taskList ...TaskList)
}

type TaskListRepository interface {
	AddTaskList(taskList ...TaskList)
	DeleteTaskList(taskList ...TaskList)
	GetTaskListById(userId int64, taskListId int64) TaskList
	GetTaskLists(userId int64) []TaskList
	UpdateTaskList(taskList ...TaskList)
}
