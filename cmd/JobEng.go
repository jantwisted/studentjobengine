package main

import (
  //  "fmt"
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "github.com/go-redis/redis"
)


type Job struct{
  JID string `json:jid,omitempty`
  Title string  `json:"title,omitempty"`
  Short_desc string `json:"shortdesc,omitempty"`
  Coordinates *Coordinates `json:"coordinates,omitempty"`
  Contact string `json:"contact,omitempty"`
  MetaData *JobMeta `json:"meta,omitempty"`
}

type JobMeta struct{
  Added_date  string `json:added_date,omitempty`
  Added_user  string `json:added_user,omitempty`
  Modified_date string `json:modified_date,omitempty`
  Views string `json:views,omitempty`
}

type Coordinates struct{
  Latitude string `json:latitude,omitempty`
  Longtitude string `json:longtitude,omitempty`
}

var alljobs []Job


func GetAllJobs(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(alljobs)
//    client := RedisCon()
//    pong, err := client.Ping().Result()
//    fmt.Println(pong, err)
}


func GetJob(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range alljobs{
        if item.JID == params["jid"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Job{})
}


func AddJob(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var job Job
    _ = json.NewDecoder(r.Body).Decode(&job)
    job.JID = params["jid"]
    alljobs = append(alljobs, job)
    json.NewEncoder(w).Encode(alljobs)
}


func RemoveJob(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range alljobs {
        if item.JID == params["jid"] {
            alljobs = append(alljobs[:index], alljobs[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(alljobs)
    }
}

func RedisCon() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB:	0,
	})
	return client
}

func main() {
    router := mux.NewRouter()
//    alljobs = append(alljobs, Job{JID: "1", Title: "John", Short_desc: "Doe", Coordinates: &Coordinates{Latitude: "0.1234", Longtitude: "0.5678"}, Contact: "janith@tuta.io", MetaData:&JobMeta{Added_date:"11", Added_user:"Jan", Modified_date:"11", Views:"5"}})
    router.HandleFunc("/jobs", GetAllJobs).Methods("GET")
    router.HandleFunc("/jobs/{id}", GetJob).Methods("GET")
    router.HandleFunc("/jobs/add", AddJob).Methods("POST")
    router.HandleFunc("/jobs/{id}", RemoveJob).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8080", router))
}
