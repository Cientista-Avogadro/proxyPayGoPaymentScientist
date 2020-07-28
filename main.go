package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"proxypay/lib"
	"proxypay/model"
	"regexp"
	"strconv"
)

var (
	id     int
	err    error
	reg    = regexp.MustCompile(`/delete/([0-9]+)`)
	regpay = regexp.MustCompile(`/confirm/([0-9]+)`)
)

func main() {
	flag.Parse()

	http.HandleFunc("/create", create)
	http.HandleFunc("/delete/", delete)
	http.HandleFunc("/payments", payments)
	http.HandleFunc("/confirm/", confirm)
	http.HandleFunc("/dopayment", mock)

	go processPayment()
	go http.ListenAndServe(":8080", nil)
	fmt.Println("Proxy Pay reference is running")
	fmt.Scanln()
}

//handle request to created reference
func create(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	//check if method is POST
	if http.MethodPost != r.Method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var reference model.Reference

	err = json.NewDecoder(r.Body).Decode(&reference) //convert json to object

	log.Println(reference)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err = lib.GenerateID() //genereted ID

	log.Printf("Generated ID %d", id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = lib.GeneratedRef(reference, id) //create ID in proxypay

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]int{
		"reference": id,
	}

	json.NewEncoder(w).Encode(&response)
}

//handle request to delete reference
func delete(w http.ResponseWriter, r *http.Request) {

	//check if Method is DELETE
	if http.MethodDelete != r.Method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	matchs := reg.FindStringSubmatch(r.URL.Path)

	if len(matchs) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(matchs)

	id, _ := strconv.Atoi(matchs[1]) //convert string to int

	err := lib.DeleteRef(id) // call proxypay for delete reference ID

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

//handle request to mock payment
func mock(w http.ResponseWriter, r *http.Request) {

	//check if Method POST
	if http.MethodPost != r.Method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var payment model.MockPayment
	json.NewDecoder(r.Body).Decode(&payment)

	err := lib.MockPayment(payment)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//Handle request payment
func payments(w http.ResponseWriter, r *http.Request) {
	var i int
	w.Header().Add("Content-Type", "application/json")

	//check if Method GET
	if http.MethodGet != r.Method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	n := r.URL.Query().Get("n") //extrat query param
	if n != "" {
		i, err = strconv.Atoi(n) //convert to int
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	pays, err := lib.GetPayments(i)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&pays)

}

func confirm(w http.ResponseWriter, r *http.Request) {
	//check if Method PUT
	if http.MethodPut != r.Method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	matchs := regpay.FindStringSubmatch(r.URL.Path)

	if len(matchs) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(matchs)

	id, _ := strconv.Atoi(matchs[1])

	err := lib.ConfirmPayment(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
