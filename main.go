package main

import (
	"fmt"
	"github.com/labstack/echo"
	"encoding/json"
    "log"
    "net/http"
	"github.com/gorilla/mux"
)


func main(){

// init files index.html
router := mux.NewRouter()
a := echo.New()
a.Static("/", "public")


// request from de page
router.HandleFunc("/all-students", GetPeopleEndpoint).Methods("GET")

// Double listener
go func ()  {
			a.Logger.Fatal(a.Start(":8080"),)
}()
log.Fatal(http.ListenAndServe(":8081", router))

}

// Funciones from here

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request){
	es, err := allstudents()

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Println("the query student was done succesfull")
	fmt.Println(es)
	json.NewEncoder(w).Encode(es)

}
/*
func for_creating() {
	e := student{
		name:   "Carlos",
		Age:    30,
		Active: true,
	}
	err := creation(e)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("student was created succesfull")
}
*/
