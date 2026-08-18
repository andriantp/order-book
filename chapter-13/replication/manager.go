package replication

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
