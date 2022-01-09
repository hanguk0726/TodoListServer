package dto

import (
	"todolist-server/domain"
)

type TaskItemDto struct {
	Title            string `json:"title"`
	Detail           string `json:"detail"`
	IsCompleted      bool   `json:"is_completed"`
	CreatedTimestamp int64  `json:"created_timestamp"`
	TaskListId       int64  `json:"task_list_id"`
	Id               int64  `json:"id"`
	UserId           string `json:"user_id"`
}

func ToTaskItemDto(t domain.TaskItem) TaskItemDto {
	return TaskItemDto{
		Title:            t.Title,
		Detail:           t.Detail,
		IsCompleted:      t.IsCompleted,
		CreatedTimestamp: t.CreatedTimestamp,
		TaskListId:       t.TaskListId,
		Id:               t.Id,
		UserId:           t.UserId,
	}
}

func (t TaskItemDto) ToTaskItem() domain.TaskItem {
	return domain.TaskItem{
		Title:            t.Title,
		Detail:           t.Detail,
		IsCompleted:      t.IsCompleted,
		CreatedTimestamp: t.CreatedTimestamp,
		TaskListId:       t.TaskListId,
		Id:               t.Id,
		UserId:           t.UserId,
	}
}
