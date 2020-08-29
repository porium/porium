package models

type Path struct {
	Name   string  `gorm:"primary_key;auto_increment:false"`
	Courses []Course `gorm:"many2many:courses_path"`
	Users  []User  `gorm:"many2many:users_path"`
}

type PathRepo interface {
	GetByName(name string) (*Path, error)
	Create(i *Path) error
	Update(i *Path) error
}
