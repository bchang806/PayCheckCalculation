package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go_Paycalc/mylib"

	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "080691"
	dbname   = "HR101"
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

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open connection to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	defer db.Close()

	// Ping the database to verify connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %q", err)
	}

	fmt.Println("Successfully connected to the database!")

	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/api/calculate", func(c *gin.Context) {
		var reqBody RequestBody

		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fedTaxWithholding := mylib.CalculateFederalTax(reqBody.AnnualIncome, reqBody.Dependent, reqBody.AddIncome, reqBody.AddDeduct, reqBody.ExtraWithhold, reqBody.FedMaritalStatus, reqBody.PayPeriod, db)
		socialSecurityTaxWithholding := mylib.CalculateSocialTax(reqBody.AnnualIncome, reqBody.PayPeriod)
		medicareTaxWithholding := mylib.CalculateMedicareTax(reqBody.AnnualIncome, reqBody.PayPeriod)
		stateTaxWithholding := mylib.CalculateStateTax(reqBody.AnnualIncome, reqBody.StateMaritalStatus, reqBody.PayPeriod)
		localTaxWithholding := mylib.CalculateLocalTax(reqBody.AnnualIncome, reqBody.LocalMaritalStatus, reqBody.PayPeriod)
		takeHomePay := mylib.CalculateTakeHomePay(reqBody.AnnualIncome, fedTaxWithholding, socialSecurityTaxWithholding, medicareTaxWithholding, stateTaxWithholding, localTaxWithholding, reqBody.PayPeriod)

		responseBody := ResponseBody{
			FedTaxWithholding:            fedTaxWithholding,
			SocialSecurityTaxWithholding: socialSecurityTaxWithholding,
			MedicareTaxWithholding:       medicareTaxWithholding,
			StateTaxWithholding:          stateTaxWithholding,
			LocalTaxWithholding:          localTaxWithholding,
			TakeHomePay:                  takeHomePay,
		}
		c.JSON(http.StatusOK, responseBody)
	})

	router.Run(":8080")

}
