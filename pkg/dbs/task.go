package dbs

import (
	"database/sql"
)

// Task представляет структуру задачи
type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
	CreatedAt   string
}

// CreateTask создает новую задачу в базе данных
func CreateTask(db *sql.DB, task *Task) error {
	_, err := db.Exec("INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3)", task.Title, task.Description, task.Status)
	return err
}

// GetTasks возвращает все задачи из базы данных
func GetTasks(db *sql.DB) ([]Task, error) {
	rows, err := db.Query("SELECT id, title, description, status, created_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

// UpdateTask обновляет задачу в базе данных
func UpdateTask(db *sql.DB, task *Task) error {
	_, err := db.Exec("UPDATE tasks SET title = $1, description = $2, status = $3 WHERE id = $4", task.Title, task.Description, task.Status, task.ID)
	return err
}

// GetTaskByID возвращает задачу по ее ID
func GetTaskByID(db *sql.DB, id int) (*Task, error) {
	var task Task
	err := db.QueryRow("SELECT id, title, description, status, created_at FROM tasks WHERE id = $1", id).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// DeleteTask удаляет задачу из базы данных по ее ID
func DeleteTask(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}
