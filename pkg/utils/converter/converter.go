package converter

import "user_task_apps/app/models"

func MapTasksToTasksRes(tasks []models.Tasks) (tasksRes []models.TasksRes) {
	for _, v := range tasks {
		tasksRes = append(tasksRes, *v.ToTasksRes())
	}
	return
}

func MapUsersToUsersRes(users []models.Users) (usersRes []models.UsersRes) {
	for _, v := range users {
		usersRes = append(usersRes, *v.ToUsersRes())
	}
	return
}
