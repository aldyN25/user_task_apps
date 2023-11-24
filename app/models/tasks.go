package models

import (
	"time"

	"user_task_apps/pkg/utils/constants"
)

type (
	Tasks struct {
		ID          uint   `gorm:"primaryKey;autoIncrement"`
		UserID      uint   `gorm:"foreignKey:UserId;references:ID"`
		Title       string `gorm:"type:varchar(255)"`
		Description string `gorm:"type:text"`
		Status      string `gorm:"type:varchar(50);default:pending"`
		CreatedAt   time.Time
		UpdatedAt   *time.Time
		// ID        uint `gorm:"column:id;primary_key;auto_increment"`
		// UserId    uint
		// Title     string
		// Priority  string
		// IsActive  bool       `gorm:"column:is_active"`
		// CreatedAt time.Time  `gorm:"column:created_at"`
		// UpdatedAt *time.Time `gorm:"column:updated_at"`
		// Users     Users      `gorm:"foreignKey:UserId;references:ID"`
	}

	TasksRes struct {
		ID          uint   `json:"id"`
		UserID      uint   `json:"user_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
	}

	TasksReq struct {
		Title       string `json:"title" binding:"required" validate:"required,min=3"`
		UserID      uint   `json:"user_id" validate:"omitempty"`
		Description string `json:"description" binding:"required"`
		Status      string `json:"status" validate:"omitempty"`
	}
)

func (Tasks) TableName() string {
	return "tasks"
}

func (n Tasks) ToTasksRes() *TasksRes {
	result := &TasksRes{
		ID:          n.ID,
		UserID:      n.UserID,
		Title:       n.Title,
		Description: n.Description,
		Status:      n.Status,
		CreatedAt:   n.CreatedAt.Format(constants.LayoutYMD),
	}
	if n.UpdatedAt != nil {
		result.UpdatedAt = n.UpdatedAt.Format(constants.LayoutYMD)
	} else {
		result.UpdatedAt = ""
	}

	return result
}
