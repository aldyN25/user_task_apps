package models

import (
	"time"

	"user_task_apps/pkg/utils/constants"
)

type (
	Users struct {
		ID        uint   `gorm:"primaryKey;autoIncrement"`
		Name      string `gorm:"type:varchar(255)"`
		Email     string `gorm:"type:varchar(255);uniqueIndex"`
		Password  string `gorm:"type:varchar(255)"`
		CreatedAt time.Time
		UpdatedAt *time.Time
		// ID        uint `gorm:"column:id;primary_key;auto_increment"`
		// Email     string
		// Name     string

		// CreatedAt time.Time  `gorm:"column:created_at"`
		// UpdatedAt *time.Time `gorm:"column:updated_at"`
	}

	UsersRes struct {
		ID        uint   `json:"id"`
		Email     string `json:"email"`
		Name      string `json:"name"`
		Password  string `json:"password"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	UsersReq struct {
		Email    string `json:"email" binding:"required" validate:"required,email"`
		Name     string `json:"name" validate:"required" binding:"required"`
		Password string `json:"password" validate:"required" binding:"required"`
	}
)

func (Users) TableName() string {
	return "users"
}

func (users Users) ToUsersRes() *UsersRes {
	result := &UsersRes{
		ID:        users.ID,
		Email:     users.Email,
		Name:      users.Name,
		Password:  users.Password,
		CreatedAt: users.CreatedAt.Format(constants.LayoutYMD),
	}

	if users.UpdatedAt != nil {
		result.UpdatedAt = users.UpdatedAt.Format(constants.LayoutYMD)
	} else {
		result.UpdatedAt = ""
	}

	return result
}
