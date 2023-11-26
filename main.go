package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	persondetails "person-crud-api/model"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	ch int
	p  []persondetails.Person
)

func getPersonDetails(w http.ResponseWriter, _ *http.Request) {
	// API Call: http://localhost:8080/personDetails
	w.Header().Set("content-type", "aplication/json")
	json.NewEncoder(w).Encode(p)
}

func getPersonDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range p {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(p)
		}
	}
}

func createPersonDetail(w http.ResponseWriter, r *http.Request) {
	// API Call: http://localhost:8080/personDetails
	w.Header().Set("content-type", "aplication/json")
	var person persondetails.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.Id = strconv.Itoa(rand.Intn(1000))
	p = append(p, person)
	json.NewEncoder(w).Encode(p)
}

func updatePersonDetail(w http.ResponseWriter, r *http.Request) {

}

func deletePersonDetail(w http.ResponseWriter, r *http.Request) {
	// API Call: http://localhost:8080/personDetails/1
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range p {
		if item.Id == params["id"] {
			p = append(p[:index], p[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(p)
}

func main() {

	x := []int{}
	y := 10
	x = append(x, y)

	for {
		fmt.Println("<==============================================>")
		fmt.Println("1. Insert Person Details")
		fmt.Println("2. Delete Person Details")
		fmt.Println("3. Update Person Details")
		fmt.Println("4. Display Person Details")
		fmt.Println("5. Exit")
		fmt.Print(">>>Choose an Option: ")
		fmt.Scan(&ch)
		fmt.Println()
		switch ch {
		case 1:
			setPersonDetails := persondetails.Person{}
			setPersonDetails.InsertPersonDetails()
			p = append(p, setPersonDetails)
			fmt.Println("Person Details is Inserted!")
		case 2:
			var pos int
			if len(p) <= 0 {
				fmt.Println("No record is available")
				continue
			}
			fmt.Print("Enter the position you want to delete: ")
			fmt.Scan(&pos)
			p = append(p[:pos-1], p[pos:]...)
			fmt.Println("Person details has deleted")
		case 3:
			var i int
			if len(p) <= 0 {
				fmt.Println("No record is available")
				continue
			}
			fmt.Print("Enter the Position for updation: ")
			fmt.Scan(&i)
			p[i].UpdatePersonDetails()
		case 4:
			if len(p) <= 0 {
				fmt.Println("No record is available")
				continue
			}
			for i := range p {
				p[i].DisplayPersonDetails()
			}
		case 5:
			// os.Exit(0)
			goto label
		default:
			fmt.Println("Please Enter correct option!!")
		}
	}

label:
	r := mux.NewRouter()

	// structure: r.HandleFunc("/endPoint", funcName).Methods("GET"/"PUT"/"POST"/"DELETE")
	r.HandleFunc("/personDetails", getPersonDetails).Methods("GET")
	r.HandleFunc("/personDetails/{id}", getPersonDetail).Methods("GET")
	r.HandleFunc("/personDetails", createPersonDetail).Methods("POST")
	r.HandleFunc("/personDetails/{id}", updatePersonDetail).Methods("PUT")
	r.HandleFunc("/personDetails/{id}", deletePersonDetail).Methods("DELETE")

	fmt.Println("Starting server at port: 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
