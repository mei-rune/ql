package ql

type Mutex struct{}

func (mu *Mutex) Lock()   {}
func (mu *Mutex) Unlock() {}

type RWMutex struct{}

func (mu *RWMutex) Lock()   {}
func (mu *RWMutex) Unlock() {}

func (mu *RWMutex) RLock()   {}
func (mu *RWMutex) RUnlock() {}
