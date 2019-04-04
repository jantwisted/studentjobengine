package ops

import(
	"encoding/json"
	"fmt"
	"net/http"
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

func AddJob(w http.ResponseWriter, r *http.Request) {
	pool := getRedisPool()
	con := pool.Get()
	defer con.Close()
	var job Job
	_ = json.NewDecoder(r.Body).Decode(&job)
	index, err := SeqNextVal(con, "jobseq")
	LogError(err)
	SaveToRedis(con, job, index)
}

func AddUser(w http.ResponseWriter, r *http.Request){
	pool := getRedisPool()
	con := pool.Get()
	defer con.Close()
	var user User
	_ = json.NewDecorder(r.Body).Decode(&user)
	index, err := SeqNextVal(con, "userseq")
	LogError(err)
}
