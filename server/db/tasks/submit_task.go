package tasks

import "lms/db"

func Submit(taskId int, login string, code string) error {
	query := `INSERT INTO solutions (student_id, task_id, solution) VALUES ((SELECT id FROM students WHERE login = $1), $2, $3);`
	_, err := db.Db.Exec(query, login, taskId, code)
	if err != nil {
		return err
	}
	return nil
}
