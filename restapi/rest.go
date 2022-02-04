package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// model of a data struct

type Data struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Cost int    `json:"cost"`
}

//init datas var as a slice of data struct
var datas []Data

//functions

//get all data

func getallData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datas)
}

// get a data based on id passed in the http request

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params

	// loop through datas and find with id

	for _, item := range datas {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Data{})
}

//creating a new data bucket in slice datas
func createData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data Data
	_ = json.NewDecoder(r.Body).Decode(&data)
	data.ID = strconv.Itoa(rand.Intn(10000000))
	datas = append(datas, data)
	json.NewEncoder(w).Encode(data)

}

// updating data based on the given id
func updateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range datas {
		if item.ID == params["id"] {
			datas = append(datas[:index], datas[index+1:]...)
			var data Data
			_ = json.NewDecoder(r.Body).Decode(&data)
			data.ID = params["id"]
			datas = append(datas, data)
			json.NewEncoder(w).Encode(data)
			return
		}
	}
	json.NewEncoder(w).Encode(datas)
}

// deleting data based on the given id in http request

func deleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range datas {
		if item.ID == params["id"] {
			datas = append(datas[:index], datas[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(datas)

}
func main() {
	// initializing a router variable
	r := mux.NewRouter()

	//mock data- todo implement db

	datas = append(datas, Data{ID: "1", Name: "aman", Cost: 20})
	datas = append(datas, Data{ID: "2", Name: "rahul", Cost: 50})

	//route handlers that would establish endpoints of our api

	r.HandleFunc("/data", getallData).Methods("GET")
	r.HandleFunc("/data/{id}", getData).Methods("GET")
	r.HandleFunc("/data", createData).Methods("POST")
	r.HandleFunc("/data/{id}", updateData).Methods("PUT")
	r.HandleFunc("/data/{id}", deleteData).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}