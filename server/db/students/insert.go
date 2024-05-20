package db

import (
	"errors"
	"lms/db"
	"lms/internal/encryption"
)

func (s *Student) Insert() error {
	if s.Login == "" || s.Password == "" {
		return errors.New("No data provided. Please, provide username and system password")
	}

	encrypted, err := encryption.Generate(s.Password)
	if err != nil {
		return err
	}

	_, err = db.Db.Exec(`
		INSERT INTO students (login, password, email, klass, parallel, school,solved,rating)
		VALUES ($1, $2, $3, $4, $5, $6,$7,$8);

`, s.Login, encrypted, s.Email, s.Klass, s.Parallel, s.School, 0, 0)

	if err != nil {
		return err
	}

	return nil
}
