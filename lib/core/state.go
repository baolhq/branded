package core

import (
	"sync"
)

var instance *StateManager
var once sync.Once

const (
	StateIdle         int = iota // Free movement, no window opened, no unit selected
	StateUnitSelected            // Unit selected, waiting for commands
	StateMenuOpen                // Menu window open
	StateCombat                  // A combat is happening
	StateDialog                  // A dialog is opening
)

type StateManager struct {
	Mu    sync.Mutex
	State int
}

// Ensuring only one instance exists
func GetStateManager() *StateManager {
	once.Do(func() {
		instance = &StateManager{State: StateIdle} // Default state
	})
	return instance
}

// SetState updates the current state
func (sm *StateManager) SetState(newState int) {
	sm.Mu.Lock()
	defer sm.Mu.Unlock()
	sm.State = newState
}

// GetState returns the current state
func (sm *StateManager) GetState() int {
	return sm.State
}
