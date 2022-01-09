package dto

import (
	"todolist-server/domain"
)

type TaskListDto struct {
	Name             string `json:"name"`
	CreatedTimestamp int64  `json:"created_timestamp"`
	Id               int64  `json:"id"`
	UserId           string `json:"user_id"`
}

func ToTaskListDto(t domain.TaskList) TaskListDto {
	return TaskListDto{
		Name:             t.Name,
		CreatedTimestamp: t.CreatedTimestamp,
		Id:               t.Id,
		UserId:           t.UserId,
	}
}

func (t TaskListDto) ToTaskList() domain.TaskList {
	return domain.TaskList{
		Name:             t.Name,
		CreatedTimestamp: t.CreatedTimestamp,
		Id:               t.Id,
		UserId:           t.UserId,
	}
}
