package main

import (
	"fmt"
	"order-book/engine"
	"order-book/gateway"
	"order-book/orderbook"
	"order-book/router"
)

func main() {
	// Create engine manager.
	manager := engine.NewManager()

	// Register matching engines.
	manager.Register(newEngine("Engine-0"))
	manager.Register(newEngine("Engine-1"))
	manager.Register(newEngine("Engine-2"))

	// Create market router.
	marketRouter := router.New(manager)

	// Create gateway.
	gw := gateway.New(marketRouter)

	fmt.Println("Gateway is ready.")

	// Submit demo orders.
	submitOrders(gw)
}

func newEngine(symbol string) *orderbook.OrderBook {
	eventStore := orderbook.NewMemoryEventStore()
	snapshotStore := orderbook.NewMemorySnapshotStore()
	book := orderbook.NewOrderBook(eventStore, snapshotStore)
	fmt.Printf("[Engine %s] Recovering...\n", symbol)
	if err := book.Recover(); err != nil {
		panic(err)
	}
	fmt.Printf("[Engine %s] Ready.\n\n", symbol)
	return book
}
