package logs

import "time"

type Core struct {
	Id        uint
	Status    string
	Feedback  string
	MenteeID  uint
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type LogsDataInterface interface{
	insert(userId uint) error
}

type LogsServiceInterface interface{
	Add(userId uint) error
}