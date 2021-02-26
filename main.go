package main
import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

// Feature Struct (Model)
var DB *sql.DB

const (
	host = "ec2-3-231-194-96.compute-1.amazonaws.com"
	port = 5432
	user = "ntcfjubqinaxjn"
	password = "602244e75e57f55bba480ffdbb5dfb3238ccafe1074b69b88e463f41e2c5c69d"
	dbname = "d8mpas9bratrsh"
)
type Feature struct {
	ID string `json:"id" `
	Name string `json:"name" `
	Description string `json:"description" `
	Votes int `json:"votes" `
}

//init feature mock data
var features []Feature
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request) 
  }
func getFeatures(w http.ResponseWriter, r * http.Request) {

    fmt.Printf("pnt is a nil pointer in func: %v\n", DB == nil) 
	rows, err := DB.Query("SELECT * FROM features")
	fmt.Println(rows, err)
	fmt.Println("ERROR 1",err)
	if err != nil {
		fmt.Println("ERROR 1")
		// handle this error better than this
		panic(err)
	  }
	  fmt.Println("ROWS")
	  fmt.Println(rows)

	  for rows.Next() {
		var id int
		var upvotes int
		var feature string
		var description string
		err = rows.Scan(&id, &feature, &description, &upvotes)
		if err != nil {
			fmt.Println("ERROR 2")
		  // handle this error
		  panic(err)
		}
		fmt.Println(id,feature, description,upvotes)
		features = append(features, Feature{strconv.Itoa(id),feature,description,upvotes})

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
func getFeature(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range features {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
json.NewEncoder(w).Encode(&Feature{})
}
func createFeature(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var feature Feature
	_ = json.NewDecoder(r.Body).Decode(&feature)
	feature.ID  = strconv.Itoa(rand.Intn(10000000))
	features = append(features, feature)

	//
	//
	
//

json.NewEncoder(w).Encode(feature)


}
func updateFeature(w http.ResponseWriter, r * http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range features {
		if item.ID == params["id"]{
			features = append(features[:index], features[index+1:]...)
			var feature Feature
	_ = json.NewDecoder(r.Body).Decode(&feature)
	feature.ID  = item.ID
	features = append(features, feature)

json.NewEncoder(w).Encode(feature)
return
		}
		
		
	}
	json.NewEncoder(w).Encode(features)

}
func deleteFeature(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range features {
		if item.ID == params["id"]{
			features = append(features[:index], features[index+1:]...)
			break
		}
		
		
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
  

//   sqlStatement := `INSERT INTO features VALUES (4,'feature 2', 'new page for signups', 100 );`
//   _, err = db.Exec(sqlStatement)
//   if err != nil {
// 	panic(err)
//   }
	
	
	//init router
	r := mux.NewRouter()
	// features = append(features, Feature{"1","more color","give the website more color on the home page",40})
	// features = append(features, Feature{"2","less color","give the website less color on the home page",69})
	// features = append(features, Feature{"3","same color","give the website same color on the home page",2})

	// router handlers
	r.HandleFunc("/api/features",getFeatures).Methods("GET")
	r.HandleFunc("/api/feature/{id}",getFeature).Methods("GET")
	r.HandleFunc("/api/features",createFeature).Methods("POST")
	r.HandleFunc("/api/feature/{id}",updateFeature).Methods("PUT")
	r.HandleFunc("/api/feature/{id}",deleteFeature).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",r))

}
