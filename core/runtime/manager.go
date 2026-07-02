package runtime

import (
	"sync"
)

type Manager struct {
	mu      sync.RWMutex
	engines map[string]Engine
	order   []string
}

func NewManager() *Manager {
	return &Manager{
		engines: make(map[string]Engine),
		order:   []string{},
	}
}

func (m *Manager) Register(e Engine) {
	m.mu.Lock()
	defer m.mu.Unlock()

	name := e.Name()
	if _, ok := m.engines[name]; ok {
		return
	}

	m.engines[name] = e
	m.order = append(m.order, name)
}

func (m *Manager) Engines() []Engine {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]Engine, 0, len(m.order))
	for _, name := range m.order {
		result = append(result, m.engines[name])
	}
	return result
}