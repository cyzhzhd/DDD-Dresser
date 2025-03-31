package com.dresser.infrastructure.db;

import com.dresser.infrastructure.config.DatabaseConfig;
import org.springframework.stereotype.Component;

@Component
public class DBFactory {
    
    public DB create(DatabaseConfig config) {
        switch (config.getType()) {
            case "memory":
                return new MemoryDB();
            case "postgres":
                return new PostgresDB(config);
            default:
                throw new IllegalArgumentException("Unsupported database type: " + config.getType());
        }
    }
} 