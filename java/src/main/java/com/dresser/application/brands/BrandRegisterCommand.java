package com.dresser.application.brands;

public class BrandRegisterCommand {
    private String name;
    
    public BrandRegisterCommand() {
    }
    
    public BrandRegisterCommand(String name) {
        this.name = name;
    }
    
    // Getter
    public String getName() {
        return name;
    }
    
    // Setter
    public void setName(String name) {
        this.name = name;
    }
} 