package db

import (
	"context"
	"fmt"
)

// DBType represents the type of database to use
type DBType string

const (
	// Memory is an in-memory database implementation
	Memory DBType = "memory"
	// Postgres is a PostgreSQL database implementation
	Postgres DBType = "postgres"
)

// Config holds configuration for database connection
type Config struct {
	Type     DBType `json:"type"`
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	DBName   string `json:"db_name,omitempty"`
}

// DB is an interface that defines generic database operations
type DB interface {
	// Create adds a new entity to the database
	Create(ctx context.Context, collection string, id int, entity interface{}) (interface{}, error)

	// FindByID retrieves an entity by ID from the database
	FindByID(ctx context.Context, collection string, id interface{}) (interface{}, error)

	// FindAll retrieves all entities of a collection from the database
	FindAll(ctx context.Context, collection string) ([]interface{}, error)

	// FindBy retrieves entities by a filter from the database
	FindBy(ctx context.Context, collection string, filter map[string]interface{}) ([]interface{}, error)

	// Update updates an entity in the database
	Update(ctx context.Context, collection string, id interface{}, entity interface{}) error

	// Delete removes an entity from the database
	Delete(ctx context.Context, collection string, id interface{}) error

	// NextID generates the next available ID for a collection
	NextID(collection string) int

	// Close database connection
	Close() error
}

// New creates a new database instance based on the provided configuration
func New(config Config) (DB, error) {
	switch config.Type {
	case Memory:
		return NewMemoryDB(), nil
	case Postgres:
		return NewPostgresDB(config), nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", config.Type)
	}
}
