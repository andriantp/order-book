package router

import (
	"fmt"
	"order-book/engine"
	"order-book/orderbook"
)

type MarketRouter struct {
	manager *engine.Manager
}

func New(manager *engine.Manager) *MarketRouter {
	return &MarketRouter{
		manager: manager,
	}
}

func (r *MarketRouter) Route(market string) (*orderbook.OrderBook, error) {
	books := r.manager.Books()
	if len(books) == 0 {
		return nil, fmt.Errorf("no matching engine available")
	}
	index := int(hashMarket(market)) % len(books)
	fmt.Printf("[Router] %s -> Engine-%d\n", market, index)
	return books[index], nil
}
