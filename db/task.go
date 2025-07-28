package db

import (
	"context"
	"time"
)

type Task struct{}

var TaskRepository = Task{}

type PostTaskPayload struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Status  string `json:"status" binding:"required"`
}

func (t Task) SaveTaskQuery(payload PostTaskPayload) (int, error) {
	var id int

	query := "Insert into tasks (title, content, status) values ($1, $2, $3) RETURNING id;"

	if err := DB.QueryRow(context.Background(), query, payload.Title, payload.Content, payload.Status).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

type TaskType struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (t Task) ReadTask() ([]TaskType, error) {
	var tasks []TaskType

	query := "Select * FROM tasks ORDER BY created_at DESC LIMIT 10;"

	rows, err := DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task TaskType
		if err = rows.Scan(&task.ID, &task.Title, &task.Content, &task.Status, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

type UpdateTaskPayload struct {
	ID        int       `json:"id" binding:"required"`
	Title     string    `json:"title" binding:"max=100"`
	Content   string    `json:"content" binding:"max=1000"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (t Task) GetTaskByID(id int) (TaskType, error) {
	var task TaskType

	query := "SELECT * FROM tasks WHERE id = $1;"
	row := DB.QueryRow(context.Background(), query, id)

	if err := row.Scan(&task.ID, &task.Title, &task.Content, &task.Status, &task.CreatedAt); err != nil {
		return TaskType{}, err
	}

	return task, nil
}

func (t Task) UpdateTask(payload UpdateTaskPayload) error {
	query := "UPDATE tasks SET title = $1, content = $2, status = $3 WHERE id = $4;"
	_, err := DB.Exec(context.Background(), query, payload.Title, payload.Content, payload.Status, payload.ID)
	return err
}

func (t Task) DeleteTaskQuery(id int) error {
	query := "DELETE FROM tasks WHERE id = $1;"
	_, err := DB.Exec(context.Background(), query, id)
	return err
}
