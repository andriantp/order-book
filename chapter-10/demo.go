package main

import (
	"fmt"

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
			symbol: "SOL-USDT",
			order: &orderbook.Order{
				ID:    3,
				Type:  orderbook.Limit,
				Side:  orderbook.Buy,
				Price: 50,
				Qty:   10,
			},
		},
		{
			symbol: "HYPE-USDT",
			order: &orderbook.Order{
				ID:    4,
				Type:  orderbook.Limit,
				Side:  orderbook.Buy,
				Price: 30,
				Qty:   2,
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
