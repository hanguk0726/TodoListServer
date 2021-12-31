package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"todolist-server/domain"
	"todolist-server/dto"
	"todolist-server/features/todolist/delivery/http/dto"

	"github.com/gin-gonic/gin"
)

type TaskListHandler struct {
	TaskListUsecase domain.TaskListUsecase
}

func NewTaskListHandler(r *gin.Engine, usecase domain.TaskListUsecase){
	handler := &TaskListHandler{
		TaskListUsecase: usecase,
	}
	r.GET("/v1/taskLists/:userId", handler.GetTaskLists)
}


func (h *TaskListHandler) GetTaskLists (c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	taskLists := h.TaskListUsecase.GetTaskLists(userId)

	taskListDtos := make([]dto.TaskListDto, len(taskLists))

	for i , v := range taskLists {
		taskListDtos[i] = dto.ToTaskListDto(v)
	}

	c.JSON(http.StatusOK, taskListDtos)
}

func (h *TaskListHandler) GetTaskListById (c *gin.Context) {

	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	taskListId, err := strconv.ParseInt(c.Param("taskListId"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}


	taskList := h.TaskListUsecase.GetTaskListById(userId, taskListId)


	c.JSON(http.StatusOK, dto.ToTaskListDto(taskList))
}


func (h *TaskListHandler) InsertTaskList (c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}
	
	jsonData, err := c.GetRawData()
	if err != nil {
		log.Fatal(err)
	}

	var taskListDtos []dto.TaskListDto

	json.Unmarshal([]byte(jsonData), &taskListDtos)



	taskLists := make([]domain.TaskList, len(taskListDtos))

	for i , v := range taskLists {
		taskLists[i] = taskListDtos.ToTaskListDto(v)
	}


	h.TaskListUsecase.AddTaskList(userId, taskLists...)

	c.Status(http.StatusOK)
}

// @GET("/v1/taskLists")
// suspend fun getTaskLists(@Query("userId") userId: String): List<TaskListDto>

// @GET("/v1/taskLists/{taskListId}")
// suspend fun getTaskListById(@Path("taskListId") taskListId: Long, @Query("userId") userId: String) : TaskListDto

// @POST("/v1/taskLists")
// suspend fun insertTaskList(vararg taskListDto: TaskListDto, @Query("userId") userId: String) : Call<ResponseBody>

// @DELETE("/v1/taskLists")
// suspend fun deleteTaskList(vararg taskListDto: TaskListDto, @Query("userId") userId: String) : Call<ResponseBody>

// @PUT("/v1/taskLists")
// suspend fun updateTaskList(vararg taskListDto: TaskListDto, @Query("userId") userId: String) : Call<ResponseBody>

// @POST
// suspend fun synchronizeTaskList(vararg taskListDto: TaskListDto, @Query("userId") userId: String) : Call<ResponseBody>
