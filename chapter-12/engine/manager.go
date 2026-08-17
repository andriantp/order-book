package engine

import (
	"fmt"
	"order-book/orderbook"
)

type Manager struct {
	books map[string]*orderbook.OrderBook
}

func NewManager() *Manager {
	return &Manager{
		books: make(map[string]*orderbook.OrderBook),
	}
}

func (m *Manager) Register(engine string, book *orderbook.OrderBook) {
	m.books[engine] = book
}

func (m *Manager) Get(engine string) (*orderbook.OrderBook, error) {
	book, ok := m.books[engine]
	if !ok {
		return nil, fmt.Errorf("matching engine %q not found", engine)
	}

	return book, nil
}
