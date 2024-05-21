package db

import (
	"database/sql"
	"fmt"
	"lms/db"
)

func GetTop10ByRating() ([]Student, error) {
	query := `
		SELECT login, email, klass, parallel, school, solved, rating
		FROM students
		ORDER BY rating DESC
		LIMIT 10;
	`

	rows, err := db.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student

	for rows.Next() {
		var student Student
		err = rows.Scan(
			&student.Login,
			&student.Email,
			&student.Klass,
			&student.Parallel,
			&student.School,
			&student.Stats.Solved,
			&student.Stats.Rating,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func Place(username string) (int, error) {
	var rating int
	err := db.Db.QueryRow("SELECT rating FROM students WHERE login = $1", username).Scan(&rating)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("student with username %s not found", username)
		}
		return 0, err
	}

	var place int
	query := `
		SELECT COUNT(*)
		FROM students
		WHERE rating > $1;
	`

	err = db.Db.QueryRow(query, rating).Scan(&place)
	if err != nil {
		return 0, err
	}

	return place + 1, nil
}
