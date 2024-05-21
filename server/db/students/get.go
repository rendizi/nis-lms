package db

import (
	"database/sql"
	"errors"
	"lms/db"
)

func (s *Student) GetInfo() error {
	query := `SELECT id, login, email, klass, parallel, school, solved, leetcode, badges, rating FROM students WHERE login = $1`
	var leetcode, badges sql.NullString
	err := db.Db.QueryRow(query, s.Login).Scan(
		&s.Id, &s.Login, &s.Email, &s.Klass, &s.Parallel, &s.School,
		&s.Stats.Solved, &leetcode, &badges, &s.Stats.Rating,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("student not found")
		}
		return err
	}

	if leetcode.Valid {
		s.Stats.Leetcode = leetcode.String
	}
	if badges.Valid {
		s.Stats.Badges = badges.String
	}

	return nil
}
