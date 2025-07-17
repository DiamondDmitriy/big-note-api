package task

import "time"

type Task struct {
	Id          int        `json:"id"`
	UserId      string     `json:"userId"`
	Name        string     `json:"name"`
	CategoryId  int        `json:"categoryId"`
	Order       int        `json:"order"`
	Done        bool       `json:"done"`
	DateCreated *time.Time `json:"dateCreated"`
	DateDone    *time.Time `json:"dateDone"`
}
