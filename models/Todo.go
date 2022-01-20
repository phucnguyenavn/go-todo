package models

type Todo struct {
	ID     int64     `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Task  string    `json:"task"`
	IsDone bool      `json:"isDone"`
}

func (todo *Todo) name() string {
	return "todo"
}
