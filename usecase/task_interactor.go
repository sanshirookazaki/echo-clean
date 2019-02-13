package usecase

import "github.com/sanshirookazaki/echo-clean/domain"

type TaskInteractor struct {
	TaskRepository TaskRepository
}

func (interactor *TaskInteractor) GetTaskAll(userid int) (tasks domain.Tasks) {
	tasks = interactor.TaskRepository.GetTaskAll(userid)
	return tasks
}

func (interactor *TaskInteractor) GetTask(id int) (tasks domain.Tasks) {
	tasks = interactor.TaskRepository.GetTask(id)
	return tasks
}

func (interactor *TaskInteractor) TaskHistory(userid int) (tasks domain.Tasks) {
	tasks = interactor.TaskRepository.TaskHistory(userid)
	return tasks
}

func (interactor *TaskInteractor) AddTask(userid int, task string) {
	interactor.TaskRepository.AddTask(userid, task)
}

func (interactor *TaskInteractor) DeleteTask(id int) {
	interactor.TaskRepository.DeleteTask(id)
}

func (interactor *TaskInteractor) FinishTask(id int) {
	interactor.TaskRepository.FinishTask(id)
}

func (interactor *TaskInteractor) EditTask(task string, id int) {
	interactor.TaskRepository.EditTask(task, id)
}
