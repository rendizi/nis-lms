package tasks

import (
	"encoding/json"
	"fmt"
	"lms/db"
)

func (t *Task) Post() (int, error) {
	testsJSON, err := json.Marshal(t.Tests)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal tests: %w", err)
	}

	query := "INSERT INTO tasks (title, description, author, difficulty, tests, firstExample"
	args := []interface{}{t.Title, t.Description, t.Author, t.Difficulty, string(testsJSON), t.FirstExample}

	if t.Image != "" {
		query += ", image"
		args = append(args, t.Image)
	}

	if t.SecondExample != "" {
		query += ", secondExample"
		args = append(args, t.SecondExample)
	}

	if t.ThirdExample != "" {
		query += ", thirdExample"
		args = append(args, t.ThirdExample)
	}

	query += ") VALUES ("
	for i := range args {
		if i > 0 {
			query += ", "
		}
		query += fmt.Sprintf("$%d", i+1)
	}
	query += ") RETURNING id"

	var id int64
	err = db.Db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
