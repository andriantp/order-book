package main

import (
	"fmt"
	"order-book/engine"
	"order-book/gateway"
	"order-book/orderbook"
	"order-book/ownership"
	"order-book/replication"
	"order-book/router"
)

func main() {
	// Create engine manager.
	manager := engine.NewManager()

	// Register matching engines.
	manager.Register("Engine-0", newEngine("Engine-0"))
	manager.Register("Engine-1", newEngine("Engine-1"))
	manager.Register("Engine-2", newEngine("Engine-2"))

	// Create ownership manager.
	ownershipManager := ownership.NewManager()

	// Assign market ownership.
	ownershipManager.Assign("BTC-USDT", "Engine-0")
	ownershipManager.Assign("ETH-USDT", "Engine-1")
	ownershipManager.Assign("SOL-USDT", "Engine-2")

	// Create replication manager.
	replicationManager := replication.NewManager()

	// Assign replicas.
	replicationManager.Assign("BTC-USDT", "Engine-1")
	replicationManager.Assign("BTC-USDT", "Engine-2")

	replicationManager.Assign("ETH-USDT", "Engine-0")
	replicationManager.Assign("ETH-USDT", "Engine-2")

	replicationManager.Assign("SOL-USDT", "Engine-0")
	replicationManager.Assign("SOL-USDT", "Engine-1")

	// Create market router.
	marketRouter := router.New(manager, ownershipManager)

	// Create gateway.
	gw := gateway.New(marketRouter, manager, replicationManager)

	fmt.Println("Gateway is ready.")

	// Submit demo orders.
	submitOrders(gw)

	fmt.Println()
	
	// Print replicated order books.
	printReplication(manager)
}

func newEngine(name string) *engine.Engine {
	eventStore := orderbook.NewMemoryEventStore()
	snapshotStore := orderbook.NewMemorySnapshotStore()

	book := orderbook.NewOrderBook(eventStore, snapshotStore)

	fmt.Printf("[Engine %s] Recovering...\n", name)
	book.Replay(eventStore.Load())
	fmt.Printf("[Engine %s] Ready.\n\n", name)

	return &engine.Engine{
		Book:       book,
		EventStore: eventStore,
	}
}
