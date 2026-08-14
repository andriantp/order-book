package gateway

import (
	"order-book/orderbook"
	"order-book/router"
)

type Gateway struct {
	router *router.MarketRouter
}

func New(router *router.MarketRouter) *Gateway {
	return &Gateway{
		router: router,
	}
}

func (g *Gateway) Route(market string, order *orderbook.Order) error {
	book, err := g.router.Route(market)
	if err != nil {
		return err
	}
	book.AddOrder(order)
	return nil
}
