package http

import (
	"todolist-server/domain"
)


type TaskItemDto struct {
	Title string `json:"title"`
	Detail string `json:"detail"`
	IsCompleted bool `json:"is_completed"`
	Timestamp string `json:"timestamp"`	
	TaskListId int64 `json:"task_list_id"`
	Id int64 `json:"id"`
}



func ToTaskItemDto(t domain.TaskItem) TaskItemDto {
	return TaskItemDto{
		Title: t.Title,
		Detail: t.Detail,
		IsCompleted: t.IsCompleted,
		Timestamp: t.Timestamp,
		TaskListId: t.TaskListId,
		Id: t.Id,
	}
}

func (t TaskItemDto) ToTaskItem() domain.TaskItem {
	return domain.TaskItem{
		Title: t.Title,
		Detail: t.Detail,
		IsCompleted: t.IsCompleted,
		Timestamp: t.Timestamp,
		TaskListId: t.TaskListId,
		Id: t.Id,
	}
}
