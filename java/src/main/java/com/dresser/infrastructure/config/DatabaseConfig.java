package com.dresser.infrastructure.config;

public class DatabaseConfig {
    private String type;
    private String host;
    private int port;
    private String user;
    private String password;
    private String dbName;
    
    public DatabaseConfig() {
    }
    
    public DatabaseConfig(String type) {
        this.type = type;
    }
    
    public DatabaseConfig(String type, String host, int port, String user, String password, String dbName) {
        this.type = type;
        this.host = host;
        this.port = port;
        this.user = user;
        this.password = password;
        this.dbName = dbName;
    }
    
    // Getters
    public String getType() {
        return type;
    }
    
    public String getHost() {
        return host;
    }
    
    public int getPort() {
        return port;
    }
    
    public String getUser() {
        return user;
    }
    
    public String getPassword() {
        return password;
    }
    
    public String getDbName() {
        return dbName;
    }
    
    // Setters
    public void setType(String type) {
        this.type = type;
    }
    
    public void setHost(String host) {
        this.host = host;
    }
    
    public void setPort(int port) {
        this.port = port;
    }
    
    public void setUser(String user) {
        this.user = user;
    }
    
    public void setPassword(String password) {
        this.password = password;
    }
    
    public void setDbName(String dbName) {
        this.dbName = dbName;
    }
} 