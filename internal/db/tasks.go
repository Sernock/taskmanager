package db

import (
	"log"
	"taskmanager/internal/models"
)

func CreateTask(task models.Tasks) {
	query := `INSERT into tasks(title, description, completed) VALUES (?, ?, ?)`
	_, err := DB.Exec(query, task.Title, task.Description, task.Completed)
	if err != nil {
		log.Println("Failed to create new task:", err)
	}
}

func GetTask() []models.Tasks {
	rows, err := DB.Query("SELECT id, title, description, completed FROM tasks")
	if err != nil {
		log.Println("Failed to fetch tasks:", err)
		return nil
	}
	defer rows.Close()

	var tasks []models.Tasks
	for rows.Next() {
		var t models.Tasks
		err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Completed)
		if err != nil {
			log.Panicln("Failed to scan task:", err)
			continue
		}
		tasks = append(tasks, t)
	}

	return tasks
}

func GetTaskByID(id int) (models.Tasks, error) {
	var t models.Tasks
	query := "SELECT id, title, description, completed FROM tasks WHERE id = ?"
	err := DB.QueryRow(query, id).Scan(&t.Id, &t.Title, &t.Description, &t.Completed)
	if err != nil {
		return t, err
	}
	return t, nil
}

func UpdateTask(task models.Tasks) {
	query := `UPDATE tasks SET title=?, description=?, completed=? WHERE id=?`
	_, err := DB.Exec(query, task.Title, task.Description, task.Completed, task.Id)
	if err != nil {
		log.Println("Failed to update task:", err)
	} 
}

func DeleteTask(id int) {
	query := `DELETE FROM tasks WHERE id=?`
	_, err := DB.Exec(query, id)
	if err != nil {
		log.Panicln("Failed to delete task:", err)
	}
} 

