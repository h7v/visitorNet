package eur

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Visitor implements visitor interface
type Visitor interface {
	VisitorForGetEUR(e *EUR) string
}


// EUR implements Currency interface
type EUR struct {
}

// Currencies implemets getting rates interface
type Currencies struct {
}

// Accept for EUR
func (e *EUR) Accept(v Visitor) string {
	return v.VisitorForGetEUR(e)
}

// VisitorForGetEUR implements GetEUR
func (c *Currencies) VisitorForGetEUR(e *EUR) string {
	return e.GetEUR()
}

// GetEUR implementation
func (e *EUR) GetEUR() string {
	resp, noEr := http.Get("https://api.exchangeratesapi.io/latest?base=EUR")
	if noEr == nil {
		type forJSONParsing map[string]interface{}
		var result map[string]forJSONParsing
		json.NewDecoder(resp.Body).Decode(&result)
		var r = result["rates"]
		fmt.Printf("EUR = %.2f\n", r["RUB"])
	}
	return "We get EUR currency rate. "
}