package classwork_students

import (
	"lms/db"
)

func GetByTeacher(login string, page, pageSize int) ([]ClassWork, error) {
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * pageSize

	query := `
        SELECT c.id, c.title, c.description, c.deadline, c.tasks_id
        FROM classwork c
        JOIN teachers t ON c.teacher_id = t.id
        WHERE t.login = $1
        ORDER BY c.id
        LIMIT $2 OFFSET $3;
    `
	rows, err := db.Db.Query(query, login, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classWorkList []ClassWork

	for rows.Next() {
		var classWorkInfo ClassWork
		var tasks string
		err := rows.Scan(&classWorkInfo.Id, &classWorkInfo.Title, &classWorkInfo.Desctiption, &classWorkInfo.Deadline, &tasks)
		if err != nil {
			return nil, err
		}
		taskSlice, err := FromPqToGo(tasks)
		if err != nil {
			return nil, err
		}
		classWorkInfo.Tasks = taskSlice
		classWorkList = append(classWorkList, classWorkInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return classWorkList, nil
}
