package db

import (
	"database/sql"
	"lms/db"
)

func Search(input string, page, pageSize int) ([]Student, error) {
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * pageSize

	query := `
		SELECT *
		FROM students
		WHERE login LIKE '%' || $1 || '%'
		ORDER BY id
		LIMIT $2 OFFSET $3;
	`

	rows, err := db.Db.Query(query, input, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student

	for rows.Next() {
		var student Student
		var leetcode, badges sql.NullString
		err = rows.Scan(
			&student.Id,
			&student.Login,
			&student.Password,
			&student.Email,
			&student.Klass,
			&student.Parallel,
			&student.School,
			&student.Stats.Solved,
			&leetcode,
			&badges,
			&student.Stats.Rating)
		student.Password = ""
		if err != nil {
			return nil, err
		}
		if leetcode.Valid {
			student.Stats.Leetcode = leetcode.String
		}
		if badges.Valid {
			student.Stats.Badges = badges.String
		}
		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}
