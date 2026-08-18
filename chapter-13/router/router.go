package router

import (
	"fmt"
	"order-book/engine"
	"order-book/ownership"
)

type MarketRouter struct {
	manager   *engine.Manager
	ownership *ownership.Manager
}

func New(manager *engine.Manager, ownership *ownership.Manager) *MarketRouter {
	return &MarketRouter{
		manager:   manager,
		ownership: ownership,
	}
}



func (r *MarketRouter) Route(market string) (*engine.Engine, error) {
	engineID, err := r.ownership.Owner(market)
	if err != nil {
		return nil, err
	}
	engine, err := r.manager.Get(engineID)
	if err != nil {
		return nil, err
	}
	fmt.Printf("[Router] %s -> %s\n", market, engineID)
	return engine, nil
}
