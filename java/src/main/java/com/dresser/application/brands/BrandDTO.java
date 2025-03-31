package com.dresser.application.brands;

import com.dresser.domain.brands.Brand;

public class BrandDTO {
    private int id;
    private String name;
    
    public BrandDTO() {
    }
    
    public BrandDTO(int id, String name) {
        this.id = id;
        this.name = name;
    }
    
    public static BrandDTO fromDomain(Brand brand) {
        return new BrandDTO(brand.getId().getValue(), brand.getName());
    }
    
    // Getters
    public int getId() {
        return id;
    }
    
    public String getName() {
        return name;
    }
    
    // Setters
    public void setId(int id) {
        this.id = id;
    }
    
    public void setName(String name) {
        this.name = name;
    }
} 