package snapshot

import (
	"bytes"
	"distributed-kv-store/internal/kvstore"
	"encoding/gob"
	"os"
)

// CreateSnapshot создает снимок состояния узла и сохраняет его в файл
func CreateSnapshot(filename string, node *kvstore.Node) error {
	// Блокировка узла для безопасного чтения данных
	node.RLock()
	defer node.RUnlock()

	// Создание буфера и кодировщика gob для сериализации данных
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(node.GetStore()); err != nil { // Используем метод GetStore для доступа к store
		return err
	}

	// Запись сериализованных данных в файл
	return os.WriteFile(filename, buf.Bytes(), 0644)
}

// RestoreFromSnapshot восстанавливает состояние узла из снимка
func RestoreFromSnapshot(filename string, node *kvstore.Node) error {
	// Блокировка узла для безопасной записи данных
	node.Lock()
	defer node.Unlock()

	// Чтение данных из файла
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Декодирование данных из формата gob
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	store := make(map[string]string)
	if err := dec.Decode(&store); err != nil {
		return err
	}

	// Обновление хранилища узла с восстановленными данными
	node.UpdateStore(store) // Используем метод UpdateStore для обновления store
	return nil
}
