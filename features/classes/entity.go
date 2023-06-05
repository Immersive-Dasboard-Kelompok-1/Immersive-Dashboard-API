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
	Insert(input Core) error
	Update(id int,input Core) error
	Deleted(id int) error

}

type ClassServiceInterface interface {
	Create(input Core) error
	Edit(id int,input Core) error
	Deleted(id int) error
}