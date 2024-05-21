package tasks

type Task struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Author        string `json:"author"`
	Difficulty    string `json:"difficulty"`
	Tests         []Test `json:"tests"`
	Image         string `json:"image"`
	FirstExample  string `json:"first_example"`
	SecondExample string `json:"second_example"`
	ThirdExample  string `json:"third_example"`
}

type Test struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}
