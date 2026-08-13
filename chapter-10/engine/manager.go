package engine

import "order-book/orderbook"

type Manager struct {
	books map[string]*orderbook.OrderBook
}

func NewManager() *Manager {
	return &Manager{
		books: make(map[string]*orderbook.OrderBook),
	}
}

func (m *Manager) Register(symbol string, book *orderbook.OrderBook) {
	m.books[symbol] = book
}

func (m *Manager) Get(symbol string) (*orderbook.OrderBook, bool) {
	book, ok := m.books[symbol]
	return book, ok
}
