
package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"encoding/json"
	"os"
)


func Connect_To_Database() (*gorm.DB) {
	db_type := string(os.Getenv("DBTYPE"))
	db_host := string(os.Getenv("DBHOST"))
	db_port := string(os.Getenv("DBPORT"))
	db_dbname := string(os.Getenv("DBNAME"))
	db_user := string(os.Getenv("DBUSER"))
	db_passwd := string(os.Getenv("DBPASSWD"))
	db, err := gorm.Open(db_type, "host="+db_host+" port="+db_port+" user="+db_user+" dbname="+db_dbname+" password="+db_passwd)
	if err != nil{
		fmt.Println(err)
	}
	return db
}

func Database_Migration(db *gorm.DB){
	db.AutoMigrate(&Job{}, &User{})
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
// this is a wrapper function for easy access
	user := Search_User_From_User_Name(user_name, db)
	return user.Password
}

// Special Functions

func GetJobsFromDistance(distance int, db *gorm.DB)([]Job){
// get job ids based on the distance
	var result []DistanceResult
	db.Raw("SELECT X.id FROM ( SELECT id, ( 3959 * acos( cos( radians(47.085895) ) * cos( radians( cast(latitude as float) ) ) * cos( radians( cast(longtitude as float) ) - radians(17.900233) ) + sin( radians(47.085895) ) * sin( radians( cast(latitude as float) ) ) ) ) AS distance FROM jobs ) X WHERE X.distance < ? ORDER BY X.distance", distance).Scan(&result)
	for i, r := range result {
		job := Search_User_From_Idx(r.id, db)
		fmt.Println(job.Title)
	}
}

