package main

import (
	"fmt"
	"order-book/orderbook"
)

func main() {
	eventStore := orderbook.NewMemoryEventStore()
	snapshotStore := orderbook.NewMemorySnapshotStore()
	book := orderbook.NewOrderBook(
		eventStore,
		snapshotStore,
	)
	seedBook(book)
	createSnapshot(book)
	processNewEvents(book)
	fmt.Println("=== ORIGINAL ORDER BOOK ===")
	book.Print()
	recovered := recoverBook(
		eventStore,
		snapshotStore,
	)
	fmt.Println("=== RECOVERED ORDER BOOK ===")
	recovered.Print()
}

func seedBook(book *orderbook.OrderBook) {
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
}

func createSnapshot(book *orderbook.OrderBook) {
	if err := book.SaveSnapshot(); err != nil {
		panic(err)
	}
}

func processNewEvents(book *orderbook.OrderBook) {
	book.AddOrder(&orderbook.Order{
		ID:    3,
		Type:  orderbook.Limit,
		Side:  orderbook.Sell,
		Price: 101,
		Qty:   2,
	})
}

func recoverBook(
	eventStore orderbook.EventStore,
	snapshotStore orderbook.SnapshotStore,
) *orderbook.OrderBook {
	book := orderbook.NewOrderBook(
		eventStore,
		snapshotStore,
	)
	if err := book.Recover(); err != nil {
		panic(err)
	}
	return book
}
