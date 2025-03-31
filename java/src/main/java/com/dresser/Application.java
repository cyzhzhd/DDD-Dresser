package com.dresser;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ApplicationListener;
import org.springframework.context.event.ContextRefreshedEvent;
import org.springframework.stereotype.Component;

@SpringBootApplication
public class Application {
    private static final Logger logger = LoggerFactory.getLogger(Application.class);
    
    public static void main(String[] args) {
        logger.info("Starting DDD-Dresser application...");
        SpringApplication.run(Application.class, args);
        logger.info("DDD-Dresser application started successfully!");
    }
    
    @Component
    public static class StartupListener implements ApplicationListener<ContextRefreshedEvent> {
        private static final Logger logger = LoggerFactory.getLogger(StartupListener.class);
        
        @Override
        public void onApplicationEvent(ContextRefreshedEvent event) {
            logger.info("API is now available at http://localhost:8080/api/");
            logger.info("Available endpoints:");
            logger.info("  - GET  /api/brands");
            logger.info("  - GET  /api/brands/{id}");
            logger.info("  - POST /api/brands");
            logger.info("  - DELETE /api/brands/{id}");
            logger.info("  - GET  /api/products");
            logger.info("  - GET  /api/products/{id}");
            logger.info("  - POST /api/products");
            logger.info("  - PUT  /api/products/{id}");
            logger.info("  - DELETE /api/products/{id}");
            logger.info("  - GET  /api/categories");
        }
    }
} 