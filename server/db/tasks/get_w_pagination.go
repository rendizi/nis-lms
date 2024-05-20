package tasks

import (
	"fmt"
	"lms/db"
	"strings"
)

type searchTaskResponse struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Difficulty string `json:"difficulty"`
}

func GetTasksWithPagination(page int, pageSize int, difficulty string, title string) ([]searchTaskResponse, error) {
	offset := (page - 1) * pageSize
	baseQuery := `SELECT id, title, author, difficulty FROM tasks`
	conditions := []string{}
	args := []interface{}{}
	argIndex := 1

	// Add filters to the query
	if difficulty != "" {
		conditions = append(conditions, fmt.Sprintf("difficulty = $%d", argIndex))
		args = append(args, difficulty)
		argIndex++
	}
	if title != "" {
		conditions = append(conditions, fmt.Sprintf("title ILIKE $%d", argIndex))
		args = append(args, "%"+title+"%")
		argIndex++
	}

	// Construct the final query
	if len(conditions) > 0 {
		baseQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	// Apply ORDER BY id ASC only if no filters are applied
	if len(conditions) == 0 {
		baseQuery += " ORDER BY id ASC"
	}

	baseQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, pageSize, offset)

	rows, err := db.Db.Query(baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []searchTaskResponse
	for rows.Next() {
		var task searchTaskResponse
		err = rows.Scan(&task.Id, &task.Title, &task.Author, &task.Difficulty)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
