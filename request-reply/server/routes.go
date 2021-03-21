package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
)



// CreateAllRouter creates all endpoints
func CreateAllRouter() *mux.Router  {
	router := mux.NewRouter()

	router.HandleFunc("/calculate-readjust-salary", handleCalculateReajustSalary)
	router.HandleFunc("/check-legal-age", handleCheckLegalAge)
	router.HandleFunc("/calculate-result-final", handleCalculateFinalResult)
	router.HandleFunc("/calculate-ideal-weight", handleCalculateIdealWeight)

	return router;
}

func handleCalculateReajustSalary (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var employee Employee
		
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		
		if err := r.Body.Close(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		if err := json.Unmarshal(body, &employee); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}


		result := CalculateReajustSalary(employee)

		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		
	default:
		fmt.Fprintf(w, "Method %s not supported", r.Method)
	}
}

func handleCheckLegalAge (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var person PersonLegalAge
		
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


		result := CheckLegalAge(person)

		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		
	default:
		fmt.Fprintf(w, "Method %s not supported", r.Method)
	}
}

func handleCalculateFinalResult (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var sn StudentNotes
		
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		
		if err := r.Body.Close(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		if err := json.Unmarshal(body, &sn); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}


		result := CalculateFinalResult(sn)

		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		
	default:
		fmt.Fprintf(w, "Method %s not supported", r.Method)
	}
}

func handleCalculateIdealWeight (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
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
		
	default:
		fmt.Fprintf(w, "Method %s not supported", r.Method)
	}
}



