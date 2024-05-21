package db

// When solved something update stats
import (
	"lms/db"
)

func (s *Student) Solved(difficulty string) error {
	var err error
	bonus := 0
	switch difficulty {
	case "easy":
		bonus = 1
	case "medium":
		bonus = 3
	case "hard":
		bonus = 5
	default:
		bonus = 1
	}

	// Use transaction to ensure data consistency
	tx, err := db.Db.Begin()
	if err != nil {
		return err
	}
	//Update student table, add to solved files 1 and to rating bonus where login is s.Login
	_, err = tx.Exec(`UPDATE students SET solved = solved + 1, rating = rating + $1::int WHERE login = $2`, bonus, s.Login)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction if there were no errors
	err = tx.Commit()
	if err != nil {
		return err
	}

	s.Stats.Solved += 1
	s.Stats.Rating += bonus
	return nil
}
