package models

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model
	CID           string `gorm:"notnull"`
	Title         string
	ProtocolID    uint
	Protocol      User
	LevelID 	  uint
	LevelID	      *Course
	nForks        uint         `gorm:"default:0"`
	Forks         []Course      `gorm:"foreignkey:LevelID"`
	Path	      []Path `gorm:"many2many:courses_path"`
}

func (t *Course) AfterCreate(tx *gorm.DB) error {
	return tx.Model(t).Association("Path").Append(t.paths).Error
}

type CourseRepo interface {
	GetByCoursekID(id uint) (*Course, error)
	Create(t *Course, Protocol *User) error
	Update(t *Course error
	GetCourseByPath(course []string) ([]*Path, error)
}
