package tasks

import (
	"github.com/pkg/errors"
	"lms/db"
	"time"
)

type Solution struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Code   string `json:"code"`
	Time   string `json:"time"`
	Author string `json:"author"`
}

func Solutions(login string, page, pageSize int) ([]Solution, error) {
	// Calculate OFFSET based on page and pageSize
	offset := (page - 1) * pageSize

	query := `
	SELECT sts.id, t.title, sts.solution, sts.submission_time
	FROM solutions sts
	JOIN tasks t ON sts.task_id = t.id
	JOIN students s ON sts.student_id = s.id
	WHERE s.login = $1
	LIMIT $2 OFFSET $3;
	`

	rows, err := db.Db.Query(query, login, pageSize, offset)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch solutions from database")
	}
	defer rows.Close()

	var solutions []Solution
	for rows.Next() {
		var solution Solution
		var submissionTime time.Time
		err = rows.Scan(&solution.Id, &solution.Title, &solution.Code, &submissionTime)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan solution row")
		}
		solution.Time = submissionTime.Format(time.RFC3339)
		solutions = append(solutions, solution)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error iterating over solution rows")
	}

	return solutions, nil
}

func DidSolvedIt(login string, id int) (bool, error) {
	query := `
SELECT COUNT(*)
FROM solutions sts
JOIN students s ON sts.student_id = s.id
WHERE s.login = $1 AND sts.task_id = $2;

`
	var count int
	err := db.Db.QueryRow(query, login, id).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "failed to count solutions")
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func WhoSolvedIt(id, page, pageSize int) ([]Solution, error) {
	// Calculate OFFSET based on page and pageSize
	offset := (page - 1) * pageSize

	query := `
	SELECT sts.id, s.login, sts.solution, sts.submission_time
	FROM solutions sts
	JOIN students s ON sts.student_id = s.id
	WHERE sts.task_id = $1
	LIMIT $2 OFFSET $3;
	`
	rows, err := db.Db.Query(query, id, pageSize, offset)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch solutions from database")
	}
	defer rows.Close()

	var solutions []Solution
	for rows.Next() {
		var solution Solution
		var submissionTime time.Time
		err = rows.Scan(&solution.Id, &solution.Author, &solution.Code, &submissionTime)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan solution row")
		}
		solution.Time = submissionTime.Format(time.RFC3339)
		solutions = append(solutions, solution)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error iterating over solution rows")
	}

	return solutions, nil
}
