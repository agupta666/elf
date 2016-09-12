package store

// KVSet represents key value sets
type KVSet map[string]string

var store = make(map[string]KVSet)

// SaveKVSet saves a key value set with a given name
func SaveKVSet(key string, kvs KVSet) {
	store[key] = kvs
}

// GetKVSet looks up a key value set by name from the store
func GetKVSet(key string) KVSet {
	return store[key]
}
