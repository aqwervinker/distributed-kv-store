package snapshot

import (
	"bytes"
	"distributed-kv-store/internal/kvstore"
	"encoding/gob"
	"os"
)

func CreateSnapshot(filename string, node *kvstore.Node) error {
	node.Mu.RLock()
	defer node.Mu.RUnlock()

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(node.Store); err != nil {
		return err
	}

	return os.WriteFile(filename, buf.Bytes(), 0644)
}

func RestoreFromSnapshot(filename string, node *kvstore.Node) error {
	node.Mu.Lock()
	defer node.Mu.Unlock()

	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	store := make(map[string]string)
	if err := dec.Decode(&store); err != nil {
		return err
	}

	node.Store = store
	return nil
}
