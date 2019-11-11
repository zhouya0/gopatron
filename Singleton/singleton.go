package Singleton

import (
	"fmt"
	"sync"
)

type Manager struct {
	name string
}

func (p Manager) Manage() {
	fmt.Printf("manage %s", p.name)
}

func (p *Manager) Set(name string) {
	p.name = name
}

var m *Manager
var once sync.Once

func GetInstance() *Manager {
	once.Do(func() {
		m = &Manager{}
	})
	return m
}
