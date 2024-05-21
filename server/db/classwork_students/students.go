package classwork_students

import (
	"lms/db"
)

type StudentInfo struct {
	StudentID int    `json:"student_id"`
	Login     string `json:"login"`
	Klass     string `json:"klass"`
	School    string `json:"school"`
}

func GetStudentInfo(classworkID int) ([]StudentInfo, error) {
	query := `
        SELECT cs.student_id, s.login, s.klass, s.school
        FROM classwork_students cs
        JOIN students s ON cs.student_id = s.id
        WHERE cs.classwork_id = $1;
    `

	rows, err := db.Db.Query(query, classworkID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var studentInfoList []StudentInfo

	for rows.Next() {
		var studentInfo StudentInfo
		err := rows.Scan(&studentInfo.StudentID, &studentInfo.Login, &studentInfo.Klass, &studentInfo.School)
		if err != nil {
			return nil, err
		}
		studentInfoList = append(studentInfoList, studentInfo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return studentInfoList, nil
}
