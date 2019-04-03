package main

import (
//	"./cmd"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/jantwisted/studentjobengine/ops"
)

func main() {
    //cmd.Execute()
    router := mux.NewRouter()
    router.HandleFunc("/jobs", ops.GetAllJobs).Methods("GET")
    router.HandleFunc("/jobs/add", ops.AddJob).Methods("POST")
    ops.LogPrint("Listening port :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
