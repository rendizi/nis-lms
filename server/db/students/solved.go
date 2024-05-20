package db

// When solved something update stats
import "lms/db"

func (s *Student) Solved(difficulty string) error {
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

	_, err := db.Db.Exec(`UPDATE students SET solved = solved + 1 and rating = rating + $1 WHERE login = $2`, bonus, s.Login)
	if err != nil {
		return err
	}
	s.Stats.Solved += 1
	s.Stats.Rating += bonus
	return nil
}
