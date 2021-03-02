package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Feature Struct (Model)
var DB *sql.DB

const (
	host     = "ec2-3-231-194-96.compute-1.amazonaws.com"
	port     = 5432
	user     = "ntcfjubqinaxjn"
	password = "602244e75e57f55bba480ffdbb5dfb3238ccafe1074b69b88e463f41e2c5c69d"
	dbname   = "d8mpas9bratrsh"
)

type Feature struct {
	ID          string `json:"id" `
	Name        string `json:"name" `
	Description string `json:"description" `
	Votes       int    `json:"votes" `
	Companyid   string `json:"companyid"`
}

type Company struct {
	ID      string `json:id`
	Company string `json:company`
}

//init feature mock data
var features []Feature
var companies []Company

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func getFeatures(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("pnt is a nil pointer in func: %v\n", DB == nil)
	rows, err := DB.Query("SELECT * FROM features")
	fmt.Println(rows, err)
	fmt.Println("ERROR 1", err)
	if err != nil {
		fmt.Println("ERROR 1")
		// handle this error better than this
		panic(err)
	}
	for rows.Next() {
		var id int
		var upvotes int
		var feature string
		var description string
		var companyid string
		err = rows.Scan(&id, &feature, &description, &upvotes, &companyid)
		if err != nil {
			fmt.Println("ERROR 2")
			// handle this error
			panic(err)
		}
		fmt.Println(id, feature, description, upvotes, companyid)
		features = append(features, Feature{strconv.Itoa(id), feature, description, upvotes, companyid})

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("ERROR 3")
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(features)
}
func getCompanies(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("pnt is a nil pointer in func: %v\n", DB == nil)
	rows, err := DB.Query("SELECT * FROM companies")
	fmt.Println(rows, err)
	fmt.Println("ERROR 1", err)
	if err != nil {
		fmt.Println("ERROR 1")
		// handle this error better than this
		panic(err)
	}
	for rows.Next() {
		var id int
		var company string

		err = rows.Scan(&id, &company)
		if err != nil {
			fmt.Println("ERROR 2")
			// handle this error
			panic(err)
		}
		fmt.Println(id, company)
		companies = append(companies, Company{strconv.Itoa(id), company})

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("ERROR 3")
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(companies)
}
func getFeature(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//
	fmt.Printf("pnt is a nil pointer in func: %v\n", DB == nil)
	rows, err := DB.Query("SELECT * FROM features WHERE id = " + params["id"])
	if err != nil {
		fmt.Println("ERROR 1")
		// handle this error better than this
		panic(err)
	}
	for rows.Next() {
		var id int
		var upvotes int
		var feature string
		var description string
		var companyid string
		err = rows.Scan(&id, &feature, &description, &upvotes, &companyid)
		if err != nil {
			fmt.Println("ERROR 2")
			// handle this error
			panic(err)
		}
		fmt.Println(id, feature, description, upvotes, companyid)
		features = append(features, Feature{strconv.Itoa(id), feature, description, upvotes, companyid})

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("ERROR 3")
		panic(err)
	}
	//

	for _, item := range features {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	json.NewEncoder(w).Encode(&Feature{})
}
func getCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//
	fmt.Printf("pnt is a nil pointer in func: %v\n", DB == nil)
	rows, err := DB.Query("SELECT * FROM companies WHERE id = " + params["id"])
	if err != nil {
		fmt.Println("ERROR 1")
		// handle this error better than this
		panic(err)
	}
	for rows.Next() {
		var id int
		var company string
		err = rows.Scan(&id, &company)
		if err != nil {
			fmt.Println("ERROR 2")
			// handle this error
			panic(err)
		}
		fmt.Println(id, company)
		companies = append(companies, Company{strconv.Itoa(id), company})

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("ERROR 3")
		panic(err)
	}
	//

	for _, item := range companies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	json.NewEncoder(w).Encode(&Company{})
}
func createFeature(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var feature Feature
	sqlQuery := `INSERT INTO features (id, feature, description, upvotes) VALUES($1, $2, $3, $4)`
	_, err := DB.Exec(sqlQuery, strconv.Itoa(rand.Intn(10000000)), "testFeature", "testing a feature", 100)
	if err != nil {
		fmt.Println("ERROR 1")
		// handle this error better than this
		panic(err)
	}
}
func createCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var company Company
	_ = json.NewDecoder(r.Body).Decode(&company)
	company.ID = strconv.Itoa(rand.Intn(10000000))
	// features = append(features, feature)
	sqlQuery := `INSERT INTO companies (id, company) VALUES($1, $2)`
	_, err := DB.Exec(sqlQuery, company.ID, company.Company)
	if err != nil {
		fmt.Println("ERROR 1")
		// handle this error better than this
		panic(err)
	}
}
func updateFeature(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	sqlQuery := `UPDATE features SET feature=$1, description=$2, upvotes=$3 WHERE id = $4`
	_, err := DB.Exec(sqlQuery, "testFeature2222", "testing a feature", 9000, params["id"])
	if err != nil {
		fmt.Println("ERROR 1")
		// handle this error better than this
		panic(err)
	}
}
func deleteFeature(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	sqlQuery := `DELETE FROM features WHERE id = $1`
	_, err := DB.Exec(sqlQuery, params["id"])
	if err != nil {
		fmt.Println("ERROR 1")
		// handle this error better than this
		panic(err)
	}
	json.NewEncoder(w).Encode(features)

}
func main() {

	db, err := sql.Open("postgres", "postgres://ntcfjubqinaxjn:602244e75e57f55bba480ffdbb5dfb3238ccafe1074b69b88e463f41e2c5c69d@ec2-3-231-194-96.compute-1.amazonaws.com:5432/d8mpas9bratrsh")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	DB = db

	fmt.Println("Successfully connected!")
	fmt.Printf("pnt is a nil pointer in main: %v\n", db == nil)

	//init router
	r := mux.NewRouter()
	// router handlers
	r.HandleFunc("/api/features", getFeatures).Methods("GET")
	r.HandleFunc("/api/companies", getCompanies).Methods("GET")
	r.HandleFunc("/api/company/{id}", getCompany).Methods("GET")
	r.HandleFunc("/api/feature/{id}", getFeature).Methods("GET")
	r.HandleFunc("/api/features", createFeature).Methods("POST")
	r.HandleFunc("/api/companies", createCompany).Methods("POST")
	r.HandleFunc("/api/feature/{id}", updateFeature).Methods("PUT")
	r.HandleFunc("/api/feature/{id}", deleteFeature).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
