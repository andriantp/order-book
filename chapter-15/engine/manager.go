package engine

import (
	"fmt"
	"order-book/marketdata"
	"order-book/orderbook"
)

type Engine struct {
	ID         string
	Book       *orderbook.OrderBook
	EventStore orderbook.EventStore
	ReadModel  *marketdata.ReadModel
}

type Manager struct {
	engines map[string]*Engine
}

func NewManager() *Manager {
	return &Manager{
		engines: make(map[string]*Engine),
	}
}

func (m *Manager) Register(id string, engine *Engine) {
	m.engines[id] = engine
}

func (m *Manager) Get(id string) (*Engine, error) {
	engine, ok := m.engines[id]
	if !ok {
		return nil, fmt.Errorf("matching engine %q not found", id)
	}

	return engine, nil
}
