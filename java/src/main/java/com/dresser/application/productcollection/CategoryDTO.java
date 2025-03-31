package com.dresser.application.productcollection;

import java.util.List;

public class CategoryDTO {
    private String name;
    private List<String> brands;
    private int count;
    
    public CategoryDTO() {
    }
    
    public CategoryDTO(String name, List<String> brands, int count) {
        this.name = name;
        this.brands = brands;
        this.count = count;
    }
    
    // Getters
    public String getName() {
        return name;
    }
    
    public List<String> getBrands() {
        return brands;
    }
    
    public int getCount() {
        return count;
    }
    
    // Setters
    public void setName(String name) {
        this.name = name;
    }
    
    public void setBrands(List<String> brands) {
        this.brands = brands;
    }
    
    public void setCount(int count) {
        this.count = count;
    }
} 