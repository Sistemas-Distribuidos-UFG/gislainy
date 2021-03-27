package main

import (
	"github.com/mitchellh/mapstructure"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
)



// CreateAllRouter creates all endpoints
func CreateAllRouter() *mux.Router  {
	router := mux.NewRouter()

	router.
		HandleFunc("/calculate-readjust-salary", handleCalculateReajustSalary).
		Methods(http.MethodPost)

	router.
		HandleFunc("/check-legal-age", handleCheckLegalAge).
		Methods(http.MethodPost)
	router.
		HandleFunc("/calculate-result-final", handleCalculateFinalResult).
		Methods(http.MethodPost)

	router.
		HandleFunc("/calculate-ideal-weight", handleCalculateIdealWeight).
		Methods(http.MethodPost)

	return router;
}

func UnmarshalBody(r *http.Request, model *map[string]interface{}) error {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err;
	}
	
	if err := r.Body.Close(); err != nil {
		return err;
	}

	if err := json.Unmarshal(body, &model); err != nil {
		return err;
	}
	
	return nil

}

func handleCalculateReajustSalary (w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var request map[string]interface{}
	var employee Employee

	err := UnmarshalBody(r, &request);
	
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	mapstructure.Decode(request, &employee)
	
	result := CalculateReajustSalary(employee)

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
		
}

func handleCheckLegalAge (w http.ResponseWriter, r *http.Request) {

	
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var request map[string]interface{}
	var person PersonLegalAge

	err := UnmarshalBody(r, &request);
	
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	mapstructure.Decode(request, &person)
	
	result := CheckLegalAge(person)

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
		

}

func handleCalculateFinalResult (w http.ResponseWriter, r *http.Request) {


	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var request map[string]interface{}
	var sn StudentNotes

	err := UnmarshalBody(r, &request);
	
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	mapstructure.Decode(request, &sn)

	result := CalculateFinalResult(sn)

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	
	
}

func handleCalculateIdealWeight (w http.ResponseWriter, r *http.Request) {

	var person PersonWeight
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	
	if err := r.Body.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := json.Unmarshal(body, &person); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := CalculateIdealWeight(person)

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	
	
}



