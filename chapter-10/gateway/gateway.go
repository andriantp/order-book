package gateway

import (
	"fmt"
	"order-book/engine"
	"order-book/orderbook"
)

type Gateway struct {
	manager *engine.Manager
}

func New(manager *engine.Manager) *Gateway {
	return &Gateway{
		manager: manager,
	}
}

func (g *Gateway) Route(symbol string, order *orderbook.Order) error {
	book, ok := g.manager.Get(symbol)
	if !ok {
		return fmt.Errorf("unknown symbol: %s", symbol)
	}
	book.AddOrder(order)
	return nil
}
