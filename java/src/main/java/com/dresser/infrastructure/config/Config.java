package com.dresser.infrastructure.config;

public class Config {
    private DatabaseConfig database;
    
    public Config() {
    }
    
    public Config(DatabaseConfig database) {
        this.database = database;
    }
    
    // Getter
    public DatabaseConfig getDatabase() {
        return database;
    }
    
    // Setter
    public void setDatabase(DatabaseConfig database) {
        this.database = database;
    }
} 