package replication

import "fmt"

type Manager struct {
	replicas map[string][]string
}

func NewManager() *Manager {
	return &Manager{
		replicas: make(map[string][]string),
	}
}

func (m *Manager) Assign(market, replica string) {
	m.replicas[market] = append(m.replicas[market], replica)
}

func (m *Manager) Replicas(market string) []string {
	return m.replicas[market]
}

func (m *Manager) Promote(
	market string,
	healthy func(string) bool) (string, error) {
	replicas := m.Replicas(market)
	for _, replica := range replicas {
		if healthy(replica) {
			return replica, nil
		}
	}
	return "", fmt.Errorf("no healthy replica for %s", market)
}
