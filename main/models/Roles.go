package models

type Roles struct {
	ID    int64 `gorm:"primaryKey"`
	Name  string
	Users []*User `gorm:"many2many:user_roles;"`
}
