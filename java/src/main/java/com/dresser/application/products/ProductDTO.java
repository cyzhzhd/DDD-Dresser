package com.dresser.application.products;

import com.dresser.domain.products.Product;
import java.math.BigDecimal;
import java.util.List;

public class ProductDTO {
    private int id;
    private int brandId;
    private String name;
    private String category;
    private BigDecimal price;
    private List<String> sizes;
    private boolean deleted;
    
    public ProductDTO() {
    }
    
    public ProductDTO(int id, int brandId, String name, String category, BigDecimal price, List<String> sizes, boolean deleted) {
        this.id = id;
        this.brandId = brandId;
        this.name = name;
        this.category = category;
        this.price = price;
        this.sizes = sizes;
        this.deleted = deleted;
    }
    
    public static ProductDTO fromDomain(Product product) {
        return new ProductDTO(
            product.getId().getValue(),
            product.getBrandId().getValue(),
            product.getName(),
            product.getCategory(),
            product.getPrice(),
            product.getSizes(),
            product.isDeleted()
        );
    }
    
    // Getters
    public int getId() {
        return id;
    }
    
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
    
    public boolean isDeleted() {
        return deleted;
    }
    
    // Setters
    public void setId(int id) {
        this.id = id;
    }
    
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
    
    public void setDeleted(boolean deleted) {
        this.deleted = deleted;
    }
} 