package main

import (
	"fmt"
	"order-book/orderbook"
)

func main() {
	store := orderbook.NewMemoryEventStore()
	book := orderbook.NewOrderBook(store)
	book.AddOrder(&orderbook.Order{
		ID:    1,
		Type:  orderbook.Limit,
		Side:  orderbook.Sell,
		Price: 100,
		Qty:   5,
	})
	book.Match(&orderbook.Order{
		ID:    2,
		Type:  orderbook.Limit,
		Side:  orderbook.Buy,
		Price: 100,
		Qty:   3,
	})
	fmt.Println("=== ORIGINAL ORDER BOOK ===")
	book.Print()
	recovered := orderbook.NewOrderBook(store)
	if err := recovered.Replay(store.Load()); err != nil {
		panic(err)
	}
	fmt.Println("=== RECOVERED ORDER BOOK ===")
	recovered.Print()
}
