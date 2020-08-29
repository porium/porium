package models

type LearningPath struct {
	Name   string  `gorm:"primary_key;auto_increment:false"`
	Courses []Course `gorm:"many2many:courses_learningpath"`
	Users  []User  `gorm:"many2many:users_learningpath"`
}

type LearningPathRepo interface {
	GetByName(name string) (*LearningPath, error)
	Create(i *LearningPath) error
	Update(i *LearningPath) error
}
