package kvstore

import (
	"sync"
)

type Node struct {
	mu    sync.RWMutex
	store map[string]string
}

// NewNode создает и возвращает новый узел Node
func NewNode() *Node {
	return &Node{
		store: make(map[string]string),
	}
}

// Put добавляет или обновляет значение по ключу в store
func (n *Node) Put(key, value string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.store[key] = value
}

// Get возвращает значение по ключу из store и флаг его существования
func (n *Node) Get(key string) (string, bool) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	value, ok := n.store[key]
	return value, ok
}

// Delete удаляет значение по ключу из store
func (n *Node) Delete(key string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	delete(n.store, key)
}

// Lock блокирует Node для записи
func (n *Node) Lock() {
	n.mu.Lock()
}

// Unlock разблокирует Node для записи
func (n *Node) Unlock() {
	n.mu.Unlock()
}

// RLock блокирует Node для чтения
func (n *Node) RLock() {
	n.mu.RLock()
}

// RUnlock разблокирует Node для чтения
func (n *Node) RUnlock() {
	n.mu.RUnlock()
}

// GetStore возвращает ссылку на store для чтения или модификации
func (n *Node) GetStore() map[string]string {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.store
}

// UpdateStore обновляет store новыми данными
func (n *Node) UpdateStore(newStore map[string]string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.store = newStore
}
