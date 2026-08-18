package gateway

import (
	"fmt"
	"order-book/engine"
	"order-book/orderbook"
	"order-book/replication"
	"order-book/router"
)

type Gateway struct {
	router      *router.MarketRouter
	manager     *engine.Manager
	replication *replication.Manager
}

func New(
	router *router.MarketRouter,
	manager *engine.Manager,
	replication *replication.Manager,
) *Gateway {
	return &Gateway{
		router:      router,
		manager:     manager,
		replication: replication,
	}
}

func (g *Gateway) Route(market string, order *orderbook.Order) error {
	engine, err := g.router.Route(market)
	if err != nil {
		return err
	}

	before := len(engine.EventStore.Load())
	engine.Book.AddOrder(order)
	events := engine.EventStore.Load()[before:]
	replicas := g.replication.Replicas(market)

	for _, replicaID := range replicas {
		replica, err := g.manager.Get(replicaID)
		if err != nil {
			return err
		}
		fmt.Printf("[Replication] %s -> %s\n", market, replicaID)
		replica.Book.Replay(events)
	}
	return nil
}
