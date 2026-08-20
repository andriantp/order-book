package gateway

import (
	"fmt"
	"order-book/detector"
	"order-book/engine"
	"order-book/orderbook"
	"order-book/replication"
	"order-book/router"
)

type Gateway struct {
	router      *router.MarketRouter
	manager     *engine.Manager
	replication *replication.Manager
	detector    *detector.Manager
}

func New(
	router *router.MarketRouter,
	manager *engine.Manager,
	replication *replication.Manager,
	detector *detector.Manager,
) *Gateway {
	return &Gateway{
		router:      router,
		manager:     manager,
		replication: replication,
		detector:    detector,
	}
}

func (g *Gateway) Route(market string, order *orderbook.Order) error {
	engine, err := g.router.Route(market)
	if err != nil {
		return err
	}

	if !g.detector.Healthy(engine.ID) {
		fmt.Printf("[Failover] %s is unavailable\n", engine.ID)
		replicaID, err := g.replication.Promote(
			market,
			g.detector.Healthy,
		)
		if err != nil {
			return err
		}
		fmt.Printf("[Failover] Promoting %s\n", replicaID)
		g.router.Transfer(market, replicaID)
		engine, err = g.router.Route(market)
		if err != nil {
			return err
		}
	}

	before := len(engine.EventStore.Load())
	engine.Book.AddOrder(order)
	events := engine.EventStore.Load()[before:]
	engine.ReadModel.Replay(events)
	replicas := g.replication.Replicas(market)

	for _, replicaID := range replicas {
		// Skip the current primary.
		if replicaID == engine.ID {
			continue
		}

		replica, err := g.manager.Get(replicaID)
		if err != nil {
			return err
		}

		fmt.Printf("[Replication] %s -> %s\n", market, replicaID)
		replica.Book.Replay(events)
		replica.ReadModel.Replay(events)
	}
	return nil
}
