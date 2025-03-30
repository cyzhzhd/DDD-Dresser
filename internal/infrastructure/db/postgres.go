package db

import (
	"context"
)

// PostgresDB implements DB interface using PostgreSQL
type PostgresDB struct {
	config Config
}

// NewPostgresDB creates a new instance of PostgresDB
func NewPostgresDB(config Config) DB {
	return &PostgresDB{
		config: config,
	}
}

// Create implements DB interface
func (db *PostgresDB) Create(ctx context.Context, collection string, id int, entity interface{}) (interface{}, error) {
	// Not implemented yet
	return nil, nil
}

// FindByID implements DB interface
func (db *PostgresDB) FindByID(ctx context.Context, collection string, id interface{}) (interface{}, error) {
	// Not implemented yet
	return nil, nil
}

// FindAll implements DB interface
func (db *PostgresDB) FindAll(ctx context.Context, collection string) ([]interface{}, error) {
	// Not implemented yet
	return nil, nil
}

// FindBy implements DB interface
func (db *PostgresDB) FindBy(ctx context.Context, collection string, filter map[string]interface{}) ([]interface{}, error) {
	// Not implemented yet
	return nil, nil
}

// Update implements DB interface
func (db *PostgresDB) Update(ctx context.Context, collection string, id interface{}, entity interface{}) error {
	// Not implemented yet
	return nil
}

// Delete implements DB interface
func (db *PostgresDB) Delete(ctx context.Context, collection string, id interface{}) error {
	// Not implemented yet
	return nil
}

// NextID implements DB interface
func (db *PostgresDB) NextID(collection string) int {
	// Not implemented yet
	return 0
}

// Close implements DB interface
func (db *PostgresDB) Close() error {
	// Not implemented yet
	return nil
}
