package tasks

import (
	"context"
	"fmt"
	"net/http"

	"user_task_apps/app/models"
	tasksRepository "user_task_apps/app/repository/tasks"
	"user_task_apps/pkg/utils/constants"
	"user_task_apps/pkg/utils/converter"
	"user_task_apps/pkg/utils/validator"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	ctx := context.Background()
	var tasks []models.Tasks
	err := tasksRepository.GetAllTasks(ctx, &tasks)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  constants.STATUS_FAIL,
			"message": "Internal Service Error",
			"error":   err.Error(),
		})
		return
	}

	tasksRes := converter.MapTasksToTasksRes(tasks)
	c.JSON(200, gin.H{
		"status":  constants.STATUS_SUCCESS,
		"message": "success",
		"data":    tasksRes,
	})
	return
}

func GetTasksById(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	tasks := models.Tasks{}

	err := tasksRepository.GetTasksById(ctx, id, &tasks)
	if err != nil {
		msgErr := fmt.Sprintf("tasks with ID %v Not Found", id)
		c.JSON(400, gin.H{
			"status":  constants.STATUS_FAIL,
			"message": msgErr,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  constants.STATUS_SUCCESS,
		"message": "success",
		"data":    tasks.ToTasksRes(),
	})
	return
}

func CreateTasks(c *gin.Context) {
	ctx := context.Background()
	tasksRequest := new(models.TasksReq)

	if err := c.ShouldBindJSON(&tasksRequest); err != nil {
		c.JSON(400, gin.H{
			"status":  constants.STATUS_FAIL,
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}
	errors := validator.ValidateRequest(tasksRequest)
	if errors != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  constants.STATUS_FAIL,
			"message": "error",
			"error":   errors,
		})
		return
	}

	tasks, _ := tasksRepository.CreateTasks(ctx, tasksRequest)

	c.JSON(200, gin.H{
		"status":  constants.STATUS_SUCCESS,
		"message": "success",
		"data":    tasks.ToTasksRes(),
	})
	return
}

func UpdateTasks(c *gin.Context) {
	ctx := context.Background()
	tasksRequest := new(models.TasksReq)

	if err := c.ShouldBindJSON(&tasksRequest); err != nil {
		c.JSON(400, gin.H{
			"status":  constants.STATUS_FAIL,
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}

	errors := validator.ValidateRequest(tasksRequest)
	if errors != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  constants.STATUS_FAIL,
			"message": "error",
			"error":   errors,
		})
		return
	}

	id := c.Param("id")
	tasks := models.Tasks{}
	err := tasksRepository.GetTasksById(ctx, id, &tasks)

	if err != nil {
		msgErr := fmt.Sprintf("tasks with ID %v Not Found", id)
		c.JSON(400, gin.H{
			"status":  "Not Found",
			"message": msgErr,
			"error":   err.Error(),
		})
		return
	}

	tasksRepository.UpdateTasks(ctx, id, &tasks, tasksRequest)

	c.JSON(200, gin.H{
		"status":  constants.STATUS_SUCCESS,
		"message": "success",
		"data":    tasks.ToTasksRes(),
	})
	return
}

func DeleteTasks(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	tasks := models.Tasks{}
	err := tasksRepository.GetTasksById(ctx, id, &tasks)

	if err != nil {
		msgErr := fmt.Sprintf("tasks with ID %v Not Found", id)
		c.JSON(400, gin.H{
			"status":  "Not Found",
			"message": msgErr,
			"error":   err.Error(),
		})
		return
	}

	tasksRepository.DeleteTasks(id, &tasks)

	c.JSON(200, gin.H{
		"status":  constants.STATUS_SUCCESS,
		"message": "Success",
		"data":    nil,
	})
	return
}
