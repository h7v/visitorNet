package visitor

// USD implements Currency interface
type USD struct {
}

// EUR implements Currency interface
type EUR struct {
}

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
