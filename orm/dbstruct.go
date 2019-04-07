package orm


type JobStack struct{
	Id string `gorm:"id"`
	Info *Job `gorm:"info"`
}

type JobStackArray []JobStack

type Job struct{
	Title string  `gorm:"title"`
	Short_desc string `gorm:"shortdesc"`
	Coordinates *Coordinates `gorm:"coordinates"`
	Contact string `gorm:"contact"`
	MetaData *JobMeta `gorm:"meta"`
}

type JobMeta struct{
	Added_date  string `gorm:"added_date"`
	Added_user  string `gorm:"added_user"`
	Modified_date string `gorm:"modified_date"`
	Views string `gorm:"views"`
}

type Coordinates struct{
	Latitude string `gorm:"latitude"`
	Longtitude string `gorm:"longtitude"`
}

type User struct{
	UserName string `gorm:"username"`
	Password string `gorm:"password"`
	FirstName string `gorm:"firstname"`
	LastName string `gorm:"lastname"`
	UserType string `gorm:"usertype"`
	UserStatus string `gorm:"userstatus"`
}

type UserCredentials struct{
	UserName string `gorm:"username"`
	Password string `gorm:"password"`
}