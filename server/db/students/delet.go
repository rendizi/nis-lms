package db

import "lms/db"

func DeleteStudent(student string) error {
	query := `DELETE FROM students WHERE login = $1`
	_, err := db.Db.Exec(query, student)
	if err != nil {
		return err
	}
	return nil
}
