package ops

import(
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	orm "../orm"
	auth "../auth"
)


// jobs specific operations

func SelectAllJobs(w http.ResponseWriter, r *http.Request) {
	db := orm.Connect_To_Database()
	var job_array []orm.Job
	job_array = orm.Get_All_Jobs(db)
	json.NewEncoder(w).Encode(job_array)
}

func InitializeDatabase(w http.ResponseWriter, r *http.Request){
	db := orm.Connect_To_Database()
	orm.Database_Migration(db)
}

func InsertJob(w http.ResponseWriter, r *http.Request){
	db := orm.Connect_To_Database()
	var job orm.Job
	_ = json.NewDecoder(r.Body).Decode(&job)
	orm.Insert_To_Job(job, db)
}

func GetJobById(w http.ResponseWriter, r *http.Request){
	db := orm.Connect_To_Database()
	params := mux.Vars(r)
	idx, err := strconv.Atoi(params["id"])
	LogError(err)
	var job orm.Job
	job = orm.Search_Job_From_Idx(idx, db)
	json.NewEncoder(w).Encode(job)
}

func GetJobByTitle(w http.ResponseWriter, r *http.Request){
	db := orm.Connect_To_Database()
	params := mux.Vars(r)
	title := params["title"]
	var job_array []orm.Job
	job_array = orm.Search_Job_From_Title(title, db)
	json.NewEncoder(w).Encode(job_array)
}

func DeleteJob(w http.ResponseWriter, r *http.Request){
	db := orm.Connect_To_Database()
	params := mux.Vars(r)
	idx, err := strconv.Atoi(params["id"])
	LogError(err)
	orm.Delete_Job(idx, db)
}

// user specific operations

func SelectAllUsers(w http.ResponseWriter, r *http.Request) {
	db := orm.Connect_To_Database()
	var user_array []orm.User
	user_array = orm.Get_All_Users(db)
	json.NewEncoder(w).Encode(user_array)
}

func InsertUser(w http.ResponseWriter, r *http.Request){
	db := orm.Connect_To_Database()
	var user orm.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.Password = auth.PasswordHash([]byte(user.Password))
	orm.Insert_To_User(user, db)
}

func GetUserById(w http.ResponseWriter, r *http.Request){
	db := orm.Connect_To_Database()
	params := mux.Vars(r)
	idx, err := strconv.Atoi(params["id"])
	LogError(err)
	var user orm.User
	user = orm.Search_User_From_Idx(idx, db)
	json.NewEncoder(w).Encode(user)
}

func GetUserByFirstName(w http.ResponseWriter, r *http.Request){
	db := orm.Connect_To_Database()
	params := mux.Vars(r)
	first_name := params["first_name"]
	var user_array []orm.User
	user_array = orm.Search_User_From_First_Name(first_name, db)
	json.NewEncoder(w).Encode(user_array)
}

func GetUserByUserName(w http.ResponseWriter, r *http.Request){
	db := orm.Connect_To_Database()
	params := mux.Vars(r)
	user_name := params["user_name"]
	var user orm.User
	user = orm.Search_User_From_User_Name(user_name, db)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	db := orm.Connect_To_Database()
	params := mux.Vars(r)
	idx, err := strconv.Atoi(params["id"])
	LogError(err)
	orm.Delete_User(idx, db)
}

func LogMeIn(w http.ResponseWriter, r *http.Request){
	db := orm.Connect_To_Database()
	var user_login_request orm.UserLoginRequest
	var user_login_response orm.UserLoginResponse
	_ = json.NewDecoder(r.Body).Decode(&user_login_request)
	hashed_password_from_db := orm.GetPassword(user_login_request.UserName, db)
	ismatch := auth.ComparePasswords(hashed_password_from_db, []byte(user_login_request.Password))
	user_login_response.UserName = user_login_request.UserName
	if ismatch {
		token, err := auth.GenerateJWT(user_login_request.UserName)
		user_login_response.JWT = token
		LogError(err)
		json.NewEncoder(w).Encode(user_login_response)
	} else {
		user_login_response.JWT = "failed!"
		json.NewEncoder(w).Encode(user_login_response)
	}
}
