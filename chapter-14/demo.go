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
