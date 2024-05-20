package tasks

import (
	"fmt"
	"lms/db"
)

func (t *Task) Post() (int, error) {
	query := ""
	args := []interface{}{}
	image := true

	if t.Image == "" {
		query = `INSERT INTO tasks (title, description, author, difficulty, tests, firstExample`
		args = append(args, t.Title, t.Description, t.Author, t.Difficulty, t.Tests, t.FirstExample)

		image = true
	} else {
		query = `INSERT INTO tasks (title, description, image, author, difficulty, tests, firstExample`
		args = append(args, t.Title, t.Description, t.Image, t.Author, t.Difficulty, t.Tests, t.FirstExample)

		image = false
	}

	if t.SecondExample == "" {
		if !image {
			query += `) VALUES ($1, $2, $3, $4, $5, $6)`
		} else {
			query += `) VALUES ($1, $2, $3, $4, $5, $6, $7)`
		}
	} else if t.ThirdExample == "" {
		if !image {
			query += ", secondExample) VALUES ($1, $2, $3, $4, $5, $6,$7,$8)"
		} else {
			query += ", secondExample) VALUES ($1, $2, $3, $4, $5, $6, $7,$8,$9)"
		}
		args = append(args, t.SecondExample)
	} else {
		if !image {
			query += ", secondExample, thirdExample) VALUES ($1, $2, $3, $4, $5, $6,$7,$9)"
		} else {
			query += ", secondExample, thirdExample) VALUES ($1, $2, $3, $4, $5, $6, $7,$8,$10)"
		}
		args = append(args, t.SecondExample, t.ThirdExample)
	}

	query += " RETURNING id"
	var id int64
	fmt.Println(args)
	err := db.Db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
