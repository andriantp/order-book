package detector

type Manager struct {
	status map[string]bool
}

func NewManager() *Manager {
	return &Manager{
		status: make(map[string]bool),
	}
}

func (m *Manager) Register(engine string) {
	m.status[engine] = true
}

func (m *Manager) Fail(engine string) {
	m.status[engine] = false
}

func (m *Manager) Recover(engine string) {
	m.status[engine] = true
}

func (m *Manager) Healthy(engine string) bool {
	return m.status[engine]
}
