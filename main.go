package main

import (
//	"./cmd"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"./ops"
)

func main() {
    //cmd.Execute()
    router := mux.NewRouter()
    router.HandleFunc("/jobs", ops.SelectAllJobs).Methods("GET")
    router.HandleFunc("/jobs/{id}", ops.GetJobById).Methods("GET")
    router.HandleFunc("/jobs/title/{title}", ops.GetJobByTitle).Methods("GET")
    router.HandleFunc("/jobs/add", ops.InsertJob).Methods("POST")
    router.HandleFunc("/jobs/delete/{id}", ops.DeleteJob).Methods("DELETE")
    router.HandleFunc("/user/add", ops.AddUser).Methods("POST")
    router.HandleFunc("/init", ops.InitializeDatabase).Methods("GET")
    ops.LogPrint("Listening port :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
