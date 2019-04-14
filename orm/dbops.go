
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
