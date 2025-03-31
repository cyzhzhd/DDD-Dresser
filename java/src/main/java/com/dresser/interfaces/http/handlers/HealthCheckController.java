package com.dresser.interfaces.http.handlers;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.time.LocalDateTime;
import java.util.HashMap;
import java.util.Map;

@RestController
@RequestMapping("/api/health")
public class HealthCheckController {
    
    private final LocalDateTime startTime = LocalDateTime.now();
    
    @GetMapping
    public ResponseEntity<Map<String, Object>> healthCheck() {
        Map<String, Object> response = new HashMap<>();
        response.put("status", "UP");
        response.put("startTime", startTime.toString());
        response.put("currentTime", LocalDateTime.now().toString());
        
        // Add system information
        Map<String, String> system = new HashMap<>();
        system.put("javaVersion", System.getProperty("java.version"));
        system.put("osName", System.getProperty("os.name"));
        system.put("osVersion", System.getProperty("os.version"));
        system.put("availableProcessors", String.valueOf(Runtime.getRuntime().availableProcessors()));
        
        response.put("system", system);
        
        return ResponseEntity.ok(response);
    }
} 