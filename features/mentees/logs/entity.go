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

type LogsDataInterface interface {
	Insert(input Core, userId uint) (uint,error)
	Update(input Core,id uint) error
	SelectById(id uint)error
}

type LogsServiceInterface interface {
	Add(input Core, userId uint) (uint,error)
	Edit(input Core, id uint) error
	GetById(id uint)error
}
