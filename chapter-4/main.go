package main

import "order-book/orderbook"

func main() {
	book := orderbook.NewOrderBook()

	book.AddOrder(&orderbook.Order{ID: 1, Side: orderbook.Buy, Price: 100, Qty: 10})
	book.AddOrder(&orderbook.Order{ID: 2, Side: orderbook.Buy, Price: 100, Qty: 5})
	book.AddOrder(&orderbook.Order{ID: 3, Side: orderbook.Buy, Price: 99, Qty: 8})
	book.AddOrder(&orderbook.Order{ID: 4, Side: orderbook.Sell, Price: 101, Qty: 12})
	book.AddOrder(&orderbook.Order{ID: 5, Side: orderbook.Sell, Price: 102, Qty: 7})

	book.Print()
}
