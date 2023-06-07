package mentee

import "time"

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
	// Logs						[]data.MenteeLogs `gorm:"foreignKey:MenteeID"`
}

type MenteeDataInterface interface {
	Insert(data Core) (menteeId uint, err error)
	Select(menteeId uint) (mentee Core, err error)
	SelectAll() (mentees []Core)
	Update(menteeId uint, data Core) (mentee Core, err error)
	Delete(menteeId uint) error
}

type MenteeServiceInterface interface {
	AddMentee(data Core) (menteeId uint, err error)
	GetMenteeById(menteeId uint) (mentee Core, err error)
	GetMentees() (mentees []Core)
	EditMentee(menteeId uint, data Core) (mentee Core, err error)
	DeleteMentee(menteeId uint) error
}