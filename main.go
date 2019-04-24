package main

import (
//	"./cmd"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"./ops"
	jwt "github.com/dgrijalva/jwt-go"
	"fmt"
)

var mySigningKey = []byte("mysupersecretphrase")

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if r.Header["Token"] != nil{
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token)(interface{}, error){
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil{
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid{
				endpoint(w, r)
			}

		} else {
			fmt.Fprint(w, "Not Authorized")
		}
	})
}


func main() {
    //cmd.Execute()
    router := mux.NewRouter()
    // job operations
    router.HandleFunc("/jobs", ops.SelectAllJobs).Methods("GET")
    router.HandleFunc("/jobs/{id}", ops.GetJobById).Methods("GET")
    router.HandleFunc("/jobs/title/{title}", ops.GetJobByTitle).Methods("GET")
    router.HandleFunc("/jobs/add", ops.InsertJob).Methods("POST")
    router.HandleFunc("/jobs/delete/{id}", ops.DeleteJob).Methods("DELETE")
    // user operations
    router.HandleFunc("/users", ops.SelectAllUsers).Methods("GET")
    router.HandleFunc("/users/{id}", ops.GetUserById).Methods("GET")
    router.HandleFunc("/users/firstname/{first_name}", ops.GetUserByFirstName).Methods("GET")
    router.HandleFunc("/users/username/{user_name}", ops.GetUserByUserName).Methods("GET")
    router.HandleFunc("/users/add", ops.InsertUser).Methods("POST")
    router.HandleFunc("/users/delete/{id}", ops.DeleteUser).Methods("DELETE")
    router.HandleFunc("/users/login", ops.LogMeIn).Methods("POST")
    // initialize the database
    router.HandleFunc("/init", ops.InitializeDatabase).Methods("GET")
    ops.LogPrint("Listening port :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
