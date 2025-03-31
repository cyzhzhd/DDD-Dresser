package com.dresser.application.products;

import java.math.BigDecimal;
import java.util.List;

public class ProductRegisterCommand {
    private int brandId;
    private String name;
    private String category;
    private BigDecimal price;
    private List<String> sizes;
    
    public ProductRegisterCommand() {
    }
    
    public ProductRegisterCommand(int brandId, String name, String category, BigDecimal price, List<String> sizes) {
        this.brandId = brandId;
        this.name = name;
        this.category = category;
        this.price = price;
        this.sizes = sizes;
    }
    
    // Getters
    public int getBrandId() {
        return brandId;
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
    public void setBrandId(int brandId) {
        this.brandId = brandId;
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