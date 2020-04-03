package main

import (
	"fmt"

	"visitorTest2/pkg/eur"
	"visitorTest2/pkg/usd"
	"visitorTest2/pkg/visitor"
)

func main() {
	ratesReview := new(visitor.Review)
	ratesReview.Add(&usd.USD{})
	ratesReview.Add(&eur.EUR{})
	result := ratesReview.Accept(&visitor.Currencies{})
	fmt.Println(result)
}
