package models

import "time"

type Todo struct {
	ID     int64     `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Tasks  string    `json:"tasks"`
	IsDone bool      `json:"isDone"`
	at     time.Time `json:"at"`
}
