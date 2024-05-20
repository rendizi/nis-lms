package tasks

import (
	"database/sql"
	"lms/db"
)

func GetTask(id string) (Task, error) {
	var query string
	if id == "random" {
		query = `SELECT * FROM tasks ORDER BY RANDOM() LIMIT 1`
	} else {
		query = `SELECT * FROM tasks WHERE id = $1`
	}

	var task Task
	var second, third sql.NullString
	var image sql.NullString

	var row *sql.Row
	if id == "random" {
		row = db.Db.QueryRow(query)
	} else {
		row = db.Db.QueryRow(query, id)
	}

	err := row.Scan(&task.Id, &task.Title, &task.Description, &task.Author, &task.Difficulty, &task.Tests, &image, &task.FirstExample, &second, &third)
	if err != nil {
		return Task{}, err
	}

	if second.Valid {
		task.SecondExample = second.String
	}
	if image.Valid {
		task.Image = image.String
	}
	if third.Valid {
		task.ThirdExample = third.String
	}

	return task, nil
}
