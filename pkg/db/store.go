package db

import "sync"

type Store struct {
	mu sync.RWMutex
	data map[string]bool // map[processName]isWhitelisted
}

func NewStore() *Store {
	return &Store{data: make(map[string]bool)}
}

func (s *Store) IsWhitelisted(process string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data[process]
}

func (s *Store) Whitelist(process string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[process] = true
}
