package database

import (
	"strconv"

	"github.com/sanshirookazaki/echo-clean/domain"
)

type TaskRepository struct {
	SQLHandler
}

func (repo *TaskRepository) GetTaskAll(userid int) (tasks domain.Tasks) {
	rows, err := repo.Query("SELECT * FROM tasks WHERE Status = 0 and userid = " + strconv.Itoa(userid))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var t domain.Task
	for rows.Next() {
		if err := rows.Scan(&t.ID, &t.UserID, &t.Task, &t.Status); err != nil {
			panic(err.Error())
		}
		tasks = append(tasks, t)
	}
	return tasks
}

func (repo *TaskRepository) GetTask(id int) (tasks domain.Tasks) {
	rows, err := repo.Query("SELECT * FROM tasks WHERE ID = " + strconv.Itoa(id))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var t domain.Task
	for rows.Next() {
		if err := rows.Scan(&t.ID, &t.UserID, &t.Task, &t.Status); err != nil {
			panic(err.Error())
		}
		tasks = append(tasks, t)
	}
	return tasks
}

func (repo *TaskRepository) TaskHistory(userid int) (tasks domain.Tasks) {
	rows, err := repo.Query("SELECT * FROM tasks WHERE Status = 1 and Userid =" + strconv.Itoa(userid))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var t domain.Task
	for rows.Next() {
		if err := rows.Scan(&t.ID, &t.UserID, &t.Task, &t.Status); err != nil {
			panic(err.Error())
		}
		tasks = append(tasks, t)
	}
	return tasks
}

func (repo *TaskRepository) AddTask(userid int, task string) {
	_, err := repo.Query("INSERT INTO tasks (Userid, Task, Status) VALUES ( " + strconv.Itoa(userid) + ", \"" + task + "\", 0)")
	if err != nil {
		panic(err.Error)
	}
}

func (repo *TaskRepository) DeleteTask(id int) {
	_, err := repo.Query("DELETE FROM tasks WHERE ID =" + strconv.Itoa(id))
	if err != nil {
		panic(err.Error())
	}
}

func (repo *TaskRepository) FinishTask(id int) {
	_, err := repo.Query("UPDATE tasks SET Status = 1 WHERE ID =" + strconv.Itoa(id))
	if err != nil {
		panic(err.Error())
	}
}

func (repo *TaskRepository) EditTask(task string, id int) {
	_, err := repo.Query("UPDATE tasks SET Task = \"" + task + "\" WHERE ID = " + strconv.Itoa(id))
	if err != nil {
		panic(err.Error())
	}
}
