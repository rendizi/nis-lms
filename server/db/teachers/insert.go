package teachers

import (
	"lms/db"
	d "lms/db/students"
)

func Insert(creds d.Student) error {
	query := `INSERT INTO teachers (login,password,email,school) VALUES ($1,$2,$3,$4)`
	_, err := db.Db.Exec(query, creds.Login, creds.Password, creds.Email, creds.School)
	if err != nil {
		return err
	}
	return nil
}
