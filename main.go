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

func getFeatures(w http.ResponseWriter, r * http.Request) {
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
	fmt.Println("Successfully connected!")
	fmt.Println(db)

	//init router
	r := mux.NewRouter()
	features = append(features, Feature{"1","more color","give the website more color on the home page",40})
	features = append(features, Feature{"2","less color","give the website less color on the home page",69})
	features = append(features, Feature{"3","same color","give the website same color on the home page",2})

	// router handlers
	r.HandleFunc("/api/features",getFeatures).Methods("GET")
	r.HandleFunc("/api/feature/{id}",getFeature).Methods("GET")
	r.HandleFunc("/api/features",createFeature).Methods("POST")
	r.HandleFunc("/api/feature/{id}",updateFeature).Methods("PUT")
	r.HandleFunc("/api/feature/{id}",deleteFeature).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",r))

}
