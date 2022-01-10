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

type TaskListHandler struct {
	TaskListUsecase domain.TaskListUsecase
}

func NewTaskListHandler(r *gin.Engine, usecase domain.TaskListUsecase) {
	handler := &TaskListHandler{
		TaskListUsecase: usecase,
	}
	r.GET("/v1/taskLists", handler.GetTaskLists)
	r.GET("/v1/taskLists/:taskListId", handler.GetTaskListById)
	r.POST("/v1/taskLists/", handler.InsertTaskList)
	r.PUT("/v1/taskLists/", handler.UpdateTaskList)
	r.DELETE("/v1/taskLists/", handler.DeleteTaskList)
	r.POST("/v1/taskLists/synchronizeTaskList", handler.SynchronizeTaskList)
}

func (h *TaskListHandler) GetTaskLists(c *gin.Context) {
	userId := c.Query("userId")

	taskLists := h.TaskListUsecase.GetTaskLists(userId)

	taskListDtos := make([]dto.TaskListDto, len(taskLists))

	for i, v := range taskLists {
		taskListDtos[i] = dto.ToTaskListDto(v)
	}

	c.JSON(http.StatusOK, taskListDtos)
}

func (h *TaskListHandler) GetTaskListById(c *gin.Context) {

	userId := c.Query("userId")

	taskListId, err := strconv.ParseInt(c.Param("taskListId"), 10, 64)

	if err != nil {
		log.Println(err)
	}

	taskList := h.TaskListUsecase.GetTaskListById(userId, taskListId)

	c.JSON(http.StatusOK, dto.ToTaskListDto(taskList))
}

func (h *TaskListHandler) InsertTaskList(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		log.Println(err)
	}

	var taskListDtos []dto.TaskListDto
	err = json.Unmarshal([]byte(jsonData), &taskListDtos)
	if err != nil {
		log.Println(err)
	}
	taskLists := make([]domain.TaskList, len(taskListDtos))

	for i, v := range taskListDtos {
		taskLists[i] = v.ToTaskList()
	}
	h.TaskListUsecase.AddTaskList(taskLists...)

	c.Status(http.StatusOK)
}

func (h *TaskListHandler) DeleteTaskList(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		log.Println(err)
	}

	var taskListDtos []dto.TaskListDto

	err = json.Unmarshal([]byte(jsonData), &taskListDtos)
	if err != nil {
		log.Println(err)
	}

	taskLists := make([]domain.TaskList, len(taskListDtos))

	for i, v := range taskListDtos {
		taskLists[i] = v.ToTaskList()
	}

	h.TaskListUsecase.DeleteTaskList(taskLists...)

	c.Status(http.StatusOK)
}

func (h *TaskListHandler) UpdateTaskList(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		log.Println(err)
	}

	var taskListDtos []dto.TaskListDto

	err = json.Unmarshal([]byte(jsonData), &taskListDtos)
	if err != nil {
		log.Println(err)
	}

	taskLists := make([]domain.TaskList, len(taskListDtos))

	for i, v := range taskListDtos {
		taskLists[i] = v.ToTaskList()
	}

	h.TaskListUsecase.UpdateTaskList(taskLists...)

	c.Status(http.StatusOK)
}

func (h *TaskListHandler) SynchronizeTaskList(c *gin.Context) {

	jsonData, err := c.GetRawData()
	if err != nil {
		log.Println(err)
	}

	var taskListDtos []dto.TaskListDto

	err = json.Unmarshal([]byte(jsonData), &taskListDtos)
	if err != nil {
		log.Println(err)
	}

	taskLists := make([]domain.TaskList, len(taskListDtos))

	for i, v := range taskListDtos {
		taskLists[i] = v.ToTaskList()
	}

	h.TaskListUsecase.SynchronizeTaskList(taskLists...)

	c.Status(http.StatusOK)
}
