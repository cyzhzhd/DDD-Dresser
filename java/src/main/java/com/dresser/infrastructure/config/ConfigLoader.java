package com.dresser.infrastructure.config;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.stereotype.Component;

import java.io.File;
import java.io.IOException;

@Component
public class ConfigLoader {
    
    public Config loadDefault() throws IOException {
        ObjectMapper mapper = new ObjectMapper();
        File configFile = new File("config.json");
        
        if (!configFile.exists()) {
            // Return default config if file doesn't exist
            return new Config(new DatabaseConfig("memory"));
        }
        
        return mapper.readValue(configFile, Config.class);
    }
} 