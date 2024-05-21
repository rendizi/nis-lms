package classwork_students

import (
	"github.com/lib/pq"
	"lms/db"
)

func (c *ClassWork) Insert(login string) (int, error) {
	query := `
        INSERT INTO classwork (title, description, teacher_id, deadline, tasks_id)
        VALUES ($1, $2, (SELECT id FROM teachers WHERE login = $3), $4, $5)
        RETURNING id;
    `
	taskArray := pq.Array(c.Tasks)
	var classWorkId int64
	err := db.Db.QueryRow(query, c.Title, c.Desctiption, login, c.Deadline, taskArray).Scan(&classWorkId)
	if err != nil {
		return 0, err
	}

	tx, err := db.Db.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	stmt, err := tx.Prepare(`
        INSERT INTO classwork_students (classwork_id, student_id)
        VALUES ($1, $2);
    `)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	for _, student := range c.Students {
		_, err = stmt.Exec(classWorkId, student)
		if err != nil {
			return 0, err
		}
	}

	return int(classWorkId), nil
}
