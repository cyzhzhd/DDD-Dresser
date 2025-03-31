package com.dresser.application.products;

import java.math.BigDecimal;
import java.util.List;

public class ProductUpdateCommand {
    private int id;
    private String name;
    private String category;
    private BigDecimal price;
    private List<String> sizes;
    
    public ProductUpdateCommand() {
    }
    
    public ProductUpdateCommand(int id, String name, String category, BigDecimal price, List<String> sizes) {
        this.id = id;
        this.name = name;
        this.category = category;
        this.price = price;
        this.sizes = sizes;
    }
    
    // Getters
    public int getId() {
        return id;
    }
    
    public String getName() {
        return name;
    }
    
    public String getCategory() {
        return category;
    }
    
    public BigDecimal getPrice() {
        return price;
    }
    
    public List<String> getSizes() {
        return sizes;
    }
    
    // Setters
    public void setId(int id) {
        this.id = id;
    }
    
    public void setName(String name) {
        this.name = name;
    }
    
    public void setCategory(String category) {
        this.category = category;
    }
    
    public void setPrice(BigDecimal price) {
        this.price = price;
    }
    
    public void setSizes(List<String> sizes) {
        this.sizes = sizes;
    }
} 