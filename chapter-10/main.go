package main

import (
	"fmt"
	"order-book/engine"
	"order-book/gateway"
	"order-book/orderbook"
)

func main() {
	// Create engine manager.
	manager := engine.NewManager()

	// Register matching engines.
	manager.Register("BTC-USDT", newEngine("BTC-USDT"))
	manager.Register("ETH-USDT", newEngine("ETH-USDT"))
	manager.Register("SOL-USDT", newEngine("SOL-USDT"))

	// Create gateway.
	gw := gateway.New(manager)
	fmt.Println("Gateway is ready.")

	// Submit order.
	submitOrders(gw)
}

func newEngine(symbol string) *orderbook.OrderBook {
	eventStore := orderbook.NewMemoryEventStore()
	snapshotStore := orderbook.NewMemorySnapshotStore()

	book := orderbook.NewOrderBook(
		eventStore,
		snapshotStore,
	)

	fmt.Printf("[Engine %s] Recovering...\n", symbol)

	if err := book.Recover(); err != nil {
		panic(err)
	}

	fmt.Printf("[Engine %s] Ready.\n\n", symbol)
	return book
}
