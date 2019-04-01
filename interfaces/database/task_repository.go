package database

import (
	"strconv"

	"github.com/sanshirookazaki/echo-clean/domain"
)

type TaskRepository struct {
	SQLHandler
}

func (repo *TaskRepository) GetTaskAll(userid int) (tasks domain.Tasks) {
	var t domain.Tasks
	repo.Table("tasks").Select("*").Where("Status = 0 and userid =" + strconv.Itoa(userid)).Scan(&t)
	return t
}

func (repo *TaskRepository) GetTask(id int) (tasks domain.Tasks) {
	var t domain.Tasks
	repo.Table("tasks").Select("*").Where("ID =" + strconv.Itoa(id)).Scan(&t)
	return t
}

func (repo *TaskRepository) TaskHistory(userid int) (tasks domain.Tasks) {
	var t domain.Tasks
	repo.Table("tasks").Select("*").Where("Status = 1 and Userid =" + strconv.Itoa(userid)).Scan(&t)
	return t
}

func (repo *TaskRepository) AddTask(userid int, task string) {
	t := domain.Task{UserID: userid, Task: task, Status: 0}
	repo.Table("tasks").Create(&t)
	repo.Save(&t)
}

func (repo *TaskRepository) DeleteTask(id int) {
	t := domain.Task{ID: id}
	repo.Table("tasks").Delete(&t)
}

func (repo *TaskRepository) FinishTask(id int) {
	t := domain.Task{ID: id}
	repo.Model(&t).Update("Status", "1")
}

func (repo *TaskRepository) EditTask(task string, id int) {
	t := domain.Task{ID: id, Task: task}
	repo.Model(&t).Update(domain.Task{Status: 1, Task: task})
}
