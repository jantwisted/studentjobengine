package ops

import(
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	orm "../orm"
)



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
