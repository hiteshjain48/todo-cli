package model

import "time"

type Tasks struct {
	Tasks []Task  `json:"tasks"`
}

type Task struct {
	ID 			int 		`json:"id"`
	Description string 		`json:"description"`
	Status 		Status 		`json:"status"`
	CreatedAt	time.Time	`json:"createdAt"`
	UpdatedAt	time.Time	`json:"updatedAt"`
}

func (ts *Tasks) GetNextID() int {
    maxId := 0
    for _, t := range ts.Tasks {
        if t.ID > maxId {
            maxId = t.ID
        }
    }
    return maxId + 1
}