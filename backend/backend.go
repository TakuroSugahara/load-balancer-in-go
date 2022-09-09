package backend

import "sync"

type Backend struct {
	URL    string `json:"url"`
	IsDead bool
	mu     sync.RWMutex
}

type Backends = []Backend

func (b *Backend) SetDead(status bool) {
	b.mu.Lock()
	b.IsDead = status
	b.mu.Unlock()
}

func (b *Backend) GetIsDead() bool {
	b.mu.RLock()
	isAlive := b.IsDead
	b.mu.RUnlock()
	return isAlive
}
