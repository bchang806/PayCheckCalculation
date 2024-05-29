package main

import (
	"encoding/json"
	"fmt"

	"go_Paycalc/mylib"

	"net/http"

	_ "github.com/lib/pq"
)

type RequestBody struct {
	AnnualIncome       float64 `json:"annualIncome"`
	Dependent          float64 `json:"dependent"`
	AddIncome          float64 `json:"addIncome"`
	AddDeduct          float64 `json:"addDeduct"`
	ExtraWithhold      float64 `json:"extraWithhold"`
	FedMaritalStatus   string  `json:"fedMaritalStatus"`
	StateMaritalStatus string  `json:"stateMaritalStatus"`
	LocalMaritalStatus string  `json:"localMaritalStatus"`
	PayPeriod          string  `json:"payPeriod"`
}

type ResponseBody struct {
	FedTaxWithholding            float64 `json:"fedTaxWithholding"`
	SocialSecurityTaxWithholding float64 `json:"socialSecurityTaxWithholding"`
	MedicareTaxWithholding       float64 `json:"medicareTaxWithholding"`
	StateTaxWithholding          float64 `json:"stateTaxWithholding"`
	LocalTaxWithholding          float64 `json:"localTaxWithholding"`
	TakeHomePay                  float64 `json:"takeHomePay"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON request body
	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Perform the calculation
	fedTaxWithholding := mylib.CalculateFederalTax(reqBody.AnnualIncome, reqBody.Dependent, reqBody.AddIncome, reqBody.AddDeduct, reqBody.ExtraWithhold, reqBody.FedMaritalStatus, reqBody.PayPeriod)
	socialSecurityTaxWithholding := mylib.CalculateSocialTax(reqBody.AnnualIncome, reqBody.PayPeriod)
	medicareTaxWithholding := mylib.CalculateMedicareTax(reqBody.AnnualIncome, reqBody.PayPeriod)
	stateTaxWithholding := mylib.CalculateStateTax(reqBody.AnnualIncome, reqBody.StateMaritalStatus, reqBody.PayPeriod)
	localTaxWithholding := mylib.CalculateLocalTax(reqBody.AnnualIncome, reqBody.LocalMaritalStatus, reqBody.PayPeriod)
	takeHomePay := mylib.CalculateTakeHomePay(reqBody.AnnualIncome, fedTaxWithholding, socialSecurityTaxWithholding, medicareTaxWithholding, stateTaxWithholding, localTaxWithholding, reqBody.PayPeriod)

	// Prepare the response
	respBody := ResponseBody{
		FedTaxWithholding:            fedTaxWithholding,
		SocialSecurityTaxWithholding: socialSecurityTaxWithholding,
		MedicareTaxWithholding:       medicareTaxWithholding,
		StateTaxWithholding:          stateTaxWithholding,
		LocalTaxWithholding:          localTaxWithholding,
		TakeHomePay:                  takeHomePay,
	}

	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")
	// Encode the response as JSON and send it
	json.NewEncoder(w).Encode(respBody)
}

func main() {

	http.HandleFunc("/api/calculate", handler)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)

}
