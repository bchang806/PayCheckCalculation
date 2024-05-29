package mylib

import (
	"database/sql"
	"fmt"
	"log"
	"math"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "080691"
	dbname   = "HR101"
)

type Fed_tax_table struct {
	pay_freq      string
	marital       string
	earning       float64
	amount        float64
	percentage    float64
	adjust_income float64
}

func CalculateTakeHomePay(annuaIncome float64, fedTax float64, socialTax float64, medicareTax float64, stateTax float64, locTax float64, payPeriod string) float64 {
	payFreq := CalcPayFreq(payPeriod)
	totalPayPerCheck := annuaIncome / float64(payFreq)
	takeHomePayPerCheck := totalPayPerCheck - fedTax - socialTax - medicareTax - stateTax - locTax
	takeHomePayPerCheck = math.Floor(takeHomePayPerCheck*100) / 100
	fmt.Printf("\nThe take home pay is: $%.2f", takeHomePayPerCheck)
	return takeHomePayPerCheck
}
func calculateAdjustedIncome(totalPayPerCheck float64, addIncome float64, addDeduct float64, adjust_income float64) float64 {
	adjustedIncome := totalPayPerCheck + addIncome - addDeduct - adjust_income

	return adjustedIncome

}

func CalculateFederalTax(annuaIncome float64, dependent float64, addIncome float64, addDeduct float64, extraWithhold float64, fedMaritalStatus string, payPeriod string) float64 {
	var fedTax float64 = 0

	payFreq := CalcPayFreq(payPeriod)
	fed_tax_table := getFedTaxTable(fedMaritalStatus)

	totalAdjustedIncome := calculateAdjustedIncome(annuaIncome, addIncome, addDeduct, fed_tax_table[1].adjust_income)
	taxable := totalAdjustedIncome
	var old_earn float64 = 0

	for _, row := range fed_tax_table {
		if taxable >= old_earn {
			if taxable < row.earning {
				fedTax = row.amount + (taxable-old_earn)*row.percentage
				break
			}
		}
		old_earn = row.earning
	}

	fedTax = fedTax / float64(payFreq)

	fedTax = fedTax - (dependent / float64(payFreq)) //dependent credit Step 3 on W4
	if fedTax < 0 {
		fedTax = 0
	}

	fedTax = fedTax + extraWithhold //extra withholding per pay period Step 4c on W4

	fedTax = math.Floor(fedTax*100) / 100
	fmt.Printf("\nThe fed tax is: $%.2f", fedTax)

	return fedTax

}

func CalculateStateTax(annuaIncome float64, stateMaritalStatus string, payPeriod string) float64 {
	var stateTax float64 = 0
	payFreq := CalcPayFreq(payPeriod)
	totalPayPerCheck := annuaIncome / float64(payFreq)
	stateTax = totalPayPerCheck * 0 // to be changed later
	fmt.Printf("\nThe state tax is: $%.2f", stateTax)
	return stateTax
}

func CalculateLocalTax(annuaIncome float64, localMaritalStatus string, payPeriod string) float64 {
	var locTax float64 = 0
	payFreq := CalcPayFreq(payPeriod)
	totalPayPerCheck := annuaIncome / float64(payFreq)
	locTax = totalPayPerCheck * 0 // to be changed later
	fmt.Printf("\nThe local tax is: $%.2f", locTax)
	return locTax
}

func CalcPayFreq(payPeriod string) int {
	var payFreq int
	switch payPeriod {
	case "M":
		payFreq = 12
	case "B":
		payFreq = 26
	case "S":
		payFreq = 24
	case "W":
		payFreq = 52
	}

	return payFreq
}
func CalculateSocialTax(annuaIncome float64, payPeriod string) float64 {
	var socialTax float64 = 0
	var socialTaxRate float64 = 0.062
	payFreq := CalcPayFreq(payPeriod)
	totalPayPerCheck := annuaIncome / float64(payFreq)

	socialTax = totalPayPerCheck * socialTaxRate
	socialTax = math.Floor(socialTax*100) / 100
	fmt.Printf("\nThe Social Security Tax is: $%.2f", socialTax)
	return socialTax
}

func CalculateMedicareTax(annuaIncome float64, payPeriod string) float64 {
	var medicareTax float64 = 0
	var medicareTaxRate float64 = 0.0145
	payFreq := CalcPayFreq(payPeriod)
	totalPayPerCheck := annuaIncome / float64(payFreq)

	medicareTax = totalPayPerCheck * medicareTaxRate
	medicareTax = math.Floor(medicareTax*100) / 100
	fmt.Printf("\nThe Medicare Tax is: $%.2f", medicareTax)
	return medicareTax
}

func getFedTaxTable(fed_marital string) []Fed_tax_table {
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

	query := "SELECT * FROM fed_tax_table where marital = $1 order by earning"
	rows, err := db.Query(query, fed_marital)
	if err != nil {
		log.Fatalf("Error querying database: %q", err)
	}
	defer rows.Close()

	var fed_tax_table []Fed_tax_table

	// Iterate over rows and print data
	for rows.Next() {
		var row Fed_tax_table
		err := rows.Scan(&row.pay_freq, &row.marital, &row.earning, &row.amount, &row.percentage, &row.adjust_income)
		if err != nil {
			log.Fatalf("Error scanning row: %q", err)
		}
		fed_tax_table = append(fed_tax_table, row)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		log.Fatalf("Error during iteration: %q", err)
	}

	return fed_tax_table
}
