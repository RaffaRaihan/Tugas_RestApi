package models

type Users struct {
	Id 		int64 	`gorm:"primaryKey" json:"id"`
	Name 	string 	`gorm:"type:varchar(255)" json:"name"`
}