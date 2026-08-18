package ownership

import "fmt"

type Manager struct {
	owners map[string]string
}

func NewManager() *Manager {
	return &Manager{
		owners: make(map[string]string),
	}
}

func (m *Manager) Assign(market, engine string) {
	m.owners[market] = engine
}

func (m *Manager) Owner(market string) (string, error) {
	engine, ok := m.owners[market]
	if !ok {
		return "", fmt.Errorf("owner for %s not found", market)
	}

	return engine, nil
}
