package mentee

import (
	"alta/immersive-dashboard-api/features/mentees/logs/data"
	"time"
)

type Core struct {
	Id        uint
	CreatedAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
	FullName        string  `validation:"required"`
	NickName        string 	
	ClassID         uint		`validation:"required"`
	Status          string 	`validation:"required"`
	Category        string 	
	Gender          string 
	Graduate        string 
	Mayor           string 
	Phone           string 
	Telegram        string 
	Discord         string 
	Institusi				string
	Email           string 
	EmergencyName   string 
	EmergencyPhone  string 
	EmergencyStatus string 
	Logs						[]data.MenteeLogs `gorm:"foreignKey:MenteeID"`
}

type RequestCore struct {
	FullName        string
	NickName        string
	ClassID					uint
	Status					string
	Category				string
	Gender          string 
	Graduate        string 
	Mayor           string 
	Phone           string 
	Telegram        string 
	Discord         string 
	Institusi				string
	Email           string 
	EmergencyName   string 
	EmergencyPhone  string 
	EmergencyStatus string
}

func RequestToCoreMentee(data RequestCore) Core {
	return Core {
		FullName: data.FullName,
		NickName: data.NickName,
		ClassID: data.ClassID,
		Status: data.Status,
		Category: data.Category,
		Gender: data.Gender,
		Graduate: data.Graduate,
		Mayor: data.Mayor,
		Phone: data.Phone,
		Telegram: data.Telegram,
		Discord: data.Discord,
		Institusi: data.Institusi,
		Email: data.Email,
		EmergencyName: data.EmergencyName,
		EmergencyPhone: data.EmergencyPhone,
		EmergencyStatus: data.EmergencyStatus,
	}
}

type MenteeDataInterface interface {
	Insert(data Core) (menteeId uint, err error)
	Select(menteeId uint) (mentee *Core, err error)
	SelectAll() (mentees []Core, err error)
	Update(menteeId uint, data Core) (mentee *Core, err error)
	Delete(menteeId uint) error
}

type MenteeServiceInterface interface {
	AddMentee(data Core) (menteeId uint, err error)
	GetMenteeById(menteeId uint) (mentee *Core, err error)
	GetMentees() (mentees []Core, err error)
	EditMentee(menteeId uint, data Core) (mentee *Core, err error)
	DeleteMentee(menteeId uint) error

}