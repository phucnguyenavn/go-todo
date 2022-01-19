package models

type User struct {
	ID       int64    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Todo     []Todo   `json:"todos" gorm:"foreignKey:TodoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Roles    []*Roles `gorm:"many2many:user_roles;"`
}
