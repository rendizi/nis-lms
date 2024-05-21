package db

type Student struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Klass    string `json:"klass"`
	Parallel int    `json:"parallel"`
	School   string `json:"school"`
	Stats    struct {
		Solved   int    `json:"solved"`
		Leetcode string `json:"leetcode"`
		Badges   string `json:"badges"`
		Rating   int    `json:"rating"`
		Rank     int    `json:"rank"`
	} `json:"stats"`
}
