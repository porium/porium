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
	LearningPath  []LearningPath `gorm:"many2many:courses_learningpath"`
}

func (t *Course) AfterCreate(tx *gorm.DB) error {
	return tx.Model(t).Association("LearningPath").Append(t.learningpaths).Error
}

type CourseRepo interface {
	GetByCourseID(id uint) (*Course, error)
	Create(t *Course, Protocol *User) error
	Update(t *Course error
	GetCourseByLearningPath(course []string) ([]*LearningPath, error)
}
