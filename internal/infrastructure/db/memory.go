package db

import (
	"context"
	"fmt"
	"sync"
)

// MemoryDB implements DB interface using in-memory storage
type MemoryDB struct {
	sync.RWMutex
	collections map[string]map[interface{}]interface{}
	lastIDs     map[string]int
}

// NewMemoryDB creates a new instance of MemoryDB
func NewMemoryDB() DB {
	return &MemoryDB{
		collections: make(map[string]map[interface{}]interface{}),
		lastIDs:     make(map[string]int),
	}
}

// Create implements DB interface
func (db *MemoryDB) Create(ctx context.Context, collection string, id int, entity interface{}) (interface{}, error) {
	db.Lock()
	defer db.Unlock()

	// Initialize collection if it doesn't exist
	if _, exists := db.collections[collection]; !exists {
		db.collections[collection] = make(map[interface{}]interface{})
	}
	// Check if entity with same ID already exists
	if _, exists := db.collections[collection][id]; exists {
		return nil, fmt.Errorf("entity with ID %v already exists in collection %s", id, collection)
	}

	// Store entity
	db.collections[collection][id] = entity
	return entity, nil
}

// FindByID implements DB interface
func (db *MemoryDB) FindByID(ctx context.Context, collection string, id interface{}) (interface{}, error) {
	db.RLock()
	defer db.RUnlock()

	// Check if collection exists
	collectionMap, exists := db.collections[collection]
	if !exists {
		return nil, fmt.Errorf("collection %s does not exist", collection)
	}

	// Find entity
	entity, exists := collectionMap[id]
	if !exists {
		return nil, fmt.Errorf("entity with ID %v not found in collection %s", id, collection)
	}

	return entity, nil
}

// FindAll implements DB interface
func (db *MemoryDB) FindAll(ctx context.Context, collection string) ([]interface{}, error) {
	db.RLock()
	defer db.RUnlock()

	// Check if collection exists
	collectionMap, exists := db.collections[collection]
	if !exists {
		return []interface{}{}, nil
	}

	// Convert map values to slice
	result := make([]interface{}, 0, len(collectionMap))
	for _, entity := range collectionMap {
		result = append(result, entity)
	}

	return result, nil
}

// FindBy implements DB interface
func (db *MemoryDB) FindBy(ctx context.Context, collection string, filter map[string]interface{}) ([]interface{}, error) {
	db.RLock()
	defer db.RUnlock()

	// Check if collection exists
	collectionMap, exists := db.collections[collection]
	if !exists {
		return []interface{}{}, nil
	}

	// For memory implementation, we'll do a simple matching on all entities
	// A real implementation would be more sophisticated
	result := make([]interface{}, 0)
	for _, entity := range collectionMap {
		if matchesFilter(entity, filter) {
			result = append(result, entity)
		}
	}

	return result, nil
}

// Update implements DB interface
func (db *MemoryDB) Update(ctx context.Context, collection string, id interface{}, entity interface{}) error {
	db.Lock()
	defer db.Unlock()

	// Check if collection exists
	collectionMap, exists := db.collections[collection]
	if !exists {
		return fmt.Errorf("collection %s does not exist", collection)
	}

	// Check if entity exists
	if _, exists := collectionMap[id]; !exists {
		return fmt.Errorf("entity with ID %v not found in collection %s", id, collection)
	}

	// Update entity
	collectionMap[id] = entity
	return nil
}

// Delete implements DB interface
func (db *MemoryDB) Delete(ctx context.Context, collection string, id interface{}) error {
	db.Lock()
	defer db.Unlock()

	// Check if collection exists
	collectionMap, exists := db.collections[collection]
	if !exists {
		return fmt.Errorf("collection %s does not exist", collection)
	}

	// Check if entity exists
	if _, exists := collectionMap[id]; !exists {
		return fmt.Errorf("entity with ID %v not found in collection %s", id, collection)
	}

	// Delete entity
	delete(collectionMap, id)
	return nil
}

// NextID implements DB interface
func (db *MemoryDB) NextID(collection string) int {
	db.Lock()
	defer db.Unlock()

	// Initialize last ID for collection if it doesn't exist
	if _, exists := db.lastIDs[collection]; !exists {
		db.lastIDs[collection] = 0
	}

	// Increment and return next ID
	db.lastIDs[collection]++
	return db.lastIDs[collection]
}

// Close implements DB interface
func (db *MemoryDB) Close() error {
	// No-op for memory DB
	return nil
}

// Helper functions

// matchesFilter checks if entity matches the given filter
func matchesFilter(entity interface{}, filter map[string]interface{}) bool {
	// For a real implementation, we'd use reflection to match fields
	// For simplicity, we'll just return true for empty filters
	return len(filter) == 0
}
