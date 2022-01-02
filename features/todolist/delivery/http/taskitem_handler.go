package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"todolist-server/domain"
	"todolist-server/features/todolist/delivery/http/dto"
)

type TaskItemHandler struct {
	TaskItemUsecase domain.TaskItemUsecase
}

func NewTaskItemHandler(r *gin.Engine, usecase domain.TaskItemUsecase) {
	handler := &TaskItemHandler{
		TaskItemUsecase: usecase,
	}
	r.GET("/v1/taskItems", handler.GetTaskItemsByTaskListId)
	r.GET("/v1/taskItems/:taskListId", handler.GetTaskItemById)
	r.POST("/v1/taskItems", handler.InsertTaskItem)
	r.PUT("/v1/taskItems", handler.UpdateTaskItem)
	r.DELETE("/v1/taskItems", handler.DeleteTask)
	r.POST("/v1/taskItems/synchronizeTaskItem", handler.SynchronizeTaskItem)
}

func (h *TaskItemHandler) GetTaskItemsByTaskListId(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("userId"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	taskListId, err := strconv.ParseInt(c.Param("taskListId"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	taskItems := h.TaskItemUsecase.GetTaskItemsByTaskListId(userId, taskListId)

	taskItemDtos := make([]dto.TaskItemDto, len(taskItems))

	for i, v := range taskItems {
		taskItemDtos[i] = dto.ToTaskItemDto(v)
	}

	c.JSON(http.StatusOK, taskItemDtos)
}

func (h *TaskItemHandler) GetTaskItemById(c *gin.Context) {

	userId, err := strconv.ParseInt(c.Query("userId"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	taskItemId, err := strconv.ParseInt(c.Param("taskItemId"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	taskItem := h.TaskItemUsecase.GetTaskItemById(userId, taskItemId)

	c.JSON(http.StatusOK, dto.ToTaskItemDto(taskItem))
}

func (h *TaskItemHandler) InsertTaskItem(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		log.Fatal(err)
	}

	var taskItemDtos []dto.TaskItemDto

	json.Unmarshal([]byte(jsonData), &taskItemDtos)

	taskItems := make([]domain.TaskItem, len(taskItemDtos))

	for i, v := range taskItemDtos {
		taskItems[i] = v.ToTaskItem()
	}

	h.TaskItemUsecase.AddTaskItem(taskItems...)

	c.Status(http.StatusOK)
}

func (h *TaskItemHandler) DeleteTask(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		log.Fatal(err)
	}

	var taskItemDtos []dto.TaskItemDto

	json.Unmarshal([]byte(jsonData), &taskItemDtos)

	taskItems := make([]domain.TaskItem, len(taskItemDtos))

	for i, v := range taskItemDtos {
		taskItems[i] = v.ToTaskItem()
	}

	h.TaskItemUsecase.DeleteTaskItem(taskItems...)

	c.Status(http.StatusOK)
}

func (h *TaskItemHandler) UpdateTaskItem(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		log.Fatal(err)
	}

	var taskItemDtos []dto.TaskItemDto

	json.Unmarshal([]byte(jsonData), &taskItemDtos)

	taskItems := make([]domain.TaskItem, len(taskItemDtos))

	for i, v := range taskItemDtos {
		taskItems[i] = v.ToTaskItem()
	}

	h.TaskItemUsecase.UpdateTaskItem(taskItems...)

	c.Status(http.StatusOK)
}

func (h *TaskItemHandler) SynchronizeTaskItem(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		log.Fatal(err)
	}

	var taskItemDtos []dto.TaskItemDto

	json.Unmarshal([]byte(jsonData), &taskItemDtos)

	taskItems := make([]domain.TaskItem, len(taskItemDtos))

	for i, v := range taskItemDtos {
		taskItems[i] = v.ToTaskItem()
	}

	h.TaskItemUsecase.UpdateTaskItem(taskItems...)

	c.Status(http.StatusOK)
}
