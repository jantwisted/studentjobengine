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
	s, err := GetFromRedis(con, "JOB")
	if err != nil{
		panic(err.Error())
	}
	joblist := Job{}
	err = json.Unmarshal([]byte(s), &joblist)
	json.NewEncoder(w).Encode(joblist)
}

func AddJob(w http.ResponseWriter, r *http.Request) {
	pool := getRedisPool()
	con := pool.Get()
	defer con.Close()
	var job Job
	_ = json.NewDecoder(r.Body).Decode(&job)
	index, err := SeqNextVal(con, "jobseq")
	if err != nil {
		fmt.Println("sequence error!")
	}
	SaveToRedis(con, job, index)
}
