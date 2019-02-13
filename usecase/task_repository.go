package usecase

import "github.com/sanshirookazaki/echo-clean/domain"

type TaskRepository interface {
	GetTaskAll(userid int) domain.Tasks
	GetTask(id int) domain.Tasks
	TaskHistory(userid int) domain.Tasks
	AddTask(userid int, task string)
	DeleteTask(id int)
	FinishTask(id int)
	EditTask(task string, id int)
}
