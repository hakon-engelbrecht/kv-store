// Package store contains interface and implementations of a store.
// A store is a map-like data structure for holding data referenced by keys.
package store

// Store is a basic map-like data structure.
// It holds string keys associated to string values.
type Store interface {
	// Sets the provided string to the provided value in the store.
	Set(key string, val string)

	// Get returns the value associated with the provided key and if the key is present.
	// If the key is not present this function will return an empty string as the value.
	Get(key string) (string, bool)

	// Delete deletes the key from the store. If the key is not present, nothing happens.
	Delete(key string) bool

	// Exists returns if the provided key is present in the store.
	Exists(key string) bool

	// Keys returns all keys contained in the store sorted alphabetically.
	Keys() []string
}
