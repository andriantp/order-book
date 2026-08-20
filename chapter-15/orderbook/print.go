package orderbook

import "fmt"

func (ob *OrderBook) Print() {
	
	fmt.Println("ASKS")

	for _, level := range ob.Asks {
		fmt.Printf("%.0f :", level.Price)

		for _, order := range level.Orders {
			fmt.Printf(" [#%d Qty:%d]", order.ID, order.Qty)
		}

		fmt.Println()
	}

	fmt.Println()

	fmt.Println("BIDS")

	for _, level := range ob.Bids {
		fmt.Printf("%.0f :", level.Price)

		for _, order := range level.Orders {
			fmt.Printf(" [#%d Qty:%d]", order.ID, order.Qty)
		}

		fmt.Println()
	}
}
