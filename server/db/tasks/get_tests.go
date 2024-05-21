package tasks

import (
	"encoding/json"
	"lms/db"
)

func GetTests(id int) ([]Test, error) {
	quert := `SELECT tests FROM tasks WHERE id = $1`
	var storedTests string
	err := db.Db.QueryRow(quert, id).Scan(&storedTests)
	if err != nil {
		return nil, err
	}
	var tests []Test
	err = json.Unmarshal([]byte(storedTests), &tests)
	if err != nil {
		return nil, err
	}
	return tests, nil
}
