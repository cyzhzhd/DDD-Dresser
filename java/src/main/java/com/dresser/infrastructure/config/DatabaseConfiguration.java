package com.dresser.infrastructure.config;

import com.dresser.infrastructure.db.DB;
import com.dresser.infrastructure.db.DBFactory;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.io.IOException;

@Configuration
public class DatabaseConfiguration {
    
    private final ConfigLoader configLoader;
    private final DBFactory dbFactory;
    
    public DatabaseConfiguration(ConfigLoader configLoader, DBFactory dbFactory) {
        this.configLoader = configLoader;
        this.dbFactory = dbFactory;
    }
    
    @Bean
    public DB database() throws IOException {
        Config config = configLoader.loadDefault();
        return dbFactory.create(config.getDatabase());
    }
} 