package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Visitor implements visitor interface
type Visitor interface {
	VisitorForGetUSD(u *USD) string
	VisitorForGetEUR(e *EUR) string
}

//Currency implements interface that Visitor must get
type Currency interface {
	Accept(v Visitor) string
}

// Currencies implemets getting rates interface
type Currencies struct {
}

// VisitorForGetUSD implements getting USD rate
func (c *Currencies) VisitorForGetUSD(u *USD) string {
	return u.GetUSD()
}

// VisitorForGetEUR implements GetEUR
func (c *Currencies) VisitorForGetEUR(e *EUR) string {
	return e.GetEUR()
}

// Review implements currencies collection for getting rates
type Review struct {
	currencies []Currency
}

// Add adding currency to collection
func (r *Review) Add(c Currency) {
	r.currencies = append(r.currencies, c)
}

// Accept implements visiter interface to get all currencies
func (r *Review) Accept(v Visitor) string {
	var result string
	for _, p := range r.currencies {
		result += p.Accept(v)

	}
	return result
}

// USD implements Currency interface
type USD struct {
}

// Accept implementation
func (u *USD) Accept(v Visitor) string {
	return v.VisitorForGetUSD(u)
}

// GetUSD implementation
func (u *USD) GetUSD() string {
	resp, noEr := http.Get("https://api.exchangeratesapi.io/latest?base=USD")
	if noEr == nil {
		type forJSONParsing map[string]interface{}
		var result map[string]forJSONParsing
		json.NewDecoder(resp.Body).Decode(&result)
		var r = result["rates"]
		fmt.Printf("USD = %.2f\n", r["RUB"])
	}
	return "We get USD curency rate. "
}

// EUR implements Currency interface
type EUR struct {
}

// Accept for EUR
func (e *EUR) Accept(v Visitor) string {
	return v.VisitorForGetEUR(e)
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

func main() {
	ratesReview := new(Review)
	ratesReview.Add(&USD{})
	ratesReview.Add(&EUR{})
	result := ratesReview.Accept(&Currencies{})
	fmt.Println(result)
}
