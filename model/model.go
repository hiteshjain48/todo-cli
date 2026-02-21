package model

import "time"

type Tasks struct {
	Tasks []Task  `json:"tasks"`
}

type Task struct {
	Id 			int 		`json:"id"`
	Description string 		`json:"description"`
	Status 		string 		`json:"status"`
	CreatedAt	time.Time	`json:"createdAt"`
	UpdatedAt	time.Time	`json:"updatedAt"`
}

func (ts *Tasks) GetNextId() int {
    maxId := 0
    for _, t := range ts.Tasks {
        if t.Id > maxId {
            maxId = t.Id
        }
    }
    return maxId + 1
}