package main

import (
	"fmt"
)

type Order struct {
	ID       string
	Side     string
	Price    int64
	Quantity int64
}

type PriceLevel struct {
	Price  int64
	Orders []Order
}

func main() {
	fmt.Println("Order Book Data Structures")

	/*order := Order{
		ID:       "ORD-001",
		Side:     "BUY",
		Price:    100,
		Quantity: 10,
	}

	if js, err := json.MarshalIndent(order, "", "  "); err == nil {
		fmt.Println(string(js))
	}*/

	/*level := PriceLevel{
		Price: 100,
		Orders: []Order{
			{
				ID:       "ORD-001",
				Side:     "BUY",
				Price:    100,
				Quantity: 10,
			},
			{
				ID:       "ORD-002",
				Side:     "BUY",
				Price:    100,
				Quantity: 20,
			},
		},
	}

	if js, err := json.MarshalIndent(level, "", "  "); err == nil {
		fmt.Println(string(js))
	}*/
}
