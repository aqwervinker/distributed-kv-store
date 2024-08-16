package consensus

import (
	"distributed-kv-store/internal/kvstore"
	"sync"
)

type Zab struct {
	nodes []*kvstore.Node
	mu    sync.RWMutex
}

func NewZab() *Zab {
	return &Zab{
		nodes: []*kvstore.Node{},
	}
}

func (z *Zab) AddNode(node *kvstore.Node) {
	z.mu.Lock()
	defer z.mu.Unlock()
	z.nodes = append(z.nodes, node)
}

func (z *Zab) Put(key, value string) {
	z.mu.Lock()
	defer z.mu.Unlock()
	for _, node := range z.nodes {
		node.Put(key, value)
	}
}

func (z *Zab) Get(key string) (string, bool) {
	z.mu.RLock()
	defer z.mu.RUnlock()
	if len(z.nodes) == 0 {
		return "", false
	}
	return z.nodes[0].Get(key)
}

func (z *Zab) Delete(key string) {
	z.mu.Lock()
	defer z.mu.Unlock()
	for _, node := range z.nodes {
		node.Delete(key)
	}
}
