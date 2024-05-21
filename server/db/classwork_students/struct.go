package classwork_students

import (
	"strconv"
	"strings"
)

type ClassWork struct {
	Id          int    `json:"id"`
	Student     int    `json:"student"`
	Teacher     int    `json:"teacher"`
	Tasks       []int  `json:"tasks"`
	Deadline    string `json:"deadline"`
	Title       string `json:"title"`
	Desctiption string `json:"description"`
	Students    []int  `json:"students"`
}

func FromPqToGo(tasks string) ([]int, error) {
	tasks = strings.Trim(tasks, "{}")
	taskStrings := strings.Split(tasks, ",")

	taskSlice := make([]int, len(taskStrings))
	for i, taskStr := range taskStrings {
		taskInt, err := strconv.Atoi(taskStr)
		if err != nil {
			return nil, err
		}
		taskSlice[i] = taskInt
	}
	return taskSlice, nil
}
