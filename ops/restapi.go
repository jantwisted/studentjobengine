package ops

import(
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	orm "../orm"
)

func GetAllJobs(w http.ResponseWriter, r *http.Request) {
	pool := getRedisPool()
	con := pool.Get()
	defer con.Close()
	keyarray, err := GetKeys(con, "JOB*")
	if err != nil{
		panic(err.Error())
	}
	job_array := JobStackArray{}
	for _, keyvalue := range keyarray{
		str, err := GetJsonValue(con, keyvalue)
		LogError(err)
		jobstackstr := GetJobStackStr(keyvalue, str)
		job_array = AppendToJobStackArray(jobstackstr, job_array)
	}
	LogPrint("Sent JSON")
	json.NewEncoder(w).Encode(job_array)
}

func GetJobStackStr(jobid string, info string)(string){
	return "{\"id\":\""+jobid+"\",\"Info\":" +info+ "}"
}

func AppendToJobStackArray(jobstackstr string, job_array JobStackArray)(JobStackArray){
	jobstack := JobStack{}
	err := json.Unmarshal([]byte(jobstackstr), &jobstack)
	LogError(err)
	return append(job_array, jobstack)
}



func AddUser(w http.ResponseWriter, r *http.Request){
	const prefix string = "USR"
	pool := getRedisPool()
	con := pool.Get()
	defer con.Close()
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	index, err := SeqNextVal(con, "userseq")
	LogError(err)
	SaveToRedis(con, user, prefix, index)
}

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
