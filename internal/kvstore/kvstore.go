package kvstore

import (
	"sync"
)

type Node struct {
	mu    sync.RWMutex
	store map[string]string
}

func NewNode() *Node {
	return &Node{
		store: make(map[string]string),
	}
}

func (n *Node) Put(key, value string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.store[key] = value
}

func (n *Node) Get(key string) (string, bool) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	value, ok := n.store[key]
	return value, ok
}

func (n *Node) Delete(key string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	delete(n.store, key)
}
