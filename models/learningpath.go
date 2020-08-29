package models

type Path struct {
	Name   string  `gorm:"primary_key;auto_increment:false"`
	Courses []Course `gorm:"many2many:tracks_instruments"`
	Paths []Path
	Users  []User  `gorm:"many2many:users_instruments"`
}

type PathRepo interface {
	GetByName(name string) (*Instrument, error)
	Create(i *Instrument) error
	Update(i *Instrument) error
}
