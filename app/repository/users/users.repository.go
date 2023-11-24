package users

import (
	"context"
	"log"
	"time"

	"user_task_apps/app/models"
	gormdb "user_task_apps/pkg/gorm.db"
)

func GetAllUsers(ctx context.Context, users *[]models.Users) error {
	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}
	err = db.Debug().Table("users").Find(&users).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUsersById(ctx context.Context, id string, users *models.Users) error {
	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}

	err = db.Debug().Table("users").First(&users, id).Error

	if err != nil {
		return err
	}

	return nil
}

func CreateUsers(ctx context.Context, usersRequest *models.UsersReq) (*models.Users, error) {
	db, err := gormdb.GetInstance()
	if err != nil {
		return nil, err
	}

	users := &models.Users{
		Email:     usersRequest.Email,
		Name:      usersRequest.Name,
		Password:  usersRequest.Password,
		CreatedAt: time.Now(),
	}

	err = db.Debug().Table("users").Create(&users).Error
	if err != nil {
		return nil, err
	}

	log.Println("log users : ", users)
	return users, nil
}

func UpdateUsers(ctx context.Context, id string, users *models.Users, usersRequest *models.UsersReq) (*models.UsersRes, error) {
	db, err := gormdb.GetInstance()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	dataUpdates := models.Users{
		Email:     usersRequest.Email,
		Name:      usersRequest.Name,
		Password:  usersRequest.Password,
		UpdatedAt: &now,
	}
	err = db.Debug().Table("users").Where("id = ?", id).Updates(&dataUpdates).Error
	if err != nil {
		return nil, err
	}

	err = db.Debug().Table("users").Where("id = ? ", id).First(&users).Error
	if err != nil {
		return nil, err
	}
	return users.ToUsersRes(), err
}

func DeleteUsers(id string, users *models.Users) error {
	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}

	err = db.Debug().Table("users").Delete(&users).Error
	if err != nil {
		return err
	}
	return nil
}
