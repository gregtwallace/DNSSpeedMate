package safemap

import (
	"encoding/json"
	"errors"
	"sync"
)

// errors
var errMapKeyDoesntExist = errors.New("specified map key does not exist")

// SafeMap is a map with a mutex
type SafeMap[K comparable, V any] struct {
	m  map[K]V
	mu sync.RWMutex
}

// NewSafeMap creates a new SafeMap
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		m: make(map[K]V),
	}
}

// MarshalJSON marshals and returns the map
func (sm *SafeMap[K, V]) MarshalJSON() ([]byte, error) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	return json.Marshal(sm.m)
}

// Read returns the value from the specified key. If the key
// does not exist, an error is returned.
func (sm *SafeMap[K, V]) Read(key K) (V, error) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	// read data
	value, exists := sm.m[key]
	if !exists {
		return value, errMapKeyDoesntExist
	}

	return value, nil
}

// Copy returns a copy of the internal map's current state
func (sm *SafeMap[K, V]) Copy() map[K]V {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	// copy to new map
	m := make(map[K]V)
	for k, v := range sm.m {
		m[k] = v
	}

	return m
}

// Add creates the named key and inserts the specified value.
// If the key already exists, true and the existing value are
// returned instead.
func (sm *SafeMap[K, V]) Add(key K, value V) (bool, V) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// if key exists, return true and the existing value
	exisingValue, exists := sm.m[key]
	if exists {
		return true, exisingValue
	}

	// if not, add the key and value
	sm.m[key] = value

	return false, value
}
