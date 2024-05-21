package teachers

import "lms/db"

func Delete(username string) error {
	query := `DELETE FROM teachers WHERE login = $1`
	_, err := db.Db.Exec(query, username)
	if err != nil {
		return err
	}
	return nil
}
