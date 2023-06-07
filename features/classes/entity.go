package classes

import (
	"alta/immersive-dashboard-api/features/mentees/mentee"
	"time"
)

type Core struct {
	Id        uint
	CreatedAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
	Name    string
	Tag     string
	UserID uint
	Mentees []mentee.Core
}

type ClassDataInterface interface {
	Insert(input Core,UserId int) error
	Update(id int, input Core) error
	Deleted(id int) error
	SelectAll()([]Core,error)

}

type ClassServiceInterface interface {
	Create(input Core,UserId int) error
	Edit(id int, input Core) error
	Deleted(id int) error
	GetAll()([]Core,error)
}