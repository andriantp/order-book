package engine

import "order-book/orderbook"

type Manager struct {
	books []*orderbook.OrderBook
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Register(book *orderbook.OrderBook) {
	m.books = append(m.books, book)
}

func (m *Manager) Books() []*orderbook.OrderBook {
	return m.books
}
