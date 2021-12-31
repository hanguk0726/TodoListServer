package http

import (
	"github.com/gin-gonic/gin"
	"todolist-server/domain"
	"net/http"
	"strconv"
)

type TaskListHandler struct {
	TaskListUsecase domain.TaskListUsecase
}

// func NewTaskListHandler(r *gin.Engine, usecase domain.TaskItemUsecase){
// 	handler := &TaskListHandler{
// 		TaskListUsecase: usecase,
// 	}
// 	r.GET("/v1/taskLists/:userId", handler.GetTaskLists)
// }


func (h *TaskListHandler) GetTaskLists (c *gin.Context) {
	userId := strconv.ParseInt(c.Param("userId"), 10, 64)
	taskLists := h.TaskListUsecase.GetTaskLists(userId)



	c.JSON(http.StatusOK, taskLists)
}