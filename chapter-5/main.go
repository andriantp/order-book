package main

import "order-book/orderbook"

func main() {
	book := orderbook.NewOrderBook()
	book.AddOrder(&orderbook.Order{
		ID:    1,
		Side:  orderbook.Sell,
		Price: 100,
		Qty:   5,
	})
	book.Match(&orderbook.Order{
		ID:    2,
		Side:  orderbook.Buy,
		Price: 100,
		Qty:   7,
	})
	book.Print()
}
