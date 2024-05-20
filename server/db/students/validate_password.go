package db

import (
	"lms/db"
	"lms/internal/encryption"
)

// Validate password
func (s *Student) Validate() error {
	var storedPassword string
	err := db.Db.QueryRow(`SELECT password FROM students WHERE login = $1`, s.Login).Scan(&storedPassword)
	if err != nil {
		return err
	}
	err = encryption.Compare(storedPassword, s.Password)
	if err != nil {
		return err
	}
	return nil
}
