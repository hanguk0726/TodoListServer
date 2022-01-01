package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"todolist-server/domain"
	"todolist-server/features/todolist/delivery/http/dto"
	"github.com/gin-gonic/gin"
)

type TaskItemHandler struct {
	TaskItemUsecase domain.TaskItemUsecase
}

func NewTaskItemHandler(r *gin.Engine, usecase domain.TaskItemUsecase){
	handler := &TaskItemHandler{
		TaskItemUsecase: usecase,
	}
	r.GET("/v1/taskLists/:userId", handler.GetTaskItemsByTaskListId)
	r.GET("/v1/taskLists/:taskListId/:userId", handler.GetTaskItemById)
	r.POST("/v1/taskLists/:userId", handler.InsertTaskItem)
	r.PUT("/v1/taskLists/:userId", handler.UpdateTaskItem)
	r.DELETE("/v1/taskLists/:userId", handler.DeleteTask)
	r.POST("/v1/taskLists/:userId", handler.SynchronizeTaskItem)
}



func (h *TaskItemHandler) GetTaskItemsByTaskListId (c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	taskListId, err := strconv.ParseInt(c.Param("taskListId"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	taskItems := h.TaskItemUsecase.GetTaskItemsByTaskListId(userId, taskListId)

	taskItemDtos := make([]dto.TaskItemDto, len(taskItems))

	for i , v := range taskItems {
		taskItemDtos[i] = dto.ToTaskItemDto(v)
	}

	c.JSON(http.StatusOK, taskItemDtos)
}

func (h *TaskItemHandler) GetTaskItemById (c *gin.Context) {

	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

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


func (h *TaskItemHandler) InsertTaskItem (c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}
	
	jsonData, err := c.GetRawData()
	if err != nil {
		log.Fatal(err)
	}

	var taskItemDtos []dto.TaskItemDto

	json.Unmarshal([]byte(jsonData), &taskItemDtos)


	taskItems := make([]domain.TaskItem, len(taskItemDtos))

	for i , v := range taskItemDtos {
		taskItems[i] = v.ToTaskItem()
	}

	h.TaskItemUsecase.AddTaskItem(userId, taskItems...)

	c.Status(http.StatusOK)
}


func (h *TaskItemHandler) DeleteTask (c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}
	
	jsonData, err := c.GetRawData()
	if err != nil {
		log.Fatal(err)
	}

	var taskItemDtos []dto.TaskItemDto

	json.Unmarshal([]byte(jsonData), &taskItemDtos)


	taskItems := make([]domain.TaskItem, len(taskItemDtos))

	for i , v := range taskItemDtos {
		taskItems[i] = v.ToTaskItem()
	}


	h.TaskItemUsecase.DeleteTaskItem(userId, taskItems...)

	c.Status(http.StatusOK)
}


func (h *TaskItemHandler) UpdateTaskItem (c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}
	
	jsonData, err := c.GetRawData()
	if err != nil {
		log.Fatal(err)
	}

	var taskItemDtos []dto.TaskItemDto

	json.Unmarshal([]byte(jsonData), &taskItemDtos)


	taskItems := make([]domain.TaskItem, len(taskItemDtos))

	for i , v := range taskItemDtos {
		taskItems[i] = v.ToTaskItem()
	}


	h.TaskItemUsecase.UpdateTaskItem(userId, taskItems...)

	c.Status(http.StatusOK)
}


func (h *TaskItemHandler) SynchronizeTaskItem (c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}
	
	jsonData, err := c.GetRawData()
	if err != nil {
		log.Fatal(err)
	}

	var taskItemDtos []dto.TaskItemDto

	json.Unmarshal([]byte(jsonData), &taskItemDtos)


	taskItems := make([]domain.TaskItem, len(taskItemDtos))

	for i , v := range taskItemDtos {
		taskItems[i] = v.ToTaskItem()
	}

	h.TaskItemUsecase.UpdateTaskItem(userId, taskItems...)

	c.Status(http.StatusOK)
}

