# PayCheckCalculation

This API is used for calculating the take home pay for employees on New W4. It is written in Go and connect to PostgreSQL Database.
Currently it would not be able to calculate the state and local tax. They are there for future development.
The federal marital status can only be used "S" -Single, "M' - Married, "HH" -Head of Household, "2M" - Married and the checkbox is checked on 2c on W4 form, "2S" - Single and the checkbox is checked on 2c on W4 form, and "2H" - Head of household and the checkbox is checked on 2c on W4 form.

Below are an example of requeset body and response

{
  "annualIncome": 75000.00,
  "dependent": 0,
  "addIncome": 0,
  "addDeduct": 0,
  "extraWithhold": 0,
  "fedMaritalStatus": "S",
  "stateMaritalStatus": "S",
  "localMaritalStatus": "S",
  "payPeriod": "S"
}

{
    "fedTaxWithholding": 347.54,
    "socialSecurityTaxWithholding": 193.75,
    "medicareTaxWithholding": 45.31,
    "stateTaxWithholding": 0,
    "localTaxWithholding": 0,
    "takeHomePay": 2538.4
}
