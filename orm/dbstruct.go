package orm

import (
	"github.com/jinzhu/gorm"
)


type JobArray []Job

type Job struct{
	gorm.Model
	JobID int `gorm:"primary_key;AUTO_INCREMENT"`
	Title string  `gorm:"type:varchar(50);unique_index"`
	Short_desc string `gorm:"type:varchar(200)`
	Location Coordinates
	Contact string `gorm:"type:varchar(100)"`
}

type Coordinates struct{
	JobID uint
	Latitude string 
	Longtitude string 
}

type User struct{
	gorm.Model
	UserID int `gorm:"primary_key;AUTO_INCREMENT"`
	UserName string `gorm:"type:varchar(50);unique_index"`
	Password string `gorm:"type:varchar(100)"`
	FirstName string `gorm:"type:varchar(100)"`
	LastName string `gorm:"type:varchar(100)"`
	UserType string `gorm:"type:varchar(50)"`
	UserStatus int 
}

