package main

import (
//	"./cmd"
	"github.com/gorilla/mux"
	"github.com/jantwisted/studentjobengine/ops"
	jwt "github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/urfave/negroni"
	"os"
)




func main() {

	ops.InitializeDatabase()
    	//cmd.Execute()

	router_noauth := mux.NewRouter()
	router_auth := mux.NewRouter()

	var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
  				ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
    					return []byte(string(os.Getenv("SIGNKEY"))), nil
  				},
  				SigningMethod: jwt.SigningMethodHS256,
	})

	// job operations

	router_auth.HandleFunc("/jobs", ops.SelectAllJobs).Methods("GET")
	router_auth.HandleFunc("/jobs/{id}", ops.GetJobById).Methods("GET")
	router_auth.HandleFunc("/jobs/title/{title}", ops.GetJobByTitle).Methods("GET")
	router_auth.HandleFunc("/jobs/add", ops.InsertJob).Methods("POST")
	router_auth.HandleFunc("/jobs/delete/{id}", ops.DeleteJob).Methods("DELETE")

	// user operations

	router_auth.HandleFunc("/users/all", ops.SelectAllUsers).Methods("GET")
	router_auth.HandleFunc("/users/{id}", ops.GetUserById).Methods("GET")
	router_auth.HandleFunc("/users/firstname/{first_name}", ops.GetUserByFirstName).Methods("GET")
	router_auth.HandleFunc("/users/username/{user_name}", ops.GetUserByUserName).Methods("GET")
	router_auth.HandleFunc("/users/add", ops.InsertUser).Methods("POST")
	router_auth.HandleFunc("/users/delete/{id}", ops.DeleteUser).Methods("DELETE")
	router_noauth.HandleFunc("/users/login", ops.LogMeIn).Methods("POST")

	// initialize the database

	//router.HandleFunc("/init", ops.InitializeDatabase).Methods("GET")


	negroni_wrapper := negroni.New(negroni.HandlerFunc(jwtMiddleware.HandlerWithNext), negroni.Wrap(router_auth))
	router_noauth.PathPrefix("/users").Handler(negroni_wrapper)

	negroni_instance := negroni.Classic()
	negroni_instance.UseHandler(router_noauth)
	negroni_instance.Run(":8080")
}
