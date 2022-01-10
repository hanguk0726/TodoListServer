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
	r.GET("/v1/taskItems", handler.GetTaskItemById)
	r.GET("/v1/taskItems/:taskListId", handler.GetTaskItemsByTaskListId)
	r.POST("/v1/taskItems", handler.InsertTaskItem)
	r.PUT("/v1/taskItems", handler.UpdateTaskItem)
	r.DELETE("/v1/taskItems", handler.DeleteTask)
	r.POST("/v1/taskItems/synchronizeTaskItem", handler.SynchronizeTaskItem)
}

func (h *TaskItemHandler) GetTaskItemsByTaskListId(c *gin.Context) {

	userId := c.Query("userId")

	taskListId, err := strconv.ParseInt(c.Param("taskListId"), 10, 64)

	if err != nil {
		log.Println(err)
	}

	taskItems := h.TaskItemUsecase.GetTaskItemsByTaskListId(userId, taskListId)

	taskItemDtos := make([]dto.TaskItemDto, len(taskItems))

	for i, v := range taskItems {
		taskItemDtos[i] = dto.ToTaskItemDto(v)
	}

	c.JSON(http.StatusOK, taskItemDtos)
}

func (h *TaskItemHandler) GetTaskItemById(c *gin.Context) {

	userId := c.Query("userId")

	taskItemId, err := strconv.ParseInt(c.Query("taskItemId"), 10, 64)

	if err != nil {
		log.Println(err)
	}

	taskItem := h.TaskItemUsecase.GetTaskItemById(userId, taskItemId)

	c.JSON(http.StatusOK, dto.ToTaskItemDto(taskItem))
}

func (h *TaskItemHandler) InsertTaskItem(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		log.Println(err)
	}

	var taskItemDtos []dto.TaskItemDto

	err = json.Unmarshal([]byte(jsonData), &taskItemDtos)

	if err != nil {
		log.Println(err)
	}

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
		log.Println(err)
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
		log.Println(err)
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
		log.Println(err)
	}

	var taskItemDtos []dto.TaskItemDto

	json.Unmarshal([]byte(jsonData), &taskItemDtos)

	taskItems := make([]domain.TaskItem, len(taskItemDtos))

	for i, v := range taskItemDtos {
		taskItems[i] = v.ToTaskItem()
	}

	h.TaskItemUsecase.SynchronizeTaskList(taskItems...)

	c.Status(http.StatusOK)
}
