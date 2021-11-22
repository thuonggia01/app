package models

import "time"

type User struct {
	ID          uint `gorm:"primary_key"`
	Email       string
	Password    string
	Status      bool
	FullName    string
	Phone       string
	Since_Date  time.Time
	Description string
	Facebook    string
	Part_Id     int
	Position_Id int
	Branch_Id   int
	Image_Id    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
