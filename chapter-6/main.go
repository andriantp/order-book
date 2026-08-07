package main

import (
	"fmt"
	"order-book/orderbook"
)

func main() {
	book := seedBook()

	fmt.Println("=== INITIAL ORDER BOOK ===")
	book.Print()
	fmt.Println()

	limitOrder()
	fmt.Println()
	marketOrder()
}

func seedBook() *orderbook.OrderBook {
	book := orderbook.NewOrderBook()

	book.AddOrder(&orderbook.Order{
		ID:    1,
		Type:  orderbook.Limit,
		Side:  orderbook.Sell,
		Price: 100,
		Qty:   5,
	})

	return book
}

func limitOrder() {
	fmt.Println("=== LIMIT ORDER ===")

	book := seedBook()
	book.Match(&orderbook.Order{
		ID:    2,
		Type:  orderbook.Limit,
		Side:  orderbook.Buy,
		Price: 99,
		Qty:   5,
	})

	book.Print()
}

func marketOrder() {
	fmt.Println("=== MARKET ORDER ===")

	book := seedBook()
	book.Match(&orderbook.Order{
		ID:   3,
		Type: orderbook.Market,
		Side: orderbook.Buy,
		Qty:  7,
	})

	book.Print()
}
