package storage

import (
	"sync"
	"time"
)

var (
	cipherKeys = make(map[string]string)
	cipherLock sync.Mutex
)

func StoreCipherKey(key string) {
	cipherLock.Lock()
	defer cipherLock.Unlock()

	cipherKeys[key] = key
}

func GetCipherKeys() map[string]string {
	cipherLock.Lock()
	defer cipherLock.Unlock()

	return cipherKeys
}

func ExpireCipherKeys(duration time.Duration) {
	time.Sleep(duration)
	cipherLock.Lock()
	defer cipherLock.Unlock()
	cipherKeys = make(map[string]string)
}
