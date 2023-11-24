package users

import (
	"context"
	"log"
	"user_task_apps/app/models"
	usersRepository "user_task_apps/app/repository/users"
	"user_task_apps/pkg/utils/constants"
	"user_task_apps/pkg/utils/converter"
	"user_task_apps/pkg/utils/validator"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.Users
	err := usersRepository.GetAllUsers(c, &users)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  constants.STATUS_FAIL,
			"message": "Internal Service Error",
			"error":   err.Error(),
		})
		return
	}

	usersRes := converter.MapUsersToUsersRes(users)

	c.JSON(200, gin.H{
		"status":  constants.STATUS_SUCCESS,
		"message": "success",
		"data":    usersRes,
	})
	return
}

func GetUsersById(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	users := models.Users{}

	err := usersRepository.GetUsersById(ctx, id, &users)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  constants.STATUS_FAIL,
			"message": "users not found",
			"error":   err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"status":  constants.STATUS_SUCCESS,
		"message": "users found",
		"data":    users.ToUsersRes(),
	})
	return
}

func CreateUsers(c *gin.Context) {
	log.Println("MASUK SINI")
	usersRequest := new(models.UsersReq)

	if err := c.ShouldBindJSON(&usersRequest); err != nil {
		c.JSON(400, gin.H{
			"status":  constants.STATUS_FAIL,
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}

	errors := validator.ValidateRequest(usersRequest)
	if errors != nil {
		c.JSON(500, gin.H{
			"status":  constants.STATUS_FAIL,
			"message": "error",
			"errors":  errors,
		})
		return
	}
	log.Println("USERS REQUEST ====== ", usersRequest)

	users, _ := usersRepository.CreateUsers(c, usersRequest)
	log.Println("USERSSS ====== ", users)

	c.JSON(200, gin.H{
		"status":  constants.STATUS_SUCCESS,
		"message": "success",
		"data":    users.ToUsersRes(),
	})
	return
}

func UpdateUsers(c *gin.Context) {
	ctx := context.Background()
	usersRequest := new(models.UsersReq)

	if err := c.ShouldBindJSON(&usersRequest); err != nil {
		c.JSON(400, gin.H{
			"status":  constants.STATUS_FAIL,
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}

	errors := validator.ValidateRequest(usersRequest)
	if errors != nil {
		c.JSON(500, gin.H{
			"status":  constants.STATUS_FAIL,
			"message": "error",
			"errors":  errors,
		})
	}

	id := c.Param("id")
	users := models.Users{}
	err := usersRepository.GetUsersById(ctx, id, &users)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  constants.STATUS_FAIL,
			"message": "users not found",
		})
	}

	usersRepository.UpdateUsers(ctx, id, &users, usersRequest)

	c.JSON(200, gin.H{
		"status":  constants.STATUS_SUCCESS,
		"message": "users updated",
		"data":    users.ToUsersRes(),
	})
	return
}

func DeleteUsers(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	users := models.Users{}
	err := usersRepository.GetUsersById(ctx, id, &users)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  constants.STATUS_FAIL,
			"message": "users not found",
		})
	}

	usersRepository.DeleteUsers(id, &users)

	c.JSON(200, gin.H{
		"status":  constants.STATUS_SUCCESS,
		"message": "Success",
		"data":    nil,
	})
}
