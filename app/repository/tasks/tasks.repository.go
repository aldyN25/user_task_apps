package tasks

import (
	"context"
	"time"

	"user_task_apps/app/models"
	gormdb "user_task_apps/pkg/gorm.db"
)

func GetAllTasks(ctx context.Context, tasks *[]models.Tasks) error {
	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}
	err = db.Debug().Table("tasks").Find(&tasks).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTasksById(ctx context.Context, id string, tasks *models.Tasks) error {
	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}

	err = db.Debug().Table("tasks").First(&tasks, id).Error

	if err != nil {
		return err
	}

	return nil
}

func CreateTasks(ctx context.Context, tasksRequest *models.TasksReq) (*models.Tasks, error) {
	db, err := gormdb.GetInstance()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	tasks := &models.Tasks{
		UserID:      tasksRequest.UserID,
		Title:       tasksRequest.Title,
		Description: tasksRequest.Description,
		Status:      tasksRequest.Status,
		CreatedAt:   now,
	}

	err = db.Debug().Table("tasks").Create(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func UpdateTasks(ctx context.Context, id string, tasks *models.Tasks, tasksRequest *models.TasksReq) (*models.TasksRes, error) {
	db, err := gormdb.GetInstance()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	dataUpdates := models.Tasks{
		Title:       tasksRequest.Title,
		Description: tasksRequest.Description,
		Status:      tasksRequest.Status,
		UpdatedAt:   &now,
	}
	err = db.Debug().Table("tasks").Where("id = ?", id).Updates(&dataUpdates).Error
	if err != nil {
		return nil, err
	}

	err = db.Debug().Table("tasks").First(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks.ToTasksRes(), nil
}

func DeleteTasks(id string, task *models.Tasks) error {
	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}

	err = db.Debug().Table("tasks").Delete(&task).Error
	if err != nil {
		return err
	}
	return nil
}
