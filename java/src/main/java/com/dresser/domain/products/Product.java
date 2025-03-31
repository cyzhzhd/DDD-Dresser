package com.dresser.domain.products;

import com.dresser.domain.brands.BrandId;

import java.math.BigDecimal;
import java.util.List;

public class Product {
    private final ProductId id;
    private final BrandId brandId;
    private String name;
    private String category;
    private BigDecimal price;
    private List<String> sizes;
    private boolean deleted;
    
    private Product(ProductId id, BrandId brandId, String name, String category, BigDecimal price, List<String> sizes) {
        this.id = id;
        this.brandId = brandId;
        this.name = name;
        this.category = category;
        this.price = price;
        this.sizes = sizes;
        this.deleted = false;
    }
    
    public static Product create(int id, int brandId, String name, String category, BigDecimal price, List<String> sizes) {
        return new Product(
            new ProductId(id),
            new BrandId(brandId),
            name,
            category,
            price,
            sizes
        );
    }
    
    public void update(String name, String category, BigDecimal price, List<String> sizes) {
        this.name = name;
        this.category = category;
        this.price = price;
        this.sizes = sizes;
    }
    
    public void markAsDeleted() {
        this.setDeleted(true);
    }
    
    // Getters
    public ProductId getId() {
        return id;
    }
    
    public BrandId getBrandId() {
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
    
    public boolean isDeleted() {
        return deleted;
    }
    
    // Setters
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
    
    public void setDeleted(boolean deleted) {
        this.deleted = deleted;
    }
} 