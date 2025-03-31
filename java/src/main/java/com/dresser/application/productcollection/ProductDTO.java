package com.dresser.application.productcollection;

import java.util.List;

public class ProductDTO {
    private String id;
    private String name;
    private String brand;
    private String category;
    private double price;
    private List<String> sizes;
    
    public ProductDTO() {
    }
    
    public ProductDTO(String id, String name, String brand, String category, double price, List<String> sizes) {
        this.id = id;
        this.name = name;
        this.brand = brand;
        this.category = category;
        this.price = price;
        this.sizes = sizes;
    }
    
    // Getters
    public String getId() {
        return id;
    }
    
    public String getName() {
        return name;
    }
    
    public String getBrand() {
        return brand;
    }
    
    public String getCategory() {
        return category;
    }
    
    public double getPrice() {
        return price;
    }
    
    public List<String> getSizes() {
        return sizes;
    }
    
    // Setters
    public void setId(String id) {
        this.id = id;
    }
    
    public void setName(String name) {
        this.name = name;
    }
    
    public void setBrand(String brand) {
        this.brand = brand;
    }
    
    public void setCategory(String category) {
        this.category = category;
    }
    
    public void setPrice(double price) {
        this.price = price;
    }
    
    public void setSizes(List<String> sizes) {
        this.sizes = sizes;
    }
} 