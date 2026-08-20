package main

import (
	"fmt"

	"order-book/engine"
	"order-book/gateway"
	"order-book/orderbook"
)

func submitOrders(gw *gateway.Gateway) {
	orders := []struct {
		symbol string
		order  *orderbook.Order
	}{
		{
			symbol: "BTC-USDT",
			order: &orderbook.Order{
				ID:    1,
				Type:  orderbook.Limit,
				Side:  orderbook.Buy,
				Price: 100,
				Qty:   5,
			},
		},
		{
			symbol: "ETH-USDT",
			order: &orderbook.Order{
				ID:    2,
				Type:  orderbook.Limit,
				Side:  orderbook.Sell,
				Price: 200,
				Qty:   3,
			},
		},
		{
			symbol: "BTC-USDT",
			order: &orderbook.Order{
				ID:    3,
				Type:  orderbook.Limit,
				Side:  orderbook.Sell,
				Price: 101,
				Qty:   2,
			},
		},
		{
			symbol: "SOL-USDT",
			order: &orderbook.Order{
				ID:    4,
				Type:  orderbook.Limit,
				Side:  orderbook.Buy,
				Price: 50,
				Qty:   10,
			},
		},
		{
			symbol: "ETH-USDT",
			order: &orderbook.Order{
				ID:    5,
				Type:  orderbook.Limit,
				Side:  orderbook.Buy,
				Price: 199,
				Qty:   1,
			},
		},
		{
			symbol: "BTC-USDT",
			order: &orderbook.Order{
				ID:    6,
				Type:  orderbook.Limit,
				Side:  orderbook.Buy,
				Price: 99,
				Qty:   4,
			},
		},
	}

	for _, req := range orders {
		fmt.Printf("\nSubmitting order for %s...\n", req.symbol)

		if err := gw.Route(req.symbol, req.order); err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		fmt.Println("Order accepted.")
	}
}

func printReplication(manager *engine.Manager) {
	engines := []string{
		"Engine-0",
		"Engine-1",
		"Engine-2",
	}
	for _, id := range engines {
		engine, err := manager.Get(id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("\n=== %s ===\n", id)
		engine.Book.Print()
	}
}

func printReadModels(manager *engine.Manager) {
	engines := []string{
		"Engine-0",
		"Engine-1",
		"Engine-2",
	}

	for _, id := range engines {
		engine, err := manager.Get(id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		rm := engine.ReadModel

		fmt.Printf("\n=== %s READ MODEL ===\n", id)

		fmt.Printf("Best Bid  : %.2f\n", rm.BestBid)
		fmt.Printf("Best Ask  : %.2f\n", rm.BestAsk)
		fmt.Printf("Last Price: %.2f\n", rm.LastPrice)
		fmt.Printf("Last Qty  : %d\n", rm.LastQty)

		fmt.Println()

		fmt.Println("Bids")
		for price, qty := range rm.Bids {
			fmt.Printf("%.2f : %d\n", price, qty)
		}

		fmt.Println()

		fmt.Println("Asks")
		for price, qty := range rm.Asks {
			fmt.Printf("%.2f : %d\n", price, qty)
		}
	}
}
