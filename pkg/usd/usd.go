package usd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Visitor interface {
	VisitorForGetUSD(u *USD) string
}

// USD implements Currency interface
type USD struct {
}

// Accept implementation
func (u *USD) Accept(v Visitor) string {
	return v.VisitorForGetUSD(u)
}

// Currencies implemets getting rates interface
type Currencies struct {
}

// VisitorForGetUSD implements getting USD rate
func (c *Currencies) VisitorForGetUSD(u *USD) string {
	return u.GetUSD()
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