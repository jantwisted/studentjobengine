
package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"encoding/json"
)

func Connect_To_Database() (*gorm.DB) {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=sje password=postgres")
	if err != nil{
		fmt.Println(err)
	}
	return db
}

func Database_Migration(db *gorm.DB){
	db.AutoMigrate(&Job{}, &Coordinates{}, &User{})
}


// jobs specific operations

func Insert_To_Job(job Job, db *gorm.DB){
	db.Create(&job)
}
	
func Get_All_Jobs(db *gorm.DB)([]Job){
	var job []Job
	_, err := json.Marshal(db.Find(&job))
	if err != nil{
		fmt.Println(err)
	}
	return job
}

func Search_Job_From_Idx(id int, db *gorm.DB)(Job){
	var job Job
	_, err := json.Marshal(db.First(&job, id))
	if err != nil{
		fmt.Println(err)
	}	
	return job
}

func Search_Job_From_Title(title string, db *gorm.DB)([]Job){
	var job []Job
	_, err := json.Marshal(db.Where("title LIKE ?", "%"+title+"%").Find(&job))
	if err != nil{
		fmt.Println(err)
	}
	return job
}

func Delete_Job(id int, db *gorm.DB){
	var job Job
	db.Where("id = ?", id).Delete(&job)
}

// user specific operations

func Insert_To_User(user User, db *gorm.DB){
	db.Create(&user)
}

func Get_All_Users(db *gorm.DB)([]User){
	var user []User
	_, err := json.Marshal(db.Find(&user))
	if err != nil{
		fmt.Println(err)
	}
	return user
}

func Search_User_From_Idx(id int, db *gorm.DB)(User){
	var user User
	_, err := json.Marshal(db.First(&user, id))
	if err != nil{
		fmt.Println(err)
	}
	return user
}

func Search_User_From_User_Name(user_name string, db *gorm.DB)(User){
	var user User
	_, err := json.Marshal(db.Where("user_name LIKE ?", "%"+user_name+"%").First(&user))
	if err != nil{
		fmt.Println(err)
	}
	return user
}

func Search_User_From_First_Name(first_name string, db *gorm.DB)([]User){
	var user []User
	_, err := json.Marshal(db.Where("first_name LIKE ?", "%"+first_name+"%").Find(&user))
	if err != nil{
		fmt.Println(err)
	}
	return user
}

func Delete_User(id int, db *gorm.DB){
	var user User
	db.Where("id = ?", id).Delete(&user)
}

func GetPassword(user_name string, db *gorm.DB)(string){
	user := Search_User_From_User_Name(user_name, db)
	return user.Password
}


